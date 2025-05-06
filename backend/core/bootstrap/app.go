package bootstrap

import (
	"context"
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/route"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/rbac"
	"github.com/KokoiRuby/rbac-based-management-system/backend/infra/persistence"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"sync"
	"time"
)

var (
	SigChan      = make(chan os.Signal, 1)
	ShutDownChan = make(chan struct{})

	RDBReady   = make(chan struct{})
	RedisReady = make(chan struct{})
	MongoReady = make(chan struct{})
)

type App struct {
	RuntimeConfig *config.RuntimeConfig
	RDB           *gorm.DB
	Redis         *redis.Client
	Mongo         *mongo.Client
	Casbin        *casbin.CachedEnforcer
}

func NewApp(ctx context.Context) *App {
	app := &App{}

	app.RuntimeConfig = config.NewRuntimeConfig()

	go func() {
		app.RDB = persistence.NewRDB(app.RuntimeConfig.RDB)
		close(RDBReady)
		app.Casbin = rbac.NewCasbin(app.RDB)
		// Handling flags
		if app.RuntimeConfig.Flags.AutoMigrate {
			utils.MigrateToRDB(app.RDB)
			os.Exit(0)
		}
	}()

	go func() {
		app.Redis = persistence.NewRedisClient(ctx, app.RuntimeConfig.Redis)
		close(RedisReady)
	}()
	go func() {
		app.Mongo = persistence.NewMongoClient(ctx, app.RuntimeConfig.Mongo)
		close(MongoReady)
	}()

	return app
}

func (app *App) CloseConnection(ctx context.Context) error {
	// Note: For RDB, defer sqlDB.Close() on-demand

	var wg sync.WaitGroup
	errChan := make(chan error, 2)
	var errs []error

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := persistence.CloseMongoConnection(app.Mongo, ctx); err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := persistence.CloseRedisConnection(app.Redis); err != nil {
			errChan <- err
		}
	}()

	wg.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("multiple errors occurred: %v", errs)
	}
	zap.S().Info("All connections closed")
	return nil
}

func (app *App) Start(ctx context.Context) {
	if app.RuntimeConfig.Gin.TLS {
		certFile := "./ssl/gin/gin.pem"
		keyFile := "./ssl/gin/gin-key.pem"

		if app.RuntimeConfig.Gin.MTLS {
			// TODO: MTLS
		}

		// Wait for dependencies
		<-RDBReady
		//<-RedisReady
		//<-MongoReady

		g := gin.Default()
		route.Setup(app.RuntimeConfig, app.RDB, g)
		err := g.RunTLS(app.RuntimeConfig.Gin.Addr(), certFile, keyFile) // Blocked
		if err != nil {
			zap.S().Fatalf("Failed to start Gin server: %v", err)
		}
	}
	g := gin.Default()
	route.Setup(app.RuntimeConfig, app.RDB, g)
	err := g.Run(app.RuntimeConfig.Gin.Addr()) // Blocked
	if err != nil {
		zap.S().Fatalf("Failed to start Gin server: %v", err)
	}
}

func (app *App) Shutdown(cancel context.CancelFunc) {
	sig := <-SigChan
	zap.S().Infof("Received signal: %v. Shutting down...", sig)
	cancel() // Call cancel to signal all goroutines to stop

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.CloseConnection(shutdownCtx); err != nil {
		zap.S().Errorf("Error closing connection: %v", err)
	}

	close(ShutDownChan) // Notify main go routine
}

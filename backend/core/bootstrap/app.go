package bootstrap

import (
	"context"
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/database"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/rbac"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"sync"
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
		app.RDB = database.NewRDB(app.RuntimeConfig.RDB)
		app.Casbin = rbac.NewCasbin(app.RDB)
		// Handling flags
		if app.RuntimeConfig.Flags.AutoMigrate {
			utils.MigrateToRDB(app.RDB)
			os.Exit(0)
		}
	}()

	go func() { app.Redis = database.NewRedisClient(ctx, app.RuntimeConfig.Redis) }()
	go func() { app.Mongo = database.NewMongoClient(ctx, app.RuntimeConfig.Mongo) }()

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
		if err := database.CloseMongoConnection(app.Mongo, ctx); err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := database.CloseRedisConnection(app.Redis); err != nil {
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

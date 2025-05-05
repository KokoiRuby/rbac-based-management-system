package main

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/route"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/bootstrap"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/lifecycle"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := logging.InitLogger()
	if err != nil {
		panic(err)
	}

	app := bootstrap.NewApp(ctx)

	signal.Notify(lifecycle.SigChan, syscall.SIGINT, syscall.SIGTERM)
	go lifecycle.GracefulShutdown(app, cancel)

	go func() {
		if app.RuntimeConfig.Gin.TLS {
			if app.RuntimeConfig.Gin.MTLS {
				// TODO
			}
			certFile := "./tls/certs/gin.pem"
			keyFile := "./tls/certs/gin-key.pem"

			g := gin.Default()
			route.Setup(app, g)
			err := g.RunTLS(app.RuntimeConfig.Gin.Addr(), certFile, keyFile) // Blocked
			if err != nil {
				zap.S().Fatalf("Failed to start Gin server: %v", err)
				return
			}
		}
		g := gin.Default()
		route.Setup(app, g)
		err = g.Run(app.RuntimeConfig.Gin.Addr()) // Blocked
		if err != nil {
			zap.S().Fatalf("Failed to start Gin server: %v", err)
			return
		}
	}()

	zap.S().Info("Program running. Press Ctrl+C to trigger graceful shutdown...")
	select {
	case <-lifecycle.ShutDownChan:
		zap.S().Info("Graceful shutdown completed.")
		os.Exit(0)
	case <-time.After(5 * time.Minute):
		zap.S().Info("Program timeout, exiting...")
	}
}

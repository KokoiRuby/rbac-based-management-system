package main

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/bootstrap"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/lifecycle"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/logging"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	wg        sync.WaitGroup
	Readiness sync.Map
)

func main() {
	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := logging.InitLogger()
	if err != nil {
		panic(err)
	}

	app := bootstrap.NewApp(ctx)

	// Readiness probe
	go func() {
		for {
			if lifecycle.IsReady() {
				zap.S().Debug("Ready")
			} else {
				zap.S().Debug("Not Ready")
			}
			time.Sleep(3 * time.Second)
		}
	}()

	// Graceful shutdown
	signal.Notify(lifecycle.SigChan, syscall.SIGINT, syscall.SIGTERM)
	go lifecycle.GracefulShutdown(app, cancel)

	zap.S().Info("Program running. Press Ctrl+C to trigger graceful shutdown...")
	select {
	case <-lifecycle.ShutDownChan:
		zap.S().Info("Graceful shutdown completed.")
		os.Exit(0)
	case <-time.After(30 * time.Second):
		zap.S().Info("Program timeout, exiting...")
	}
}

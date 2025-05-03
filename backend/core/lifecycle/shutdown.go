package lifecycle

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/bootstrap"
	"go.uber.org/zap"
	"os"
	"time"
)

var (
	SigChan      = make(chan os.Signal, 1)
	ShutDownChan = make(chan struct{})
)

func GracefulShutdown(app *bootstrap.App, cancel context.CancelFunc) {

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

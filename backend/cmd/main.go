package main

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/logging"
	"go.uber.org/zap"
)

func main() {

	config.ParseAndLoad() // Configuration

	logging.Init() // Logging

	zap.L().Debug("DEBUG")
	zap.L().Info("INFO")

	select {}

}

package database

import (
	"context"
	"database/sql"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

func NewRDB(cfg runtime.RDBConfig) *gorm.DB {

	dialector := cfg.GetDSN()
	if dialector == nil {
		zap.S().Fatal("database dialector is nil")
	}

	// Open initialize db session based on dialector
	db, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		zap.S().Fatalf("failed to initialize database connection: %v", err)
	}

	// Get DB connection pool
	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Fatalf("failed to get sql database: %v", err)
	}
	// Ping to verify initial connection
	err = sqlDB.Ping()
	if err != nil {
		zap.S().Fatalf("failed to ping database: %v", err)
	}
	global.Readiness.Store("rdb", true)

	// Connection Pool
	// https://gorm.io/docs/generic_interface.html#Connection-Pool
	// TODO: To config
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Start a goroutine for heartbeat
	// TODO: Backoff by factor
	go heartBeatRDB(context.Background(), sqlDB, cfg)

	zap.S().Infof("Connected to RDB successfully")
	return db
}

func heartBeatRDB(ctx context.Context, db *sql.DB, cfg runtime.RDBConfig) {
	interval := cfg.PingInterval
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	zap.S().Debugf("Starting heartbeat to RDB with interval %vs", interval)

	for {
		select {
		case <-ticker.C:
			err := db.Ping()
			if err != nil {
				zap.S().Warnf("Failed to heartbeat RDB: %v", err)
				global.Readiness.CompareAndSwap("rdb", true, false)
			} else {
				global.Readiness.CompareAndSwap("rdb", false, true)
			}
		case <-ctx.Done():
			zap.S().Debugf("Stopping heartbeat to RDB due to context cancellation")
			return
		}
	}
}

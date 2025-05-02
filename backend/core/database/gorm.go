package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

func NewGormDB(cfg runtime.RDBConfig) (*gorm.DB, error) {

	dialector := cfg.GetDSN()
	if dialector == nil {
		return nil, errors.New("failed to get dialector")
	}

	// Open initialize db session based on dialector
	db, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	// Get DB connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// Ping to verify initial connection
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

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
	return db, nil
}

func heartBeatRDB(ctx context.Context, db *sql.DB, cfg runtime.RDBConfig) {
	interval := cfg.PingInterval
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	zap.S().Debugf("Starting heartbeat to RDB with interval %v", interval)

	for {
		select {
		case <-ticker.C:
			err := db.Ping()
			if err != nil {
				zap.S().Warnf("Failed to heartbeat RDB: %v", err)
				global.Readiness["rdb"] = false
			}
			//zap.S().Debugf("Heartbeat to Redis successful")
			global.Readiness["rdb"] = true

		case <-ctx.Done():
			zap.S().Debugf("Stopping heartbeat to RDB due to context cancellation")
			return
		}
	}
}

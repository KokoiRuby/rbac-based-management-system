package database

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

func NewRedisClient(cfg runtime.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Pass,
		DB:       cfg.DB,
	})

	// Ping to verify initial connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.PingTimeout)*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	global.Readiness.Store("redis", true)

	// Start a goroutine for heartbeat
	// TODO: Backoff by factor
	go heartBeatRedis(context.Background(), client, cfg)

	zap.S().Info("Connected to Redis successfully")
	return client, nil
}

func heartBeatRedis(ctx context.Context, client *redis.Client, cfg runtime.RedisConfig) {
	interval := cfg.PingInterval
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	zap.S().Debugf("Starting heartbeat to Redis with interval %v", interval)

	for {
		select {
		case <-ticker.C:
			pingCtx, cancel := context.WithTimeout(ctx, time.Duration(cfg.PingTimeout)*time.Second)
			defer cancel()

			_, err := client.Ping(pingCtx).Result()
			if err != nil {
				zap.S().Warnf("Failed to heartbeat Redis: %v", err)
				global.Readiness.Swap("redis", false)
			}
			//zap.S().Debugf("Heartbeat to Redis successful")
			global.Readiness.CompareAndSwap("redis", false, true)

		case <-ctx.Done():
			zap.S().Debugf("Stopping heartbeat to Redis due to context cancellation")
			return
		}
	}
}

package persistence

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

func NewRedisClient(ctx context.Context, cfg runtime.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Pass,
		DB:       cfg.DB,
	})

	// Ping to verify initial connection
	pingCtx, cancel := context.WithTimeout(ctx, time.Duration(cfg.PingTimeout)*time.Second)
	defer cancel()

	_, err := client.Ping(pingCtx).Result()
	if err != nil {
		zap.S().Warnf("Failed to ping Redis: %v", err)
	}
	global.Readiness.Store("redis", true)

	// Start a goroutine for heartbeat
	// TODO: Backoff by factor
	go heartBeatRedis(ctx, client, cfg)

	zap.S().Info("Connected to Redis successfully")
	return client
}

func heartBeatRedis(ctx context.Context, client *redis.Client, cfg runtime.RedisConfig) {
	interval := cfg.PingInterval
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	zap.S().Debugf("Starting heartbeat to Redis with interval %vs", interval)

	for {
		select {
		case <-ticker.C:
			pingCtx, cancel := context.WithTimeout(ctx, time.Duration(cfg.PingTimeout)*time.Second)
			defer cancel()

			_, err := client.Ping(pingCtx).Result()
			if err != nil {
				zap.S().Warnf("Failed to heartbeat Redis: %v", err)
				global.Readiness.CompareAndSwap("redis", true, false)
			} else {
				global.Readiness.CompareAndSwap("redis", false, true)
			}
		case <-ctx.Done():
			zap.S().Debugf("Stopping heartbeat to Redis due to context cancellation")
			global.Readiness.CompareAndSwap("redis", true, false)
			return
		}
	}
}

func CloseRedisConnection(client *redis.Client) error {
	if client != nil {
		if err := client.Close(); err != nil {
			return err
		} else {
			zap.S().Info("Redis connection closed")
			return nil
		}
	} else {
		zap.S().Warn("Redis client was nil, skipping close.")
		return nil
	}
}

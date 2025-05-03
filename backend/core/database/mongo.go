package database

import (
	"context"
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

func NewMongoClient(cfg runtime.MongoConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ConnectionTimeout)*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s",
		cfg.Pass,
		cfg.Pass,
		cfg.Host,
		cfg.App,
	)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Ping to verify initial connection
	pingCtx, cancel := context.WithTimeout(ctx, time.Duration(cfg.PingTimeout)*time.Second)
	defer cancel()
	err = client.Ping(pingCtx, nil)
	if err != nil {
		return nil, err
	}
	global.Readiness.Store("mongo", true)

	// Start a goroutine for heartbeat
	// TODO: Backoff by factor
	go heartBeatMongo(context.Background(), client, cfg)

	zap.S().Info("Connected to MongoDB successfully")
	return client, nil
}

func heartBeatMongo(ctx context.Context, client *mongo.Client, cfg runtime.MongoConfig) {
	interval := cfg.PingInterval
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	zap.S().Debugf("Starting heartbeat to MongoDB with interval %v", interval)

	for {
		select {
		case <-ticker.C:
			pingCtx, cancel := context.WithTimeout(ctx, time.Duration(cfg.PingTimeout)*time.Second)
			defer cancel()

			err := client.Ping(pingCtx, nil)
			if err != nil {
				zap.S().Warnf("Failed to heartbeat MongoDB: %v", err)
				global.Readiness.Swap("mongo", false)
			}
			//zap.S().Debugf("Heartbeat to Redis successful")
			global.Readiness.CompareAndSwap("mongo", false, true)

		case <-ctx.Done():
			zap.S().Debugf("Stopping heartbeat to MongoDB due to context cancellation")
			return
		}
	}
}

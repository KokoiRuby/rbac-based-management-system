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

func NewMongoClient(ctx context.Context, cfg runtime.MongoConfig) *mongo.Client {
	connectCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ConnectionTimeout)*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s",
		cfg.Pass,
		cfg.Pass,
		cfg.Host,
		cfg.App,
	)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(connectCtx, opts)
	if err != nil {
		zap.S().Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Ping to verify initial connection
	pingCtx, cancel := context.WithTimeout(ctx, time.Duration(cfg.PingTimeout)*time.Second)
	defer cancel()
	err = client.Ping(pingCtx, nil)
	if err != nil {
		zap.S().Fatalf("failed to ping MongoDB: %v", err)
	}
	global.Readiness.Store("mongo", true)

	// Start a goroutine for heartbeat
	// TODO: Backoff by factor
	go heartBeatMongo(ctx, client, cfg)

	zap.S().Info("Connected to MongoDB successfully")
	return client
}

func heartBeatMongo(ctx context.Context, client *mongo.Client, cfg runtime.MongoConfig) {
	interval := cfg.PingInterval
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	zap.S().Debugf("Starting heartbeat to MongoDB with interval %vs", interval)

	for {
		select {
		case <-ticker.C:
			pingCtx, cancel := context.WithTimeout(ctx, time.Duration(cfg.PingTimeout)*time.Second)
			defer cancel()

			err := client.Ping(pingCtx, nil)
			if err != nil {
				zap.S().Warnf("Failed to heartbeat MongoDB: %v", err)
				global.Readiness.CompareAndSwap("mongo", true, false)
			} else {
				global.Readiness.CompareAndSwap("mongo", false, true)
			}
		case <-ctx.Done():
			zap.S().Debugf("Stopping heartbeat to MongoDB due to context cancellation")
			return
		}
	}
}

func CloseMongoConnection(client *mongo.Client, ctx context.Context) error {
	if client != nil {
		if err := client.Disconnect(ctx); err != nil {
			return err
		} else {
			zap.S().Info("MongoDB connection closed")
			return nil
		}
	} else {
		zap.S().Warn("MongoDB client was nil, skipping close.")
		return nil
	}
}

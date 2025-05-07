package repository

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) service.RedisCache {
	return &redisCache{client: client}
}

func (r redisCache) SetKeyWithTTL(c context.Context, key string, value string, ttl time.Duration) (string, error) {
	return r.client.Set(c, key, value, ttl).Result()
}

func (r redisCache) IsKeyExist(c context.Context, key string) (bool, error) {
	exists, err := r.client.Exists(c, key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (r redisCache) DelKey(c context.Context, key string) error {
	_, err := r.client.Del(c, key).Result()
	if err != nil {
		return err
	}
	return nil
}

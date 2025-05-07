package service

import (
	"context"
	"time"
)

type RedisCache interface {
	SetKeyWithTTL(c context.Context, key string, value string, ttl time.Duration) (string, error)
	IsKeyExist(c context.Context, key string) (bool, error)
}

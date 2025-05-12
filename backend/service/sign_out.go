package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type signoutService struct {
	cache          service.RedisCache
	contextTimeout time.Duration
}

func NewSignoutService(cache service.RedisCache, duration time.Duration) service.SignoutService {
	return &signoutService{
		cache:          cache,
		contextTimeout: duration,
	}
}

func (s signoutService) ExtractExpireAtFromToken(tokenString string) (*jwt.NumericDate, error) {
	return utils.ExtractExpireAtFromToken(tokenString)
}

func (s signoutService) SetKeyWithTTLToCache(c *gin.Context, key string, value string, ttl time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.SetKeyWithTTL(ctx, key, value, ttl)
}

func (s signoutService) IsSignedOut(c *gin.Context, key string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.IsKeyExist(ctx, key)
}

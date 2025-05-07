package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type SignoutService interface {
	ExtractExpireAtFromToken(tokenString string, cfg runtime.JWT) (*jwt.NumericDate, error)
	SetKeyWithTTLToCache(c *gin.Context, key string, value string, ttl time.Duration) (string, error)
	IsSignedOut(c *gin.Context, key string) (bool, error)
}

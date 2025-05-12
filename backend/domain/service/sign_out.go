package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type SignoutService interface {
	ExtractExpireAtFromToken(tokenString string) (*jwt.NumericDate, error)
	SetKeyWithTTLToCache(c *gin.Context, key string, value string, ttl time.Duration) (string, error)
	IsSignedOut(c *gin.Context, key string) (bool, error)
}

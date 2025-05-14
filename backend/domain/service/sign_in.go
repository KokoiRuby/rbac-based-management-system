package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
	"time"
)

type SigninService interface {
	GetUserByEmail(c *gin.Context, email string) (*model.User, error)
	CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error)
	CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error)
	IsKeyExist(c *gin.Context, key string) (bool, error)
	FlagSignin(c *gin.Context, key string, value string, ttl time.Duration) (string, error)
	UnFlagSignout(c *gin.Context, key string) error
}

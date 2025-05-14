package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type signinService struct {
	userRDB        service.UserRDB
	cache          service.RedisCache
	contextTimeout time.Duration
}

func NewSigninService(rdb service.UserRDB, cache service.RedisCache, timeout time.Duration) service.SigninService {
	return &signinService{
		userRDB:        rdb,
		cache:          cache,
		contextTimeout: timeout,
	}
}

func (s signinService) GetUserByEmail(c *gin.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.userRDB.GetByEmail(ctx, email)
}

func (s signinService) CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error) {
	return utils.CreateAccessToken(user, cfg)
}

func (s signinService) CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, cfg)
}

func (s signinService) UnFlagSignout(c *gin.Context, key string) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.DelKey(ctx, key)
}

func (s signinService) FlagSignin(c *gin.Context, key string, value string, ttl time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.SetKeyWithTTL(ctx, key, value, ttl)
}

func (s signinService) IsKeyExist(c *gin.Context, key string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.IsKeyExist(ctx, key)
}

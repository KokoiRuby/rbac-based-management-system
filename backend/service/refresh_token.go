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

type refreshTokenService struct {
	userRDB        service.UserRDB
	contextTimeout time.Duration
}

func NewRefreshTokenService(rdb service.UserRDB, timeout time.Duration) service.RefreshTokenService {
	return &refreshTokenService{
		userRDB:        rdb,
		contextTimeout: timeout,
	}
}

func (s refreshTokenService) ExtractIDFromToken(token string, cfg runtime.JWT) (uint, error) {
	return utils.ExtractIDFromToken(token, cfg)
}

func (s refreshTokenService) GetUserByID(c *gin.Context, id uint) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.userRDB.GetByID(ctx, id)
}

func (s refreshTokenService) CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error) {
	return utils.CreateAccessToken(user, cfg)
}

func (s refreshTokenService) CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, cfg)
}

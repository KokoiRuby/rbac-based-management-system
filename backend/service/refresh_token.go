package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
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

func (s refreshTokenService) ExtractIDFromToken(requestToken string, cfg runtime.JWT) (uint, error) {
	return utils.ExtractIDFromToken(requestToken, cfg)
}

func (s refreshTokenService) GetUserByID(c context.Context, id uint) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userRDB.GetByID(ctx, id)
}

func (r refreshTokenService) CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error) {
	return utils.CreateAccessToken(user, cfg)
}

func (r refreshTokenService) CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, cfg)
}

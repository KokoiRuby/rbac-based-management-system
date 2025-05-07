package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"time"
)

type signinService struct {
	userRDB        service.UserRDB
	contextTimeout time.Duration
}

func NewSigninService(rdb service.UserRDB, timeout time.Duration) service.SigninService {
	return &signinService{
		userRDB:        rdb,
		contextTimeout: timeout,
	}
}

func (s signinService) GetUserByEmail(c context.Context, email string) (model.User, error) {
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

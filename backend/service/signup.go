package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"time"
)

type signupService struct {
	userRDB        service.UserRDB
	contextTimeout time.Duration
}

func NewSignupService(rdb service.UserRDB, timeout time.Duration) service.SignupService {
	return &signupService{
		userRDB:        rdb,
		contextTimeout: timeout,
	}
}

func (s signupService) Create(c context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userRDB.Create(ctx, user)
}

func (s signupService) GetUserByEmail(c context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userRDB.GetByEmail(ctx, email)
}

func (s signupService) CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error) {
	return utils.CreateAccessToken(user, cfg)
}

func (s signupService) CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, cfg)
}

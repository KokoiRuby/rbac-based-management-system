package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"gopkg.in/gomail.v2"
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

func (s signupService) GetUserByEmail(c context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.userRDB.GetByEmail(ctx, email)
}

func (s signupService) CreateConfirmToken(req *model.SignupConfirmRequest, cfg runtime.JWT) (confirmToken string, err error) {
	return utils.CreateConfirmToken(req, cfg)
}

func (s signupService) SendConfirmEmail(msg *gomail.Message, cfg runtime.SMTPConfig) error {
	return utils.SendEmail(msg, cfg)
}

func (s signupService) Confirm(token string, cfg runtime.JWT) (*model.SignupRequest, error) {
	return utils.ExtractCredFromToken(token, cfg)
}

func (s signupService) CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error) {
	return utils.CreateAccessToken(user, cfg)
}

func (s signupService) Create(c context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.userRDB.Create(ctx, user)
}

func (s signupService) CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, cfg)
}

package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"gopkg.in/gomail.v2"
)

type SignupService interface {
	Create(c context.Context, user *model.User) error
	GetUserByEmail(c context.Context, email string) (model.User, error)
	SendConfirmEmail(msg *gomail.Message, cfg runtime.SMTPConfig) error
	CreateConfirmToken(user *model.SignupConfirmRequest, cfg runtime.JWT) (confirmToken string, err error)
	Confirm(token string, cfg runtime.JWT) (*model.SignupRequest, error)
	CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error)
	CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error)
}

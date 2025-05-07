package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"time"
)

type SignupService interface {
	Create(c context.Context, user *model.User) error
	GetUserByEmail(c context.Context, email string) (model.User, error)
	SendConfirmEmail(msg *gomail.Message, cfg runtime.SMTPConfig) error
	CreateConfirmToken(user *model.SignupConfirmRequest, cfg runtime.JWT) (confirmToken string, err error)
	Confirm(token string, cfg runtime.JWT) (*model.SignupRequest, error)
	CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error)
	CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error)
	SetKeyWithTTLToCache(c *gin.Context, key string, value string, ttl time.Duration) (string, error)
	DelKeyFromCache(c *gin.Context, key string) error
	IsKeyExist(c context.Context, key string) (bool, error)
}

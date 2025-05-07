package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"time"
)

type ForgotPasswordService interface {
	GetUserByEmail(c *gin.Context, email string) (model.User, error)
	CreateConfirmToken(user *model.ForgotPasswordRequest, cfg runtime.JWT) (confirmToken string, err error)
	SendConfirmEmail(msg *gomail.Message, cfg runtime.SMTPConfig) error
	Confirm(token string, cfg runtime.JWT) (string, error)
	SetKeyWithTTLToCache(c *gin.Context, key string, value string, ttl time.Duration) (string, error)
	DelKeyFromCache(c *gin.Context, key string) error
	IsKeyExist(c *gin.Context, key string) (bool, error)
	UpdateUser(c *gin.Context, user *model.User) error
}

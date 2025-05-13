package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type UpdateUserService interface {
	GetByID(c *gin.Context, id uint) (model.User, error)
	UpdateUser(c *gin.Context, user *model.User) error
	ValidateUserNameUniqueness(c *gin.Context, user *model.User, username string) (bool, error)
	ValidateEmailUniqueness(c *gin.Context, user *model.User, email string) (bool, error)
	CreateConfirmToken(req *model.UserUpdateConfirmRequest, cfg runtime.JWT) (confirmToken string, err error)
	SendConfirmEmail(msg *gomail.Message, cfg runtime.SMTPConfig) error
	Confirm(token string) (*model.UserUpdateConfirmRequest, error)
}

package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
)

type UpdateUserService interface {
	GetByID(c *gin.Context, id uint) (model.User, error)
	UpdateUser(c *gin.Context, user *model.User) error
	ValidateUserNameUniqueness(c *gin.Context, user *model.User, username string) (bool, error)
	ValidateEmailUniqueness(c *gin.Context, user *model.User, email string) (bool, error)
}

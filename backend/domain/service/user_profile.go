package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
)

type UserProfileService interface {
	GetUserByID(c *gin.Context, id uint) (model.User, error)
}

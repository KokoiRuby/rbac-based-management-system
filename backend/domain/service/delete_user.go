package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
)

type DeleteUserService interface {
	GetUserByID(c *gin.Context, id uint) (*model.User, error)
	DeleteByID(c *gin.Context, id uint) (int64, error)
	DeleteUsers(c *gin.Context, ids []uint) (int64, error)
}

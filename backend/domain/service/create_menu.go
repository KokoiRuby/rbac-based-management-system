package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
)

type CreateMenuService interface {
	CreateMenu(c *gin.Context, user *model.Menu) error
	GetMenuByID(c *gin.Context, id uint) (*model.Menu, error)
	GetMenuByName(c *gin.Context, name string) (*model.Menu, error)
}

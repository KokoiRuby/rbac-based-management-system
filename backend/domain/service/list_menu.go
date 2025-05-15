package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
)

type ListMenusService interface {
	ListMenus(c *gin.Context, opt model.QueryOptions) ([]*model.Menu, int64, error)
}

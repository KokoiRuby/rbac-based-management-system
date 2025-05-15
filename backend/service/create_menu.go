package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/gin-gonic/gin"
	"time"
)

type createMenuService struct {
	rdb            service.MenuRDB
	contextTimeout time.Duration
}

func NewCreateMenuService(rdb service.MenuRDB, contextTimeout time.Duration) service.CreateMenuService {
	return &createMenuService{
		rdb:            rdb,
		contextTimeout: contextTimeout,
	}
}

func (s createMenuService) CreateMenu(c *gin.Context, menu *model.Menu) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.Create(ctx, menu)
}

func (s createMenuService) GetMenuByName(c *gin.Context, name string) (*model.Menu, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByName(ctx, name)
}

func (s createMenuService) GetMenuByID(c *gin.Context, id uint) (*model.Menu, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByID(ctx, id)
}

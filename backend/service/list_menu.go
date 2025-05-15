package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/gin-gonic/gin"
	"time"
)

type listMenusService struct {
	rdb            service.MenuRDB
	contextTimeout time.Duration
}

func NewListMenusService(rdb service.MenuRDB, contextTimeout time.Duration) service.ListMenusService {
	return &listMenusService{
		rdb:            rdb,
		contextTimeout: contextTimeout,
	}
}

func (s listMenusService) ListMenus(c *gin.Context, opt model.QueryOptions) ([]*model.Menu, int64, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.ListByCond(ctx, opt)
}

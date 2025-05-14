package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/gin-gonic/gin"
	"time"
)

type listUsersService struct {
	rdb            service.UserRDB
	contextTimeout time.Duration
}

func NewListUsersService(rdb service.UserRDB, contextTimeout time.Duration) service.ListUsersService {
	return &listUsersService{
		rdb:            rdb,
		contextTimeout: contextTimeout,
	}
}

func (s listUsersService) ListUsers(c *gin.Context, opt model.QueryOptions) ([]*model.User, int64, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.ListByCond(ctx, opt)
}

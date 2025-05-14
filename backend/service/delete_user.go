package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/gin-gonic/gin"
	"time"
)

type deleteUserService struct {
	rdb            service.UserRDB
	contextTimeout time.Duration
}

func NewDeleteUserService(rdb service.UserRDB, contextTimeout time.Duration) service.DeleteUserService {
	return &deleteUserService{
		rdb:            rdb,
		contextTimeout: contextTimeout,
	}
}

func (s deleteUserService) GetUserByID(c *gin.Context, id uint) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByID(ctx, id)
}

func (s deleteUserService) DeleteByID(c *gin.Context, id uint) (int64, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.DeleteByID(ctx, id)
}

func (s deleteUserService) DeleteUsers(c *gin.Context, ids []uint) (int64, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.DeleteByIDs(ctx, ids)
}

package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/gin-gonic/gin"
	"time"
)

type closeUserService struct {
	rdb            service.UserRDB
	contextTimeout time.Duration
}

func NewCloseUserService(rdb service.UserRDB, contextTimeout time.Duration) service.CloseUserService {
	return &closeUserService{
		rdb:            rdb,
		contextTimeout: contextTimeout,
	}
}

func (s closeUserService) GetUserByID(c *gin.Context, id uint) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByID(ctx, id)
}

func (s closeUserService) DeleteByID(c *gin.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.DeleteByID(ctx, id)
}

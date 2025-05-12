package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/gin-gonic/gin"
	"time"
)

type resetPasswordService struct {
	rdb            service.UserRDB
	contextTimeout time.Duration
}

func NewResetPasswordService(rdb service.UserRDB, timeout time.Duration) service.ResetPasswordService {
	return &resetPasswordService{
		rdb:            rdb,
		contextTimeout: timeout,
	}
}

func (s resetPasswordService) GetByID(c context.Context, id uint) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByID(ctx, id)
}

func (s resetPasswordService) UpdateUser(c *gin.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.Update(ctx, user)
}

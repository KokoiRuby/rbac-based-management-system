package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/gin-gonic/gin"
	"time"
)

type userProfileService struct {
	rdb            service.UserRDB
	contextTimeout time.Duration
}

func NewUserProfileService(rdb service.UserRDB, contextTimeout time.Duration) service.UserProfileService {
	return &updateUserService{
		rdb:            rdb,
		contextTimeout: contextTimeout,
	}
}

func (s userProfileService) GetUserByID(c *gin.Context, id uint) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByID(ctx, id)
}

package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
)

type ListUsersService interface {
	ListUsers(c *gin.Context, opt model.QueryOptions) ([]*model.User, int64, error)
}

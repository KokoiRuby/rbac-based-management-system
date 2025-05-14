package handler

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type ListUsersHandler struct {
	ListUsersService service.ListUsersService
	RuntimeConfig    *config.RuntimeConfig
}

func (handler *ListUsersHandler) ListUsers(c *gin.Context) {
	req := middleware.GetBind[model.UserListRequest](c)

	likes := map[string]any{
		"username": req.Username,
		"email":    req.Email,
		"role":     req.Role,
	}

	opt := model.QueryOptions{
		Pagination: req.Pagination,
		Likes:      likes,
	}

	users, cnt, err := handler.ListUsersService.ListUsers(c, opt)
	if err != nil {
		zap.S().Errorf("failed to list users: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to list users.")
		return
	}

	utils.OKWithList(c, http.StatusOK, cnt, users)
	return

}

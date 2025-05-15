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

type ListMenusHandler struct {
	ListMenusService service.ListMenusService
	RuntimeConfig    *config.RuntimeConfig
}

func (handler *ListMenusHandler) ListMenus(c *gin.Context) {
	req := middleware.GetBind[model.MenuListRequest](c)

	likes := map[string]any{
		"name":  req.Name,
		"title": req.Title,
	}

	opt := model.QueryOptions{
		Pagination: req.Pagination,
		Likes:      likes,
	}

	users, cnt, err := handler.ListMenusService.ListMenus(c, opt)
	if err != nil {
		zap.S().Errorf("failed to list menus: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to list menus.")
		return
	}

	utils.OKWithList(c, http.StatusOK, cnt, users)
	return
}

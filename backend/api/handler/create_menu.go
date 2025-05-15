package handler

import (
	"errors"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type CreateMenuHandler struct {
	CreateMenuService service.CreateMenuService
	RuntimeConfig     *config.RuntimeConfig
}

func (handler *CreateMenuHandler) Create(c *gin.Context) {
	req := middleware.GetBind[model.CreateMenuRequest](c)

	_, err := handler.CreateMenuService.GetMenuByName(c, req.Name)
	if err == nil {
		utils.FailWithMsg(c, http.StatusConflict, "Menu already exists.")
		return
	} else {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			zap.S().Errorf("failed to get menu by name: %v", err)
			utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to create menu.")
			return
		}
	}

	if req.ParentMenuID != nil {
		_, err = handler.CreateMenuService.GetMenuByID(c, *req.ParentMenuID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utils.FailWithMsg(c, http.StatusNotFound, "Parent menu id does not exist.")
				return

			}
			zap.S().Errorf("failed to get parent menu by id: %v", err)
			utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to create menu.")
			return
		}
	}

	menu := &model.Menu{
		Name:         req.Name,
		Path:         req.Path,
		Component:    req.Component,
		ParentMenuID: req.ParentMenuID,
		Sort:         req.Sort,
		Meta: model.Meta{
			Icon:  req.Icon,
			Title: req.Title,
		},
		Enable: req.Enable,
	}

	err = handler.CreateMenuService.CreateMenu(c, menu)
	if err != nil {
		zap.S().Errorf("failed to create menu: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to create menu.")
	}

	utils.OK(c, http.StatusOK, map[string]any{
		"id": menu.ID,
	}, "Create menu successfully.")
	return

}

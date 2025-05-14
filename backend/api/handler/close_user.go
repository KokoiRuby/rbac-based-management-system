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
	"strings"
)

type CloseUserHandler struct {
	CloseUserService service.CloseUserService
	RuntimeConfig    *config.RuntimeConfig
}

func (handler *CloseUserHandler) Close(c *gin.Context) {
	// TODO: Binding won't take effect in HTTP DELETE
	req := middleware.GetBind[model.CloseUserRequest](c)

	// TODO: Left-shift-able to frontend
	if strings.ToLower(req.ConfirmMsg) != "agreed" {
		utils.FailWithMsg(c, http.StatusBadRequest, `Please type "Agreed" to close this user.`)
		return
	}

	claims, ok := c.Get("claims")
	if !ok {
		zap.S().Error("failed to get claims from context")
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to update user.")
		return
	}

	id := claims.(*utils.CustomClaims).UserID
	_, err := handler.CloseUserService.GetUserByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to close user.")
		return
	}

	_, err = handler.CloseUserService.DeleteByID(c, id)
	if err != nil {
		zap.S().Errorf("failed to delete user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to close user.")
		return
	}

	utils.OKWithMsg(c, http.StatusOK, "User closed successfully.")
	return
}

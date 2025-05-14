package handler

import (
	"errors"
	"fmt"
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

type DeleteUserHandler struct {
	DeleteUserService service.DeleteUserService
	RuntimeConfig     *config.RuntimeConfig
}

func (handler *DeleteUserHandler) DeleteAUser(c *gin.Context) {
	req := middleware.GetBind[model.DeleteUserRequest](c)

	// Optional: Get before Delete
	id := req.UserID
	_, err := handler.DeleteUserService.GetUserByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to delete user.")
		return
	}

	_, err = handler.DeleteUserService.DeleteByID(c, id)
	if err != nil {
		zap.S().Errorf("failed to delete user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to delete user.")
		return
	}

	utils.OKWithMsg(c, http.StatusOK, "User deleted successfully.")
	return
}

func (handler *DeleteUserHandler) DeleteUsers(c *gin.Context) {
	req := middleware.GetBind[model.DeleteUsersRequest](c)

	ids := req.UserIDs
	cnt, err := handler.DeleteUserService.DeleteUsers(c, ids)
	if err != nil {
		zap.S().Errorf("failed to delete users by ids: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to delete users.")
		return
	}

	utils.OKWithMsg(c, http.StatusOK, fmt.Sprintf("Deleted %d users successfully.", cnt))
	return

}

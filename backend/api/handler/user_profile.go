package handler

import (
	"errors"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type UserProfileHandler struct {
	UserProfileService service.UserProfileService
	RuntimeConfig      *config.RuntimeConfig
}

func (handler *UserProfileHandler) GetProfile(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		zap.S().Error("failed to get claims from context")
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to get user profile.")
		return
	}

	id := claims.(*utils.CustomClaims).UserID
	user, err := handler.UserProfileService.GetUserByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to get user profile.")
		return
	}

	profile := model.UserProfile{
		UserID:   user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		RoleList: user.GetRoleList(),
	}

	utils.OKWithData(c, http.StatusOK, profile)

}

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

type RefreshTokenHandler struct {
	RefreshTokenService service.RefreshTokenService
	RuntimeConfig       *config.RuntimeConfig
}

func (handler *RefreshTokenHandler) Refresh(c *gin.Context) {
	req := middleware.GetBind[model.RefreshTokenRequest](c)

	id, err := handler.RefreshTokenService.ExtractIDFromToken(req.RefreshToken)
	if err != nil {
		zap.S().Errorf("failed to extract id from refresh token: %v", err)
		utils.FailWithMsg(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	user, err := handler.RefreshTokenService.GetUserByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to refresh token.")
		return
	}

	accessToken, err := handler.RefreshTokenService.CreateAccessToken(
		&user,
		handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create access token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to refresh token.")
	}

	refreshToken, err := handler.RefreshTokenService.CreateRefreshToken(
		&user,
		handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create refresh token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to refresh token.")
	}

	resp := model.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.OKWithData(c, http.StatusOK, resp)
	return
}

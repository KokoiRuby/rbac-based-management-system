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

type SignupHandler struct {
	SignupService service.SignupService
	RuntimeConfig *config.RuntimeConfig
}

func (handler *SignupHandler) Signup(c *gin.Context) {
	req := middleware.GetBind[model.SignupRequest](c)

	_, err := handler.SignupService.GetUserByEmail(c, req.Email)
	if err == nil {
		utils.FailWithMsg(c, http.StatusConflict, "Email already registered.")
		return
	} else {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			zap.S().Errorf("failed to get user by email: %v", err)
			utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
			return
		}
	}

	hashedPassword, err := utils.Encrypt(req.Password)
	if err != nil {
		zap.S().Errorf("failed to encrypt password: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	err = handler.SignupService.Create(c, user)
	if err != nil {
		zap.S().Errorf("failed to create user: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
	}

	accessToken, err := handler.SignupService.CreateAccessToken(
		user,
		handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create access token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
	}

	refreshToken, err := handler.SignupService.CreateRefreshToken(
		user,
		handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create refresh token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
	}

	resp := model.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.OKWithData(c, http.StatusOK, resp)
	return
}

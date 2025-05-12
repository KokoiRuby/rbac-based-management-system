package handler

import (
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type SigninHandler struct {
	SigninService service.SigninService
	RuntimeConfig *config.RuntimeConfig
}

func (handler *SigninHandler) Signin(c *gin.Context) {
	req := middleware.GetBind[model.SigninRequest](c)

	// TODO: SMS OTP or captcha

	user, err := handler.SigninService.GetUserByEmail(c, req.Email)
	if err != nil {
		utils.FailWithMsg(c, http.StatusBadRequest, "Invalid credentials")
		return
	}

	if !utils.Validate(user.Password, req.Password) {
		utils.FailWithMsg(c, http.StatusBadRequest, "Invalid credentials")
		return
	}

	// Unflag signout in cache if signed out already
	key := fmt.Sprintf("signout_%s", req.Email)
	err = handler.SigninService.UnFlagSignout(c, key)
	if err != nil {
		zap.S().Errorf("failed to unflag signout: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signin.")
		return
	}

	accessToken, err := handler.SigninService.CreateAccessToken(
		&user,
		handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create access token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signin.")
	}
	c.Header("Authorization", "Bearer "+accessToken)

	refreshToken, err := handler.SigninService.CreateRefreshToken(
		&user,
		handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create refresh token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signin.")
	}

	resp := model.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	utils.OKWithData(c, http.StatusOK, resp)
	return
}

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

	accessToken, err := handler.SigninService.CreateAccessToken(
		&user,
		handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create access token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
	}

	refreshToken, err := handler.SigninService.CreateRefreshToken(
		&user,
		handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create refresh token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
	}

	resp := model.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.Header("Access-Token", accessToken)
	utils.OKWithData(c, http.StatusOK, resp)
	return
}

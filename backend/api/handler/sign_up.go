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
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type SignupHandler struct {
	SignupService service.SignupService
	RuntimeConfig *config.RuntimeConfig
}

func (handler *SignupHandler) Signup(c *gin.Context) {
	req := middleware.GetBind[model.SignupRequest](c)

	// Check key in cache if signup is ongoing
	key := fmt.Sprintf("signup_%v", req.Email)
	exists, err := handler.SignupService.IsKeyExist(c, key)
	if err != nil {
		zap.S().Errorf("failed to get key from cache: %v", err)
	}
	if exists {
		utils.OKWithMsg(c, http.StatusOK, "Signup is still ongoing.")
		return
	}

	_, err = handler.SignupService.GetUserByEmail(c, req.Email)
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

	confirmReq := &model.SignupConfirmRequest{
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	confirmToken, err := handler.SignupService.CreateConfirmToken(confirmReq, handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create confirm token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
		return
	}

	// TODO: param-ize host
	signupURL := fmt.Sprintf("https://localhost:%v/signup/confirm?token=%v", handler.RuntimeConfig.Gin.Port, confirmToken)
	msg := fmt.Sprintf("Complete your signup process in 5min via this URL: %v", signupURL)
	m := gomail.NewMessage()
	m.SetHeader("From", req.Email)
	m.SetHeader("To", req.Email)
	m.SetHeader("Subject", "RBAC-based Management System Signup")
	m.SetBody("text/html", msg)

	err = handler.SignupService.SendConfirmEmail(m, handler.RuntimeConfig.SMTP)
	if err != nil {
		zap.S().Errorf("failed to send confirm email: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
		return
	}
	zap.S().Debugf("Send confirmation mail to %v successfully.", req.Email)

	// Set key to cache to flag ongoing signup
	_, err = handler.SignupService.SetKeyWithTTLToCache(c, key, "", time.Duration(handler.RuntimeConfig.JWT.ConfirmExpire)*time.Minute)
	if err != nil {
		zap.S().Errorf("failed to set key to cache: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
		return
	}

	utils.OKWithMsg(c, http.StatusOK, "Please go to your mailbox to confirm and finish the sign up process.")
	return
}

func (handler *SignupHandler) SignupConfirm(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		utils.FailWithMsg(c, http.StatusBadRequest, "Missing token.")
		return
	}

	req, err := handler.SignupService.Confirm(tokenString, handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to confirm token: %v", err)
		utils.FailWithMsg(c, http.StatusBadRequest, "Invalid token.")
		return
	}

	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err = handler.SignupService.CreateUser(c, user)
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

	// Delete key in cache to unflag ongoing signup
	key := fmt.Sprintf("signup_%v", req.Email)
	err = handler.SignupService.DelKeyFromCache(c, key)
	if err != nil {
		zap.S().Errorf("failed to delete key from cache: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
		return
	}

	resp := model.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.OKWithData(c, http.StatusOK, resp)
	return
}

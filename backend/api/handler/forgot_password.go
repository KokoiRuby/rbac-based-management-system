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

type ForgotPasswordHandler struct {
	ForgotPasswordService service.ForgotPasswordService
	RuntimeConfig         *config.RuntimeConfig
}

func (handler *ForgotPasswordHandler) Forgot(c *gin.Context) {
	req := middleware.GetBind[model.ForgotPasswordRequest](c)

	// Check key in cache if password reset is ongoing
	key := fmt.Sprintf("reset_%v", req.Email)
	exists, err := handler.ForgotPasswordService.IsKeyExist(c, key)
	if err != nil {
		zap.S().Errorf("failed to get key from cache: %v", err)
	}
	if exists {
		utils.OKWithMsg(c, http.StatusOK, "Password reset is still ongoing.")
		return
	}

	_, err = handler.ForgotPasswordService.GetUserByEmail(c, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by email: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	confirmReq := &model.ForgotPasswordRequest{
		Email: req.Email,
	}

	confirmToken, err := handler.ForgotPasswordService.CreateConfirmToken(confirmReq, handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to create confirm token: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	// TODO: param-ize host
	signupURL := fmt.Sprintf("https://localhost:%v/forgotPassword/confirm?token=%v", handler.RuntimeConfig.Gin.Port, confirmToken)
	msg := fmt.Sprintf("Reset your password in 5min via this URL: %v", signupURL)
	m := gomail.NewMessage()
	m.SetHeader("From", req.Email)
	m.SetHeader("To", req.Email)
	m.SetHeader("Subject", "RBAC-based Management System Reset Password")
	m.SetBody("text/html", msg)

	err = handler.ForgotPasswordService.SendConfirmEmail(m, handler.RuntimeConfig.SMTP)
	if err != nil {
		zap.S().Errorf("failed to send confirm email: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}
	zap.S().Debugf("Send confirmation mail to %v successfully.", req.Email)

	// Set key to cache to flag ongoing password reset
	_, err = handler.ForgotPasswordService.SetKeyWithTTLToCache(c, key, "", time.Duration(handler.RuntimeConfig.JWT.ConfirmExpire)*time.Minute)
	if err != nil {
		zap.S().Errorf("failed to set key to cache: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	utils.OKWithMsg(c, http.StatusOK, "Please go to your mailbox to confirm and reset your password.")
	return
}

func (handler *ForgotPasswordHandler) ForgotConfirm(c *gin.Context) {
	req := middleware.GetBind[model.ForgotPasswordConfirmRequest](c)
	if req.NewPassword != req.NewPasswordConfirm {
		utils.FailWithMsg(c, http.StatusBadRequest, "Passwords do not match.")
		return
	}

	tokenString := c.Query("token")
	if tokenString == "" {
		utils.FailWithMsg(c, http.StatusBadRequest, "Missing token.")
		return
	}

	email, err := handler.ForgotPasswordService.Confirm(tokenString)
	if err != nil {
		zap.S().Errorf("failed to confirm token: %v", err)
		utils.FailWithMsg(c, http.StatusBadRequest, "Invalid token.")
		return
	}

	user, err := handler.ForgotPasswordService.GetUserByEmail(c, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by email: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	hashedPassword, err := utils.Encrypt(req.NewPassword)
	if err != nil {
		zap.S().Errorf("failed to encrypt password: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	user.Password = hashedPassword
	err = handler.ForgotPasswordService.UpdateUser(c, user)
	if err != nil {
		zap.S().Errorf("failed to update user: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	// Delete key in cache to unflag ongoing password reset
	key := fmt.Sprintf("reset_%v", email)
	err = handler.ForgotPasswordService.DelKeyFromCache(c, key)
	if err != nil {
		zap.S().Errorf("failed to delete key from cache: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	utils.OKWithMsg(c, http.StatusOK, "Password reset successfully.")
	return

}

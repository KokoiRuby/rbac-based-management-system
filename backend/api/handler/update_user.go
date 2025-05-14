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
)

type UpdateUserHandler struct {
	UpdateUserService service.UpdateUserService
	RuntimeConfig     *config.RuntimeConfig
}

func (handler *UpdateUserHandler) Update(c *gin.Context) {
	req := middleware.GetBind[model.UserUpdate](c)

	claims, ok := c.Get("claims")
	if !ok {
		zap.S().Error("failed to get claims from context")
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to update user.")
		return
	}

	id := claims.(*utils.CustomClaims).UserID
	user, err := handler.UpdateUserService.GetUserByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to update user.")
		return
	}

	// TODO: How to perform transaction in this layer?

	if req.Nickname != "" && user.Nickname != req.Nickname {
		user.Nickname = req.Nickname
	}

	// Validate uniqueness
	if req.Username != "" && user.Username != req.Username {
		ok, err := handler.UpdateUserService.ValidateUserNameUniqueness(c, &user, req.Username)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				zap.S().Errorf("failed to validate username uniqueness: %v", err)
				utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to update user.")
				return
			}
		}
		if !ok {
			utils.FailWithMsg(c, http.StatusConflict, "Username is already taken.")
			return
		}
		user.Username = req.Username

	}

	// Validate uniqueness
	if req.Email != "" && user.Email != req.Email {
		ok, err := handler.UpdateUserService.ValidateEmailUniqueness(c, &user, req.Email)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				zap.S().Errorf("failed to validate username uniqueness: %v", err)
				utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to update user.")
				return
			}
		}
		if !ok {
			utils.FailWithMsg(c, http.StatusConflict, "Email is already taken.")
			return
		}

		// TODO: Send verification to new email and update only when verified.
		confirmReq := &model.UserUpdateConfirmRequest{
			UserID:   user.ID,
			Username: req.Username,
			Nickname: req.Nickname,
			Email:    req.Email,
		}

		confirmToken, err := handler.UpdateUserService.CreateConfirmToken(confirmReq, handler.RuntimeConfig.JWT)
		if err != nil {
			zap.S().Errorf("failed to create confirm token: %v", err)
			utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
			return
		}

		// TODO: param-ize host
		signupURL := fmt.Sprintf("https://localhost:%v/v1/user/update/confirm?token=%v", handler.RuntimeConfig.Gin.Port, confirmToken)
		msg := fmt.Sprintf("Complete your user update process in 5min via this URL: %v", signupURL)
		m := gomail.NewMessage()
		m.SetHeader("From", user.Email)
		m.SetHeader("To", req.Email)
		m.SetHeader("Subject", "RBAC-based Management System Update email")
		m.SetBody("text/html", msg)

		err = handler.UpdateUserService.SendConfirmEmail(m, handler.RuntimeConfig.SMTP)
		if err != nil {
			zap.S().Errorf("failed to send confirm email: %v", err)
			utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signup.")
			return
		}
		zap.S().Debugf("Send confirmation mail to %v successfully.", req.Email)

		utils.OKWithMsg(c, http.StatusOK, "Please go to your mailbox to confirm and finish the user update process.")
		return
	}

	err = handler.UpdateUserService.UpdateUser(c, &user)
	if err != nil {
		zap.S().Errorf("failed to update user: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to update user.")
		return
	}

	utils.OKWithMsg(c, http.StatusOK, "Update user successfully.")
	return
}

func (handler *UpdateUserHandler) UpdateConfirm(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		utils.FailWithMsg(c, http.StatusBadRequest, "Missing token.")
		return
	}

	req, err := handler.UpdateUserService.Confirm(tokenString)
	if err != nil {
		zap.S().Errorf("failed to confirm token: %v", err)
		utils.FailWithMsg(c, http.StatusBadRequest, "Invalid token.")
		return
	}

	id := req.UserID
	user, err := handler.UpdateUserService.GetUserByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to update user.")
		return
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Username != "" {
		user.Username = req.Username
	}

	user.Email = req.Email

	err = handler.UpdateUserService.UpdateUser(c, &user)
	if err != nil {
		zap.S().Errorf("failed to update user: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to update user.")
	}

	utils.OKWithMsg(c, http.StatusOK, "Update user successfully.")
	return
}

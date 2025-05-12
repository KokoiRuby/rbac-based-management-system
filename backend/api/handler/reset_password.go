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

type ResetPasswordHandler struct {
	ResetPasswordService service.ResetPasswordService
	RuntimeConfig        *config.RuntimeConfig
}

func (handler *ResetPasswordHandler) Reset(c *gin.Context) {
	req := middleware.GetBind[model.ResetPasswordRequest](c)

	claims, ok := c.Get("claims")
	if !ok {
		zap.S().Error("failed to get claims from context")
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	id := claims.(*utils.CustomClaims).UserID
	user, err := handler.ResetPasswordService.GetByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	if !utils.Validate(user.Password, req.OldPassword) {
		utils.FailWithMsg(c, http.StatusBadRequest, "Invalid credentials")
		return
	}

	hashedNewPassword, err := utils.Encrypt(req.NewPassword)
	if err != nil {
		zap.S().Errorf("failed to encrypt new password: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	user.Password = hashedNewPassword
	err = handler.ResetPasswordService.UpdateUser(c, &user)
	if err != nil {
		zap.S().Errorf("failed to update user: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to reset password.")
		return
	}

	// Unset token in header & force to re-login.
	c.Writer.Header().Del("Bearer")

	utils.OKWithMsg(c, http.StatusOK, "Reset password successfully.")
	return
}

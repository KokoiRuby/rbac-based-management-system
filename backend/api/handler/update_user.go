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
	user, err := handler.UpdateUserService.GetByID(c, id)
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

	// Uniqueness check required
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

	// Uniqueness check required
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
		user.Email = req.Email
	}

	// TODO: Send verification to new email and update only when verified.

	err = handler.UpdateUserService.UpdateUser(c, &user)
	if err != nil {
		zap.S().Errorf("failed to update user: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to update user.")
		return
	}

	utils.OKWithMsg(c, http.StatusOK, "Update user successfully.")
	return
}

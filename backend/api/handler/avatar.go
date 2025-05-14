package handler

import (
	"errors"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"path"
	"path/filepath"
)

var exts = map[string]struct{}{
	".jpg": {},
	".png": {},
}

type UploadAvatarHandler struct {
	AvatarService service.AvatarService
	RuntimeConfig *config.RuntimeConfig
}

func (handler *UploadAvatarHandler) Upload(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		zap.S().Error("failed to get claims from context")
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to upload avatar.")
		return
	}

	id := claims.(*utils.CustomClaims).UserID
	email := claims.(*utils.CustomClaims).Email

	// TODO: Left-shift-able to frontend
	fileHeader, err := c.FormFile("file")
	if err != nil {
		utils.FailWithMsg(c, http.StatusBadRequest, "Please select an image file.")
		return
	}

	// TODO: Left-shift-able to frontend
	if fileHeader.Size > handler.RuntimeConfig.Upload.Avatar.Size*1024*1024 {
		utils.FailWithMsg(c, http.StatusBadRequest, "Please select an image file.")
		return
	}

	// TODO: Left-shift-able to frontend
	ext := filepath.Ext(fileHeader.Filename)
	if _, ok := exts[ext]; !ok {
		utils.FailWithMsg(c, http.StatusBadRequest, "Image extension not supported.")
		return
	}

	// TODO: Left-shift-able to frontend
	// TODO: thumbnail before persistence to save bandwidth
	// TODO: thumbnail can also be cached in browser cache

	// To local storage
	baseDir := path.Join("/app/static/uploads", handler.RuntimeConfig.Upload.Avatar.Dir, email)
	err = handler.AvatarService.UploadToLocal(c, baseDir, fileHeader)
	if err != nil {
		zap.S().Errorf("failed to upload avatar to local: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to upload avatar to local.")
		return
	}

	// To S3
	err = handler.AvatarService.UploadToS3(c, handler.RuntimeConfig.AWS, email, fileHeader)
	if err != nil {
		zap.S().Errorf("failed to upload avatar to s3: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to upload avatar to s3.")
		return
	}

	// S3 key to db
	user, err := handler.AvatarService.GetUserByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailWithMsg(c, http.StatusNotFound, "User not found.")
			return
		}
		zap.S().Errorf("failed to get user by id: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to upload avatar.")
		return
	}

	user.Avatar = fileHeader.Filename
	err = handler.AvatarService.UpdateUser(c, user)
	if err != nil {
		zap.S().Errorf("failed to update user: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to upload avatar.")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{}, "Upload avatar successfully.")
}

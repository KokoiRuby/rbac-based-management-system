package handler

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

// Upload TODO: To AWS S3
func (handler *UploadAvatarHandler) Upload(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		zap.S().Error("failed to get claims from context")
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to upload avatar.")
		return
	}
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

	// To local
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

	utils.OK(c, http.StatusOK, gin.H{}, "Upload avatar successfully.")
}

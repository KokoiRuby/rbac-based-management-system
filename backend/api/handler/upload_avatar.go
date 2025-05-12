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
	"strings"
)

var exts = map[string]struct{}{
	".jpg": {},
	".png": {},
}

type UploadAvatarHandler struct {
	UploadAvatarService service.UploadAvatarService
	RuntimeConfig       *config.RuntimeConfig
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
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if _, ok := exts[ext]; !ok {
		utils.FailWithMsg(c, http.StatusBadRequest, "Image extension not supported.")
		return
	}

	dst := path.Join("/app/static/uploads", handler.RuntimeConfig.Upload.Avatar.Dir, email, fileHeader.Filename)
	err = c.SaveUploadedFile(fileHeader, dst)
	if err != nil {
		zap.S().Errorf("failed to save and upload file: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to upload avatar.")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{}, "Upload avatar successfully.")
}

package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type AvatarService interface {
	UploadToLocal(c *gin.Context, baseDir string, fileHeader *multipart.FileHeader) error
	UploadToS3(c *gin.Context, cfg runtime.AWS, dir string, fileHeader *multipart.FileHeader) error
}

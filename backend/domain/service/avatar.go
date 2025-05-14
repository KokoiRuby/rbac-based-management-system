package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type AvatarService interface {
	UploadToLocal(c *gin.Context, baseDir string, fileHeader *multipart.FileHeader) error
	UploadToS3(c *gin.Context, cfg runtime.AWS, dir string, fileHeader *multipart.FileHeader) error
	GetUserByID(c *gin.Context, id uint) (*model.User, error)
	UpdateUser(c *gin.Context, user *model.User) error
}

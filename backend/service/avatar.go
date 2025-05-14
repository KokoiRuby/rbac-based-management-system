package service

import (
	"context"
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type AvatarService struct {
	rdb            service.UserRDB
	objStore       service.AvatarObjectStorage
	contextTimeout time.Duration
}

func NewAvatarService(rdb service.UserRDB, objStore service.AvatarObjectStorage, timeout time.Duration) service.AvatarService {
	return &AvatarService{
		rdb:            rdb,
		objStore:       objStore,
		contextTimeout: timeout,
	}
}

func (s AvatarService) UploadToS3(c *gin.Context, cfg runtime.AWS, dir string, fileHeader *multipart.FileHeader) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()

	toUpload, err := fileHeader.Open()
	if err != nil {
		zap.S().Errorf("failed to open file header: %v", err)
		return err
	}

	key := path.Join(dir, fileHeader.Filename)
	return s.objStore.Put(ctx, cfg.S3.Bucket, key, toUpload)
}

// UploadToLocal TODO: Shift to handler layer?
func (s AvatarService) UploadToLocal(c *gin.Context, baseDir string, fileHeader *multipart.FileHeader) error {

	originalFilename := fileHeader.Filename
	ext := filepath.Ext(originalFilename)
	baseName := strings.TrimSuffix(originalFilename, ext)
	dir := path.Join(baseDir, originalFilename)

	// De-duplicate
	counter := 0
	for {
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			break
		}

		toUpload, err := fileHeader.Open()
		if err != nil {
			zap.S().Errorf("failed to open file header: %v", err)
			return err
		}
		toUploadHash := utils.GetMD5(toUpload)

		// sudo chown -R $USER:$USER ./static/uploads/avatars
		existed, err := os.Open(dir)
		if err != nil {
			zap.S().Errorf("failed to open file: %v", err)
			return err
		}
		existedHash := utils.GetMD5(existed)

		if existedHash == toUploadHash {
			return nil
		}

		counter++
		newFilename := fmt.Sprintf("%s (%d)%s", baseName, counter, ext)
		dir = path.Join(baseDir, newFilename)
	}

	err := c.SaveUploadedFile(fileHeader, dir)
	if err != nil {
		zap.S().Errorf("failed to save and upload file: %v", err)
		return err
	}

	return nil
}

func (s AvatarService) GetUserByID(c *gin.Context, id uint) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByID(ctx, id)
}

func (s AvatarService) UpdateUser(c *gin.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.Update(ctx, user)
}

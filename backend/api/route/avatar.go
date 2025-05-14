package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/handler"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/infra/repository"
	"github.com/KokoiRuby/rbac-based-management-system/backend/service"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func NewUploadAvatarRouter(cfg *config.RuntimeConfig, db *gorm.DB, objStore *s3.Client, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db)
	avatarRepo := repository.NewAvatarRepository(objStore)

	h := handler.UploadAvatarHandler{
		AvatarService: service.NewAvatarService(userRepo, avatarRepo, time.Duration(cfg.Gin.Timeout)),
		RuntimeConfig: cfg,
	}
	group.POST("/upload/avatar",
		middleware.AuthNMiddleware,
		h.Upload)
}

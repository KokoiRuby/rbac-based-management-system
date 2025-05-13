package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/handler"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/infra/repository"
	"github.com/KokoiRuby/rbac-based-management-system/backend/service"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"time"
)

func NewUploadAvatarRouter(cfg *config.RuntimeConfig, objStore *s3.Client, group *gin.RouterGroup) {
	repo := repository.NewAvatarRepository(objStore)
	h := handler.UploadAvatarHandler{
		AvatarService: service.NewAvatarService(repo, time.Duration(cfg.Gin.Timeout)),
		RuntimeConfig: cfg,
	}
	group.POST("/upload/avatar",
		middleware.AuthNMiddleware,
		h.Upload)
}

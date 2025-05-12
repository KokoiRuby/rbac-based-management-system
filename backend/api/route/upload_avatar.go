package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/handler"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/service"
	"github.com/gin-gonic/gin"
	"time"
)

func NewUploadAvatarRouter(cfg *config.RuntimeConfig, group *gin.RouterGroup) {
	h := handler.UploadAvatarHandler{
		UploadAvatarService: service.NewUploadAvatarService(time.Duration(cfg.Gin.Timeout)),
		RuntimeConfig:       cfg,
	}
	group.POST("/upload/avatar",
		middleware.AuthNMiddleware,
		h.Upload)
}

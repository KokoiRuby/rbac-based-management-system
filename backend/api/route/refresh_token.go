package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/handler"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/infra/repository"
	"github.com/KokoiRuby/rbac-based-management-system/backend/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func NewRefreshTokenRouter(cfg *config.RuntimeConfig, db *gorm.DB, group *gin.RouterGroup) {
	repo := repository.NewUserRepository(db)
	h := handler.RefreshTokenHandler{
		RefreshTokenService: service.NewRefreshTokenService(repo, time.Duration(cfg.Gin.Timeout)),
		RuntimeConfig:       cfg,
	}
	group.POST("/refresh", middleware.BindFormMiddleware[model.RefreshTokenRequest], h.Refresh)
}

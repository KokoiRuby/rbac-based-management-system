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

func NewListMenusRouter(cfg *config.RuntimeConfig, db *gorm.DB, group *gin.RouterGroup) {
	repo := repository.NewMenuRepository(db)
	h := handler.ListMenusHandler{
		ListMenusService: service.NewListMenusService(repo, time.Duration(cfg.Gin.Timeout)),
		RuntimeConfig:    cfg,
	}
	group.GET("/list",
		middleware.AuthNMiddleware,
		middleware.BindQueryMiddleware[model.MenuListRequest],
		h.ListMenus)
}

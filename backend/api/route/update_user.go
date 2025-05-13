package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/handler"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/infra/repository"
	"github.com/KokoiRuby/rbac-based-management-system/backend/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

func NewUpdateUserRouter(cfg *config.RuntimeConfig, db *gorm.DB, client *redis.Client, group *gin.RouterGroup) {
	repo := repository.NewUserRepository(db)
	cache := repository.NewRedisCache(client)
	h := handler.UpdateUserHandler{
		UpdateUserService: service.NewUpdateUserService(repo, cache, time.Duration(cfg.Gin.Timeout)),
		RuntimeConfig:     cfg,
	}
	group.POST("/update",
		middleware.AuthNMiddleware,
		middleware.BindFormMiddleware[model.UserUpdate],
		h.Update)
	group.POST("/update/confirm",
		h.UpdateConfirm)

}

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
	"time"
)

func NewSignoutRouter(cfg *config.RuntimeConfig, client *redis.Client, group *gin.RouterGroup) {
	repo := repository.NewRedisRepository(client)
	h := handler.SignoutHandler{
		SignoutService: service.NewSignoutService(repo, time.Duration(cfg.Gin.Timeout)),
		RuntimeConfig:  cfg,
	}
	group.POST("/signout", middleware.BindFormMiddleware[model.SignoutRequest], h.Signout)
}

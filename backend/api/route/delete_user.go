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

func NewDeleteUserRouter(cfg *config.RuntimeConfig, db *gorm.DB, group *gin.RouterGroup) {
	repo := repository.NewUserRepository(db)
	h := handler.DeleteUserHandler{
		DeleteUserService: service.NewDeleteUserService(repo, time.Duration(cfg.Gin.Timeout)),
		RuntimeConfig:     cfg,
	}
	// TODO: Delete role list before deleting the user.
	group.DELETE("/:userID/delete",
		middleware.AuthNMiddleware,
		middleware.BindUriMiddleware[model.DeleteUserRequest],
		h.DeleteAUser)
	group.DELETE("/deletes",
		middleware.AuthNMiddleware,
		middleware.BindJsonMiddleware[model.DeleteUsersRequest],
		h.DeleteUsers)

}

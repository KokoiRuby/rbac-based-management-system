package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewResetPasswordRouter(cfg *config.RuntimeConfig, db *gorm.DB, group *gin.RouterGroup) {
	//repo := repository.NewUserRepository(db)
	//h := handler.ResetPasswordHandler{
	//	ResetPasswordService: service.NewResetPasswordService(repo, time.Duration(cfg.Gin.Timeout)),
	//	RuntimeConfig:        cfg,
	//}
	//group.POST("/:id/resetPassword", middleware.AuthNMiddleware, h.ResetPassword)
}

//| 路由                                | 说明                |
//| --------------------------------- | ----------------- |
//| `POST /v1/user/:id/resetPassword` | 管理员为指定用户重置密码      |
//| `POST /v1/user/resetPassword`     | 当前用户自己重置密码（无需 ID） |

package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Setup(cfg *config.RuntimeConfig, db *gorm.DB, cache *redis.Client, objStore *s3.Client, g *gin.Engine) {
	// Statics
	g.Static("/uploads", "./static/uploads")

	// Public APIs
	publicRouter := g.Group("")
	//NewCreateAdminRouter(cfg, db, publicRouter)
	NewSignupRouter(cfg, db, cache, publicRouter)
	NewSigninRouter(cfg, db, cache, publicRouter)
	NewRefreshTokenRouter(cfg, db, publicRouter)
	NewForgotPasswordRouter(cfg, db, cache, publicRouter)

	// Protected APIs
	probeRouter := g.Group("")
	NewReadinessRouter(probeRouter) // TODO: admin
	NewLivenessRouter(probeRouter)  // TODO: admin

	protectedRouter := g.Group("v1")
	userGroup := protectedRouter.Group("user")
	NewSignoutRouter(cfg, cache, userGroup)
	NewResetPasswordRouter(cfg, db, userGroup)
	NewUploadAvatarRouter(cfg, db, objStore, userGroup)
	NewUpdateUserRouter(cfg, db, cache, userGroup)
	NewUserProfileRouter(cfg, db, userGroup)
	NewListUsersRouter(cfg, db, userGroup) // TODO: admin
	NewCloseUserRouter(cfg, db, userGroup)
	NewDeleteUserRouter(cfg, db, userGroup) // TODO: admin

	menuGroup := protectedRouter.Group("menu")
	NewCreateMenuRouter(cfg, db, menuGroup)
	NewListMenusRouter(cfg, db, menuGroup)
}

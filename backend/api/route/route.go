package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/bootstrap"
	"github.com/gin-gonic/gin"
)

func Setup(app *bootstrap.App, g *gin.Engine) {
	// Statics
	g.Static("/uploads", "./static/uploads")

	// Probes
	probeRouter := g.Group("")
	NewReadinessRouter(probeRouter)
	NewLivenessRouter(probeRouter)

	// Public APIs
	publicRouter := g.Group("v1")
	_ = publicRouter

	// Protected APIs
	protectedRouter := g.Group("v1")
	_ = protectedRouter
}

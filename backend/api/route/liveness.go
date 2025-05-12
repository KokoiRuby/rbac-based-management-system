package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/handler"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewLivenessRouter(group *gin.RouterGroup) {
	h := handler.LivenessHandler{}
	//group.GET("/liveness", h.Probe)
	group.GET("/liveness", middleware.LimitMiddleware(3), h.Probe)
}

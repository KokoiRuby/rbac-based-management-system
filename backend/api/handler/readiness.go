package handler

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/lifecycle"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReadinessHandler struct {
}

func (ReadinessHandler) Probe(c *gin.Context) {
	if lifecycle.IsReady() {
		utils.OKWithMsg(c, http.StatusOK, "ready")
	} else {
		utils.FailWithMsg(c, http.StatusServiceUnavailable, "not ready")
	}
}

package handler

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LivenessHandler struct {
}

func (LivenessHandler) Probe(c *gin.Context) {
	utils.OKWithMsg(c, http.StatusOK, "up")
}

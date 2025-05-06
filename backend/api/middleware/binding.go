package middleware

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindJsonMiddleware[T any](c *gin.Context) {
	var req T
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithBindingErr(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}
	c.Set("request", req)
	return
}

func BindQueryMiddleware[T any](c *gin.Context) {
	var req T
	err := c.ShouldBindQuery(&req)
	if err != nil {
		utils.FailWithBindingErr(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}
	c.Set("request", req)
	return
}

func BindUriMiddleware[T any](c *gin.Context) {
	var req T
	err := c.ShouldBindUri(&req)
	if err != nil {
		utils.FailWithBindingErr(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}
	c.Set("request", req)
	return
}

func BindFormMiddleware[T any](c *gin.Context) {
	var req T
	err := c.ShouldBind(&req)
	if err != nil {
		utils.FailWithBindingErr(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}
	c.Set("request", req)
	return
}

func GetBind[T any](c *gin.Context) (cr T) {
	return c.MustGet("request").(T)
}

package utils

import (
	"github.com/gin-gonic/gin"
)

type status int
type code int

var statusToCode = map[status]code{}

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func OK(c *gin.Context, status int, data any, msg string) {
	c.JSON(status, Response{
		Code:    0,
		Data:    data,
		Message: msg,
	})
}

func OKWithData(c *gin.Context, status int, data any) {
	OK(c, status, data, "Success")
}

func OKWithMsg(c *gin.Context, status int, msg string) {
	OK(c, status, gin.H{}, msg)
}

func OKWithList(c *gin.Context, status int, count int64, list any) {
	OK(c, status, map[string]any{
		"count": count,
		"list":  list,
	}, "Success")
}

func Fail(c *gin.Context, status int, code int, msg string) {
	// TODO: map[status] = code
	c.JSON(status, Response{
		Code:    code,
		Data:    gin.H{},
		Message: msg,
	})
}

func FailWithMsg(c *gin.Context, status int, msg string) {
	// TODO: map[status] = code
	Fail(c, status, 1, msg)
}

func FailWithErr(c *gin.Context, status int, err error) {
	// TODO: map[status] = code
	Fail(c, status, 1, err.Error())
}

func FailWithBindingErr(c *gin.Context, status int, err error) {
	msg := ValidateError(err)
	// TODO: map[status] = code
	Fail(c, status, 1, msg)
}

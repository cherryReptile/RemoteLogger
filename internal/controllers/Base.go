package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (c *BaseController) ERROR(ctx *gin.Context, code int, err error) {
	ctx.Header("Content-Type", "application/json")
	ctx.Status(code)

	ctx.JSON(code, gin.H{
		"status": http.StatusText(code),
		"error":  err.Error(),
	})
}

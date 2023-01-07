package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Test(ctx *gin.Context) {
	ctx.MustGet("token")
	t, _ := ctx.Get("token")
	ctx.JSON(http.StatusOK, gin.H{"message": "test", "token": t})
}

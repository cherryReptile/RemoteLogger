package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
	BaseController
}

func (c *TestController) Init() {}

func (c *TestController) Test(ctx *gin.Context) {
	ctx.MustGet("token")
	ctx.JSON(http.StatusOK, "test")
	t, _ := ctx.Get("token")
	ctx.JSON(http.StatusOK, t)
}

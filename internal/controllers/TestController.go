package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
	BaseController
	DatabaseController
}

func (c *TestController) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{"status": http.StatusText(http.StatusOK)})
}

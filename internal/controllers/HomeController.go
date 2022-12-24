package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Init() {
	//
}

func (c *HomeController) AddApp(ctx *gin.Context) {

}

func (c *HomeController) AddGithub(ctx *gin.Context) {

}

func (c *HomeController) AddGoogle(ctx *gin.Context) {

}

func (c *HomeController) AddTelegram(ctx *gin.Context) {

}

func (c *HomeController) Test(ctx *gin.Context) {
	ctx.MustGet("token")
	t, _ := ctx.Get("token")
	ctx.JSON(http.StatusOK, gin.H{"message": "test", "token": t})
}

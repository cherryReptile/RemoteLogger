package router

import (
	"github.com/cherryReptile/WS-AUTH/rest/controllers"
	"github.com/gin-gonic/gin"
)

func NewGoogleRouter(group *gin.RouterGroup, gc *controllers.GoogleController) {
	gg := group.Group("/google")
	gg.GET("/", gc.RedirectToGoogle)
	gg.GET("/token", gc.GetAccessToken)
	gg.POST("/login", gc.Login)
}

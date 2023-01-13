package router

import (
	controllers2 "github.com/cherryReptile/WS-AUTH/rest/controllers"
	"github.com/gin-gonic/gin"
)

func NewGitHubRouter(group *gin.RouterGroup, gc *controllers2.GithubController) {
	gg := group.Group("/github")
	gg.GET("/", gc.RedirectToGoogle)
	gg.GET("/token", gc.GetAccessToken)
	gg.POST("/login", gc.Login)
}

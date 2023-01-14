package router

import (
	"github.com/cherryReptile/WS-AUTH/rest/controllers"
	"github.com/gin-gonic/gin"
)

func NewAppRouter(group *gin.RouterGroup, ac *controllers.AppController) {
	ag := group.Group("/app")
	ag.POST("/register", ac.Register)
	ag.POST("/login", ac.Login)
}

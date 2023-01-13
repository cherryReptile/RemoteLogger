package router

import (
	"github.com/cherryReptile/WS-AUTH/rest/controllers"
	"github.com/gin-gonic/gin"
)

func NewProfileRouter(group *gin.RouterGroup, c *controllers.ProfileController) {
	p := group.Group("/profile")
	p.POST("/create", c.Create)
	p.GET("/get", c.Get)
	p.PATCH("/update", c.Update)
	p.DELETE("/delete", c.Delete)
}

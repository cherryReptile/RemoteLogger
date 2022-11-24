package middlewars

import (
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"net/http"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := helpers.CheckAuthHeader(c)
		if !ok || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "auth header is missing or value is required")
			return
		}
		c.Set("token", token)
		//c.Next()
	}
}

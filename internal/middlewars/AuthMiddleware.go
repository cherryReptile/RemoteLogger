package middlewars

import (
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/client"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"net/http"
)

func CheckAuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := helpers.CheckAuthHeader(c)
		if !ok || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "auth header is missing or value is required"})
			return
		}
		c.Set("token", token)
		c.Next()
	}
}

func CheckUserAndToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := helpers.GetAndCastToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		res, err := client.CheckAuth(&api.TokenRequest{Token: token})
		if err != nil || !res.Ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		//c.Next()
	}
}

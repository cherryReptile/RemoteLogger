package middlewars

import (
	"github.com/gin-gonic/gin"
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

//func CheckFromGithub() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t, ok := c.Get("token")
//		if !ok {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
//			return
//		}
//
//		res, err := helpers.RequestToGithub(t.(string))
//		defer res.Body.Close()
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusBadRequest, err)
//			return
//		}
//
//		if res.StatusCode != 200 {
//			c.AbortWithStatusJSON(http.StatusBadRequest, "authorization failed")
//			return
//		}
//
//		c.Next()
//	}
//}

func CheckUserAndToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, ok := c.Get("token")
		if !ok || t == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "auth token is missing"})
			return
		}

		service, err := c.Cookie("service")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		switch service {
		case "github":
			CheckGithub(c, t.(string))
		case "app":
			CheckApp(c, t.(string))
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "unknown service"})
			return
		}

		//c.Next()
	}
}

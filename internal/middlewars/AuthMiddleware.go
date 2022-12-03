package middlewars

import (
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"github.com/pavel-one/GoStarter/internal/models"
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
			CheckGoogleOrGitHub(c, new(models.GithubUser), t.(string))
		case "google":
			CheckGoogleOrGitHub(c, new(models.GoogleUser), t.(string))
		case "app":
			CheckApp(c, new(models.AppUser), t.(string))
		case "telegram":
			CheckApp(c, new(models.TelegramUser), t.(string))
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "unknown service"})
			return
		}

		//c.Next()
	}
}

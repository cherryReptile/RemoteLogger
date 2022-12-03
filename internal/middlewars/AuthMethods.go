package middlewars

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pavel-one/GoStarter/internal/appauth"
	"github.com/pavel-one/GoStarter/internal/models"
	"net/http"
)

func CheckApp(c *gin.Context, t string) {
	user := new(models.AppUser)
	claims, err := appauth.GetClaims(t)
	if err != nil {
		if err.(*jwt.ValidationError).Errors == 16 {
			db, ok := user.CheckAndUpdateDb(claims.Unique)
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			token, _ := user.GetTokenByStr(db, t)
			if token.ID == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			token.Delete(db)
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	db, ok := user.CheckAndUpdateDb(claims.Unique)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	tokenModel, err := user.GetTokenByStr(db, t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if tokenModel.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
		return
	}
	c.Set("user", user.Email)
}

func CheckGoogleOrGitHub(c *gin.Context, t, service string) {
	var user models.OAuthModel
	switch service {
	case "github":
		user = new(models.GithubUser)
	case "google":
		user = new(models.GoogleUser)
	}

	unique, err := c.Cookie("user")
	if err != nil || unique == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unknown user"})
		return
	}

	db, ok := user.CheckAndUpdateDb(unique)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	token, err := user.GetTokenByStr(db, t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if token.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "token not found"})
		return
	}
}

func CheckTelegram(c *gin.Context) {
	//
}

package middlewars

import (
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/internal/appauth"
	"github.com/pavel-one/GoStarter/internal/models"
	"net/http"
)

func CheckApp(c *gin.Context, t string) {
	user := new(models.AppUser)
	claims, err := appauth.GetClaims(t)

	db, ok := user.CheckDb(claims.Login)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, "user not found")
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

func CheckGithub(c *gin.Context, t string) {
	user := new(models.GithubUser)
	login, err := c.Cookie("user")
	if err != nil || login == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "unknown user")
		return
	}

	db, ok := user.CheckDb(login)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, "user not found")
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

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

	tokenModel, err := user.GetAccessToken(db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if tokenModel.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
		return
	}
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

	token, err := user.GetAccessToken(db)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if token.Token != t {
		c.AbortWithStatusJSON(http.StatusBadRequest, "please use your last token")
		return
	}
}

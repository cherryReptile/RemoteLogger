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
			c.AbortWithStatusJSON(http.StatusUnauthorized, "auth header is missing or value is required")
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
		user := new(models.GithubUser)
		t, ok := c.Get("token")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "auth token is missing")
			return
		}

		service, err := c.Cookie("service")
		if err != nil || service == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "unknown service")
			return
		}

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

		//c.Next()
	}
}

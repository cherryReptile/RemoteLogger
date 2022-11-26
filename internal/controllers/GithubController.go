package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"github.com/pavel-one/GoStarter/internal/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"os"
)

type GithubAuthController struct {
	BaseController
	Config *oauth2.Config
}

var GitRedirectLogin = "/api/v1/auth/github/login"

func (c *GithubAuthController) Init() {
	c.Config = &oauth2.Config{}
	c.Config.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	c.Config.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	c.Config.Endpoint = github.Endpoint
}

func (c *GithubAuthController) RedirectForAuth(ctx *gin.Context) {
	c.Config.RedirectURL = "http://" + os.Getenv("DOMAIN") + GitRedirectLogin
	u := c.Config.AuthCodeURL(c.setCookie(ctx))
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

func (c *GithubAuthController) Login(ctx *gin.Context) {
	token := new(models.AccessToken)
	oauthState, err := ctx.Cookie("oauthstate")

	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	if ctx.Query("state") != oauthState {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("invalid state"))
		return
	}

	code := ctx.Query("code")
	ctxC := context.Background()

	tok, err := c.Config.Exchange(ctxC, code)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := c.getGitHubUser(tok.AccessToken)

	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	db, ok := user.CheckDb(user.Login)
	if db == nil || !ok {
		db, err = user.Create(user.Login)
		if err != nil {
			c.ERROR(ctx, http.StatusBadRequest, err)
			return
		}
	}

	if user.ID == 0 {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("user not found"))
		return
	}

	token.Token = tok.AccessToken
	token.UserID = user.ID
	if err = token.Create(db); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	c.setUIDCookie(ctx, user.Login)

	ctx.JSON(http.StatusOK, token.Token)
	ctx.JSON(http.StatusOK, user)
}

func (c *GithubAuthController) setCookie(ctx *gin.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("oauthstate", state, 3600, GitRedirectLogin, os.Getenv("DOMAIN"), false, true)

	return state
}

func (c *GithubAuthController) setUIDCookie(ctx *gin.Context, login string) {
	path := "/api/v1/home"
	ctx.SetCookie("service", "github", 3600, path, os.Getenv("DOMAIN"), false, true)
	ctx.SetCookie("user", login, 3600, path, os.Getenv("DOMAIN"), false, true)
}

func (c *GithubAuthController) getGitHubUser(token string) (*models.GithubUser, error) {
	user := new(models.GithubUser)
	//client := &http.Client{}
	//req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	//
	//if err != nil {
	//	return nil, err
	//}
	//req.Header.Set("Authorization", "Bearer "+token)
	//
	//res, err := client.Do(req)
	//
	//if err != nil {
	//	return nil, err
	//}

	res, err := helpers.RequestToGithub(token)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

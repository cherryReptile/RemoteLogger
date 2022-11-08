package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"os"
)

type GithubAuthController struct {
	BaseController
	DatabaseController
	Config *oauth2.Config
}

var GitRedirectLogin = "/api/v1/auth/github/login"

func (c *GithubAuthController) Init(db *sqlx.DB) {
	c.DB = db
	c.Config = &oauth2.Config{}
	c.Config.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	c.Config.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	c.Config.Endpoint = github.Endpoint

}

func (c *GithubAuthController) RedirectForAuth(ctx *gin.Context) {
	c.Config.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	c.Config.RedirectURL = "http://" + os.Getenv("DOMAIN") + GitRedirectLogin
	u := c.Config.AuthCodeURL(c.setCookie(ctx))
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

func (c *GithubAuthController) Login(ctx *gin.Context) {
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

	reqUser, err := c.getGitHubUser(tok.AccessToken)

	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, reqUser)
}

func (c *GithubAuthController) setCookie(ctx *gin.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("oauthstate", state, 3600, GitRedirectLogin, os.Getenv("DOMAIN"), false, true)

	return state
}

func (c *GithubAuthController) getGitHubUser(token string) (any, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var reqUser struct {
		Id    int    `json:"id"`
		Login string `json:"login"`
	}

	err = json.NewDecoder(res.Body).Decode(&reqUser)

	if err != nil {
		return nil, err
	}

	return reqUser, nil
}

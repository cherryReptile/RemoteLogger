package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/client"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"github.com/pavel-one/GoStarter/internal/resources/requests"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"os"
	"strings"
)

type GithubAuthController struct {
	BaseJwtAuthController
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
	u := c.Config.AuthCodeURL(c.setOAuthStateCookie(ctx, GitRedirectLogin, os.Getenv("DOMAIN")))
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

	res, err := client.GithubLogin(&api.GitHubRequest{Login: reqUser.Login})
	if err != nil {
		e := strings.Split(err.Error(), "=")
		c.ERROR(ctx, http.StatusBadRequest, errors.New(e[2]))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": res.Struct, "token": res.TokenStr})
}

func (c *GithubAuthController) Logout(ctx *gin.Context) {
	token, err := helpers.GetAndCastToken(ctx)
	if err != nil {
		c.ERROR(ctx, http.StatusUnauthorized, err)
		return
	}

	res, err := client.Logout(&api.TokenRequest{Token: token})
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": res.Message})
}

func (c *GithubAuthController) getGitHubUser(token string) (*requests.GithubUser, error) {
	user := new(requests.GithubUser)
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

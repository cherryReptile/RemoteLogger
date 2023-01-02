package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/internal/helpers"
	"github.com/cherryReptile/WS-AUTH/internal/resources/requests"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"net/http"
	"os"
)

type GithubAuthController struct {
	BaseOAuthController
	GithubService api.AuthGithubServiceClient
	Config        *oauth2.Config
}

var GitRedirectToExchangeToken = "/api/v1/auth/github/token"

func (c *GithubAuthController) Init(gs api.AuthGithubServiceClient) {
	c.GithubService = gs
	c.Config = &oauth2.Config{}
	c.Config.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	c.Config.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	c.Config.Endpoint = github.Endpoint
}

func (c *GithubAuthController) RedirectToGoogle(ctx *gin.Context) {
	c.Config.RedirectURL = "http://" + os.Getenv("DOMAIN") + GitRedirectToExchangeToken
	u := c.Config.AuthCodeURL(c.setOAuthStateCookie(ctx, GitRedirectToExchangeToken, os.Getenv("DOMAIN")))
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

func (c *GithubAuthController) GetAccessToken(ctx *gin.Context) {
	code, err := c.checkOAuthStateCookie(ctx)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	tok, err := c.Config.Exchange(context.Background(), code)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"github_access_token": tok.AccessToken})
}

func (c *GithubAuthController) Login(ctx *gin.Context) {
	t := new(requests.Token)
	if err := ctx.ShouldBindJSON(t); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	login, body, err := c.getGitHubUserAndBody(t.Token)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := c.GithubService.Login(context.Background(), &api.OAuthRequest{Username: login, Data: body})
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": res.Struct, "token": res.TokenStr})
}

func (c *GithubAuthController) AddAccount(ctx *gin.Context) {
	t := new(requests.Token)
	if err := ctx.ShouldBindJSON(t); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	uuid, err := helpers.GetAndCastUserUUID(ctx)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	login, body, err := c.getGitHubUserAndBody(t.Token)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := c.GithubService.AddAccount(context.Background(), &api.AddOauthRequest{
		UserUUID: uuid,
		Request:  &api.OAuthRequest{Username: login, Data: body},
	})
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Struct})
}

func (c *GithubAuthController) getGitHubUserAndBody(token string) (string, []byte, error) {
	user := new(requests.GithubUser)
	res, err := helpers.RequestToGithub(token)
	if err != nil {
		return "", nil, err
	}

	if res.StatusCode != http.StatusOK {
		return "", nil, errors.New("github oauth2 failed because returning not ok code")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}

	err = json.Unmarshal(body, user)

	if err != nil {
		return "", nil, err
	}

	return user.Login, body, nil
}

package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/internal/helpers"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"net/http"
	"os"
)

type GithubController struct {
	BaseOAuthController
	GithubService api.AuthGithubServiceClient
	Config        *oauth2.Config
}

type GithubUser struct {
	ID        uint   `json:"id" binding:"required"`
	Login     string `json:"login" binding:"required"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url" binding:"required"`
}

var GitRedirectToExchangeToken = "/api/v1/auth/github/token"

func (c *GithubController) Init(gs api.AuthGithubServiceClient) {
	c.GithubService = gs
	c.Config = &oauth2.Config{}
	c.Config.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	c.Config.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	c.Config.Endpoint = github.Endpoint
}

func (c *GithubController) RedirectToGoogle(ctx *gin.Context) {
	c.Config.RedirectURL = "http://" + os.Getenv("DOMAIN") + GitRedirectToExchangeToken
	u := c.Config.AuthCodeURL(c.setOAuthStateCookie(ctx, GitRedirectToExchangeToken, os.Getenv("DOMAIN")))
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

func (c *GithubController) GetAccessToken(ctx *gin.Context) {
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

func (c *GithubController) Login(ctx *gin.Context) {
	t := new(Token)
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

func (c *GithubController) AddAccount(ctx *gin.Context) {
	t := new(Token)
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

func (c *GithubController) getGitHubUserAndBody(token string) (string, []byte, error) {
	user := new(GithubUser)
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

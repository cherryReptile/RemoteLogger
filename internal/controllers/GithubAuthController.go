package controllers

import (
	"context"
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
	BaseOAuthController
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

	db, ok := user.CheckAndUpdateDb(user.Login)
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

	c.setServiceCookie(ctx, "github", os.Getenv("DOMAIN"))
	c.setUIDCookie(ctx, user.Login, os.Getenv("DOMAIN"))

	ctx.JSON(http.StatusOK, gin.H{"user": user, "token": token.Token})
}

func (c *GithubAuthController) Logout(ctx *gin.Context) {
	if err := c.LogoutFromApp(ctx, new(models.GithubUser)); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
}

func (c *GithubAuthController) getGitHubUser(token string) (*models.GithubUser, error) {
	user := new(models.GithubUser)
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

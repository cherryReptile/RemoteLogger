package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"net/http"
	"net/url"
	"os"
)

type GithubAuthController struct {
	BaseController
	DatabaseController
	Config *oauth2.Config
}

func (c *GithubAuthController) Init(db *sqlx.DB) {
	c.DB = db
	c.Config = &oauth2.Config{}
	c.Config.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	c.Config.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	c.Config.Endpoint = github.Endpoint

}

func (c *GithubAuthController) RedirectForAuth(ctx *gin.Context) {
	u := url.Values{}
	u.Set("client_id", os.Getenv("GITHUB_CLIENT_ID"))
	u.Set("redirect_uri", "http://"+os.Getenv("DOMAIN")+"/api/v1/auth/github/login")
	ctx.Redirect(http.StatusTemporaryRedirect, c.Config.Endpoint.AuthURL+"?"+u.Encode())
}

func (c *GithubAuthController) Login(ctx *gin.Context) {
	code := ctx.Query("code")
	ctxC := context.Background()

	tok, err := c.Config.Exchange(ctxC, code)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, string(body))
}

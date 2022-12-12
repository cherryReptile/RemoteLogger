package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"github.com/pavel-one/GoStarter/internal/resources/requests"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
	"strings"
)

type GoogleAuthController struct {
	BaseAuthController
	GoogleService api.AuthGoogleServiceClient
	Config        *oauth2.Config
}

func (c *GoogleAuthController) Init(gs api.AuthGoogleServiceClient) {
	c.GoogleService = gs
	c.Config = &oauth2.Config{}
	c.Config.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	c.Config.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	c.Config.Scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}
	c.Config.Endpoint = google.Endpoint
}

var GoogleRedirectLogin = "/api/v1/auth/google/login"

func (c *GoogleAuthController) RedirectForAuth(ctx *gin.Context) {
	c.Config.RedirectURL = "http://" + "localhost" + GoogleRedirectLogin
	u := c.Config.AuthCodeURL(c.setOAuthStateCookie(ctx, GoogleRedirectLogin, "localhost"))
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

func (c *GoogleAuthController) Login(ctx *gin.Context) {
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

	reqUser, err := c.getGoogleUser(tok.AccessToken)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := c.GoogleService.Login(context.Background(), &api.GoogleRequest{Email: reqUser.Email})
	if err != nil {
		e := strings.Split(err.Error(), "=")
		c.ERROR(ctx, http.StatusBadRequest, errors.New(e[2]))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": res.Struct, "token": res.TokenStr})
}

func (c *GoogleAuthController) getGoogleUser(token string) (*requests.GoogleUser, error) {
	user := new(requests.GoogleUser)
	res, err := helpers.RequestToGoogle(token)
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

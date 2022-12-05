package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/internal/appauth"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"github.com/pavel-one/GoStarter/internal/models"
	"github.com/pavel-one/GoStarter/internal/resources/requests"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
)

var GoogleRedirectLogin = "/api/v1/auth/google/login"

type GoogleAuthController struct {
	BaseJwtAuthController
	Config *oauth2.Config
}

func (c *GoogleAuthController) Init(db *sqlx.DB) {
	c.DB = db
	c.Config = &oauth2.Config{}
	c.Config.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	c.Config.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	c.Config.Scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}
	c.Config.Endpoint = google.Endpoint
}

func (c *GoogleAuthController) RedirectForAuth(ctx *gin.Context) {
	c.Config.RedirectURL = "http://" + "localhost" + GoogleRedirectLogin
	u := c.Config.AuthCodeURL(c.setOAuthStateCookie(ctx, GoogleRedirectLogin, "localhost"))
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

func (c *GoogleAuthController) Login(ctx *gin.Context) {
	user := new(models.User)
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

	reqUser, err := c.getGoogleUser(tok.AccessToken)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	user.FindByUniqueAndService(c.DB, reqUser.Email, "google")
	if user.ID == 0 {
		user.UniqueRaw = reqUser.Email
		user.AuthorizedBy = "google"
		if err = user.Create(c.DB); err != nil {
			c.ERROR(ctx, http.StatusBadRequest, err)
			return
		}
	}

	if user.ID == 0 {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("user not found"))
		return
	}

	tokenStr, err := appauth.GenerateToken(user.ID, user.UniqueRaw, user.AuthorizedBy)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	token.Token = tokenStr
	token.UserID = user.ID
	if err = token.Create(c.DB); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func (c *GoogleAuthController) Logout(ctx *gin.Context) {
	if err := c.LogoutFromApp(ctx, c.DB); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
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

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
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
)

var GoogleRedirectLogin = "/api/v1/auth/google/login"

type GoogleAuthController struct {
	BaseController
	Config *oauth2.Config
}

func (c *GoogleAuthController) Init() {
	c.Config = &oauth2.Config{}
	c.Config.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	c.Config.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	c.Config.Scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}
	c.Config.Endpoint = google.Endpoint
}

func (c *GoogleAuthController) RedirectForAuth(ctx *gin.Context) {
	c.Config.RedirectURL = "http://" + "localhost" + GoogleRedirectLogin
	u := c.Config.AuthCodeURL(c.setOAuthStateCookie(ctx))
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

func (c *GoogleAuthController) Login(ctx *gin.Context) {
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

	user, err := c.getGoogleUser(tok.AccessToken)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	db, ok := user.CheckDb(user.Email)
	if db == nil || !ok {
		db, err = user.Create(user.Email)
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

	c.setUIDCookie(ctx, user.Email)

	ctx.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func (c *GoogleAuthController) setOAuthStateCookie(ctx *gin.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("oauthstate", state, 3600, "/api/v1/auth/google", "localhost", false, true)

	return state
}

func (c *GoogleAuthController) setUIDCookie(ctx *gin.Context, login string) {
	path := "/api/v1/home"
	ctx.SetCookie("service", "github", 3600, path, os.Getenv("DOMAIN"), false, true)
	ctx.SetCookie("user", login, 3600, path, os.Getenv("DOMAIN"), false, true)
}

func (c *GoogleAuthController) getGoogleUser(token string) (*models.GoogleUser, error) {
	user := new(models.GoogleUser)
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

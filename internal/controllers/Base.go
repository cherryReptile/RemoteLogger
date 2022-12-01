package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/internal/models"
	"net/http"
)

type BaseController struct {
}

type BaseOAuthController struct {
	BaseController
}

func (c *BaseController) ERROR(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"status": http.StatusText(code),
		"error":  err.Error(),
	})
}

func (c *BaseOAuthController) LogoutFromApp(ctx *gin.Context, user models.OAuthModel) error {
	login, err := ctx.Cookie("user")
	if err != nil {
		return err
	}

	t, ok := ctx.Get("token")
	if !ok {
		return errors.New("cannot get token")
	}

	db, ok := user.CheckAndUpdateDb(login)
	if !ok {
		return errors.New("user not found")
	}

	token, err := user.GetTokenByStr(db, t.(string))
	if err != nil {
		return err
	}

	if token.ID == 0 {
		return errors.New("token not found")
	}

	if err = token.Delete(db); err != nil {
		return err
	}

	return nil
}

func (c *BaseOAuthController) setOAuthStateCookie(ctx *gin.Context, path, domain string) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("oauthstate", state, 3600, path, domain, false, true)

	return state
}

func (c *BaseOAuthController) setUIDCookie(ctx *gin.Context, service, unique, domain string) {
	path := "/api/v1/home"
	ctx.SetCookie("service", service, 3600, path, domain, false, true)
	ctx.SetCookie("user", unique, 3600, path, domain, false, true)
}

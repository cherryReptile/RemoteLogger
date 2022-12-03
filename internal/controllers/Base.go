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

type BaseJwtAuthController struct {
	BaseController
}

var homePath = "/api/v1/home"

func (c *BaseController) ERROR(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"status": http.StatusText(code),
		"error":  err.Error(),
	})
}

func (c *BaseController) setServiceCookie(ctx *gin.Context, service, domain string) {
	ctx.SetCookie("service", service, 3600, homePath, domain, false, true)
}

func (c *BaseController) setUIDCookie(ctx *gin.Context, unique, domain string) {
	ctx.SetCookie("user", unique, 3600, homePath, domain, false, true)
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

func (c *BaseJwtAuthController) LogoutFromApp(ctx *gin.Context, user models.OAuthModel) error {
	t, ok := ctx.Get("token")
	if !ok {
		return errors.New("cannot get token")
	}

	unique, ok := ctx.Get("user")
	if !ok {
		return errors.New("cannot get user")
	}

	db, ok := user.CheckAndUpdateDb(unique.(string))
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

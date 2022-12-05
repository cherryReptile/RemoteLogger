package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/internal/models"
	"net/http"
)

type BaseController struct {
}

type Database struct {
	DB *sqlx.DB
}

type BaseJwtAuthController struct {
	Database
	BaseController
}

var homePath = "/api/v1/home"

func (c *BaseController) ERROR(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"status": http.StatusText(code),
		"error":  err.Error(),
	})
}

func (c *BaseJwtAuthController) LogoutFromApp(ctx *gin.Context, db *sqlx.DB) error {
	user := new(models.User)
	t, ok := ctx.Get("token")
	if !ok {
		return errors.New("cannot get token")
	}

	unique, ok := ctx.Get("user")
	if !ok {
		return errors.New("cannot get user")
	}

	service, ok := ctx.Get("service")
	if !ok {
		return errors.New("unknown service")
	}

	if err := user.FindByUniqueAndService(db, unique.(string), service.(string)); err != nil {
		return err
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

func (c *BaseController) setServiceCookie(ctx *gin.Context, service, domain string) {
	ctx.SetCookie("service", service, 3600, homePath, domain, false, true)
}

func (c *BaseController) setUIDCookie(ctx *gin.Context, unique, domain string) {
	ctx.SetCookie("user", unique, 3600, homePath, domain, false, true)
}

func (c *BaseJwtAuthController) setOAuthStateCookie(ctx *gin.Context, path, domain string) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("oauthstate", state, 3600, path, domain, false, true)

	return state
}

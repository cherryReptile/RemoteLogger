package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type BaseController struct {
}

type BaseAuthController struct {
	BaseController
}

func (c *BaseController) ERROR(ctx *gin.Context, code int, err error) {
	e := strings.Split(err.Error(), "=")
	if len(e) != 1 {
		err = errors.New(e[2])
	}

	ctx.JSON(code, gin.H{
		"status": http.StatusText(code),
		"error":  err.Error(),
	})
}

func (c *BaseAuthController) setOAuthStateCookie(ctx *gin.Context, path, domain string) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("oauthstate", state, 3600, path, domain, false, true)

	return state
}

package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

type BaseJwtAuthController struct {
	BaseController
}

func (c *BaseController) ERROR(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"status": http.StatusText(code),
		"error":  err.Error(),
	})
}

func (c *BaseJwtAuthController) setOAuthStateCookie(ctx *gin.Context, path, domain string) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("oauthstate", state, 3600, path, domain, false, true)

	return state
}

package controllers

import (
	"context"
	"crypto/hmac"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

type TelegramController struct {
	BaseOAuthController
	TelegramService api.AuthTelegramServiceClient
}

type TelegramUser struct {
	ID        uint      `json:"id" binding:"required"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	PhotoURL  string    `json:"photo_url" binding:"required"`
	AuthDate  time.Time `json:"auth_date" binding:"required"`
	Hash      string    `json:"hash" binding:"required"`
}

func (c *TelegramController) Init(ts api.AuthTelegramServiceClient) {
	c.TelegramService = ts
}

func (c *TelegramController) Login(ctx *gin.Context) {
	reqUser := new(TelegramUser)
	if err := ctx.ShouldBindJSON(reqUser); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	checkAuthData := strings.Join([]string{reqUser.AuthDate.String(), reqUser.FirstName, reqUser.Hash, string(reqUser.ID), reqUser.LastName, reqUser.PhotoURL, reqUser.Username}, "\n")
	if !c.CheckHash(checkAuthData) {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("this not telegram request"))
		return
	}

	res, err := c.TelegramService.Login(context.Background(), &api.TelegramRequest{Username: reqUser.Username})
	if err != nil {
		e := strings.Split(err.Error(), "=")
		c.ERROR(ctx, http.StatusBadRequest, errors.New(e[2]))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": res.Struct, "token": res.TokenStr})
}

func (c *TelegramController) CheckHash(dataCheckString string) bool {
	if hmac.Equal([]byte(dataCheckString), []byte(os.Getenv("TG_BOT_TOKEN"))) {
		return true
	}

	return false
}

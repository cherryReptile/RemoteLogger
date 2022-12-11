package controllers

import (
	"crypto/hmac"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/client"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"github.com/pavel-one/GoStarter/internal/resources/requests"
	"net/http"
	"os"
	"strings"
)

type TelegramAuthController struct {
	BaseAuthController
}

func (c *TelegramAuthController) Init() {
	//
}

func (c *TelegramAuthController) Login(ctx *gin.Context) {
	reqUser := new(requests.TelegramUser)
	if err := ctx.ShouldBindJSON(reqUser); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	checkAuthData := strings.Join([]string{reqUser.AuthDate.String(), reqUser.FirstName, reqUser.Hash, string(reqUser.ID), reqUser.LastName, reqUser.PhotoURL, reqUser.Username}, "\n")
	if !c.CheckHash(checkAuthData) {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("this not telegram request"))
		return
	}

	res, err := client.TelegramLogin(&api.TelegramRequest{Username: reqUser.Username})
	if err != nil {
		e := strings.Split(err.Error(), "=")
		c.ERROR(ctx, http.StatusBadRequest, errors.New(e[2]))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": res.Struct, "token": res.TokenStr})
}

func (c *TelegramAuthController) Logout(ctx *gin.Context) {
	token, err := helpers.GetAndCastToken(ctx)
	if err != nil {
		c.ERROR(ctx, http.StatusUnauthorized, err)
		return
	}

	res, err := client.Logout(&api.TokenRequest{Token: token})
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": res.Message})
}

func (c *TelegramAuthController) CheckHash(dataCheckString string) bool {
	if hmac.Equal([]byte(dataCheckString), []byte(os.Getenv("TG_BOT_TOKEN"))) {
		return true
	}

	return false
}

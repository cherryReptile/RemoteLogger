package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"net/http"
)

type LogoutController struct {
	BaseAuthController
	LogoutService api.LogoutServiceClient
}

func (c *LogoutController) Init(ls api.LogoutServiceClient) {
	c.LogoutService = ls
}

func (c *LogoutController) Logout(ctx *gin.Context) {
	token, err := helpers.GetAndCastToken(ctx)
	if err != nil {
		c.ERROR(ctx, http.StatusUnauthorized, err)
		return
	}

	res, err := c.LogoutService.Logout(context.Background(), &api.TokenRequest{Token: token})
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": res.Message})
}

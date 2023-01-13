package controllers

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/internal/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutController struct {
	BaseOAuthController
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

package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/internal/resources/requests"
	"net/http"
)

type AppAuthController struct {
	BaseAuthController
	AppService api.AuthAppServiceClient
}

func (c *AppAuthController) Init(as api.AuthAppServiceClient) {
	c.AppService = as
}

func (c *AppAuthController) Register(ctx *gin.Context) {
	reqU := new(requests.UserRequest)
	if err := ctx.ShouldBindJSON(reqU); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := c.AppService.Register(context.Background(), &api.AppRequest{Email: reqU.Email, Password: reqU.Password})
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": res.Struct, "token": res.TokenStr})
}

func (c *AppAuthController) Login(ctx *gin.Context) {
	reqU := new(requests.UserRequest)
	if err := ctx.ShouldBindJSON(reqU); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := c.AppService.Login(context.Background(), &api.AppRequest{Email: reqU.Email, Password: reqU.Password})
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": res.Struct, "token": res.TokenStr})
}

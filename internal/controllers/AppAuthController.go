package controllers

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/internal/helpers"
	"github.com/cherryReptile/WS-AUTH/internal/resources/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppAuthController struct {
	BaseOAuthController
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

func (c *AppAuthController) AddAccount(ctx *gin.Context) {
	reqU := new(requests.UserRequest)
	if err := ctx.ShouldBindJSON(reqU); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	uuid, err := helpers.GetAndCastUserUUID(ctx)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := c.AppService.AddAccount(context.Background(), &api.AddAppRequest{
		UserUUID: uuid,
		Request:  &api.AppRequest{Email: reqU.Email, Password: reqU.Password},
	})
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Struct})
}

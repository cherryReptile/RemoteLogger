package controllers

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/internal/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppController struct {
	BaseController
	AppService api.AuthAppServiceClient
}

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (c *AppController) Init(as api.AuthAppServiceClient) {
	c.AppService = as
}

func (c *AppController) Register(ctx *gin.Context) {
	reqU := new(UserRequest)
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

func (c *AppController) Login(ctx *gin.Context) {
	reqU := new(UserRequest)
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

func (c *AppController) AddAccount(ctx *gin.Context) {
	reqU := new(UserRequest)
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

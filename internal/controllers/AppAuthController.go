package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/client"
	"github.com/pavel-one/GoStarter/internal/resources/requests"
	"net/http"
	"strings"
)

type AppAuthController struct {
	BaseJwtAuthController
}

func (c *AppAuthController) Init(db *sqlx.DB) {
	//c.DB = db
}

func (c *AppAuthController) Register(ctx *gin.Context) {
	//user := new(models.User)
	//tokenModel := new(models.AccessToken)
	reqU := new(requests.UserRequest)
	if err := ctx.ShouldBindJSON(reqU); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := client.Register(&api.AppRequest{Email: reqU.Email, Password: reqU.Password})
	if err != nil {
		e := strings.Split(err.Error(), "=")
		c.ERROR(ctx, http.StatusBadRequest, errors.New(e[2]))
		return
	}
	//user.FindByUniqueAndService(c.DB, reqU.Email, "app")
	//if user.ID != 0 {
	//	c.ERROR(ctx, http.StatusBadRequest, errors.New("this user already exists"))
	//	return
	//}
	//
	//hashP, err := bcrypt.GenerateFromPassword([]byte(reqU.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	c.ERROR(ctx, http.StatusBadRequest, err)
	//	return
	//}
	//
	//user.Password = string(hashP)
	//user.UniqueRaw = reqU.Email
	//user.AuthorizedBy = "app"
	//
	//err = user.Create(c.DB)
	//if err != nil {
	//	c.ERROR(ctx, http.StatusBadRequest, err)
	//	return
	//}
	//
	//if user.ID == 0 {
	//	c.ERROR(ctx, http.StatusBadRequest, errors.New("user not found"))
	//	return
	//}
	//
	//tokenStr, err := appauth.GenerateToken(user.ID, user.UniqueRaw, user.AuthorizedBy)
	//if err != nil {
	//	c.ERROR(ctx, http.StatusBadRequest, err)
	//	return
	//}
	//
	//tokenModel.Token = tokenStr
	//tokenModel.UserID = user.ID
	//
	//if err = tokenModel.Create(c.DB); err != nil {
	//	c.ERROR(ctx, http.StatusBadRequest, err)
	//	return
	//}

	ctx.JSON(http.StatusOK, gin.H{"user": res.Struct, "token": res.TokenStr})
	//ctx.JSON(http.StatusOK, gin.H{"user": user, "token": tokenModel})
}

func (c *AppAuthController) Login(ctx *gin.Context) {
	//user := new(models.User)
	//tokenModel := new(models.AccessToken)

	reqU := new(requests.UserRequest)
	if err := ctx.ShouldBindJSON(reqU); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := client.Login(&api.AppRequest{Email: reqU.Email, Password: reqU.Password})
	if err != nil {
		e := strings.Split(err.Error(), "=")
		c.ERROR(ctx, http.StatusBadRequest, errors.New(e[2]))
		return
	}

	//user.UniqueRaw = reqU.Email
	//if err := user.FindByUniqueAndService(c.DB, user.UniqueRaw, "app"); err != nil {
	//	c.ERROR(ctx, http.StatusBadRequest, err)
	//	return
	//}
	//
	//if user.ID == 0 {
	//	c.ERROR(ctx, http.StatusBadRequest, errors.New("user not found"))
	//	return
	//}
	//
	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqU.Password)); err != nil {
	//	c.ERROR(ctx, http.StatusBadRequest, err)
	//	return
	//}
	//
	//tokenStr, err := appauth.GenerateToken(user.ID, user.UniqueRaw, "app")
	//if err != nil {
	//	c.ERROR(ctx, http.StatusBadRequest, err)
	//	return
	//}
	//
	//tokenModel.Token = tokenStr
	//tokenModel.UserID = user.ID
	//if err = tokenModel.Create(c.DB); err != nil {
	//	c.ERROR(ctx, http.StatusBadRequest, err)
	//	return
	//}

	ctx.JSON(http.StatusOK, gin.H{"user": res.Struct, "token": res.TokenStr})
}

func (c *AppAuthController) Logout(ctx *gin.Context) {
	if err := c.LogoutFromApp(ctx, c.DB); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
}

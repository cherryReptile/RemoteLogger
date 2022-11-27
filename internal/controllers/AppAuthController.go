package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/internal/appauth"
	"github.com/pavel-one/GoStarter/internal/models"
	"github.com/pavel-one/GoStarter/internal/resources/requests"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

type AppAuthController struct {
	BaseController
}

func (c *AppAuthController) Init() {
}

func (c *AppAuthController) Register(ctx *gin.Context) {
	user := new(models.AppUser)
	tokenModel := new(models.AccessToken)
	reqU := new(requests.UserRequest)
	if err := ctx.ShouldBindJSON(reqU); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	user.Email = reqU.Email
	db, _ := user.CheckDb(user.Email)
	if db != nil {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("this user already exists"))
		return
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(reqU.Password), bcrypt.DefaultCost)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	user.Password = string(hashP)
	db, err = user.Create(user.Email)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	if user.ID == 0 {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("user not found"))
		return
	}

	tokenStr, err := appauth.GenerateToken(user)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	tokenModel.Token = tokenStr
	tokenModel.UserID = user.ID

	if err = tokenModel.Create(db); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	c.setServiceCookie(ctx)
	ctx.JSON(http.StatusOK, gin.H{"user": user, "token": tokenModel})
}

func (c *AppAuthController) Login(ctx *gin.Context) {
	user := new(models.AppUser)
	tokenModel := new(models.AccessToken)

	reqU := new(requests.UserRequest)
	if err := ctx.ShouldBindJSON(reqU); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	user.Email = reqU.Email
	db, ok := user.CheckDb(user.Email)
	if db == nil || !ok {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("user not found"))
		return
	}

	if user.ID == 0 {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("user not found"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqU.Password)); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	tokenStr, err := appauth.GenerateToken(user)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	tokenModel.Token = tokenStr
	tokenModel.UserID = user.ID
	if err = tokenModel.Create(db); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	c.setServiceCookie(ctx)
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusText(http.StatusOK), "token": tokenModel.Token})
}

func (c *AppAuthController) Logout(ctx *gin.Context) {
	user := new(models.AppUser)
	t, ok := ctx.Get("token")
	if !ok {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("cannot get token"))
		return
	}

	email, ok := ctx.Get("user")
	if !ok {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("cannot get user"))
		return
	}

	db, ok := user.CheckDb(email.(string))
	if !ok {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("user not found"))
		return
	}

	token, err := user.GetTokenByStr(db, t.(string))
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	if token.ID == 0 {
		c.ERROR(ctx, http.StatusBadRequest, errors.New("token not found"))
		return
	}

	if err = token.Delete(db); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "logout successfully"})
}

func (c *AppAuthController) setServiceCookie(ctx *gin.Context) {
	path := "/api/v1/home"
	ctx.SetCookie("service", "app", 3600, path, os.Getenv("DOMAIN"), false, true)
}

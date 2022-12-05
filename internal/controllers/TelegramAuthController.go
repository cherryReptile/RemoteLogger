package controllers

import (
	"crypto/hmac"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/internal/appauth"
	"github.com/pavel-one/GoStarter/internal/models"
	"github.com/pavel-one/GoStarter/internal/resources/requests"
	"net/http"
	"os"
	"strings"
)

type TelegramAuthController struct {
	BaseJwtAuthController
}

func (c *TelegramAuthController) Init(db *sqlx.DB) {
	c.DB = db
}

func (c *TelegramAuthController) Login(ctx *gin.Context) {
	tokenModel := new(models.AccessToken)
	user := new(models.User)
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

	user.FindByUniqueAndService(c.DB, reqUser.Username, "telegram")
	if user.ID == 0 {
		user.UniqueRaw = reqUser.Username
		user.AuthorizedBy = "telegram"
		if err := user.Create(c.DB); err != nil {
			c.ERROR(ctx, http.StatusBadRequest, err)
			return
		}
	}

	tokenStr, err := appauth.GenerateToken(user.ID, user.UniqueRaw, user.AuthorizedBy)
	if err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	tokenModel.Token = tokenStr
	tokenModel.UserID = user.ID
	if err = tokenModel.Create(c.DB); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user, "token": tokenModel.Token})
}

func (c *TelegramAuthController) Logout(ctx *gin.Context) {
	if err := c.LogoutFromApp(ctx, c.DB); err != nil {
		c.ERROR(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
}

func (c *TelegramAuthController) CheckHash(dataCheckString string) bool {
	if hmac.Equal([]byte(dataCheckString), []byte(os.Getenv("TG_BOT_TOKEN"))) {
		return true
	}

	return false
}

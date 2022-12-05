package controllers

//
//import (
//	"crypto/hmac"
//	"errors"
//	"github.com/gin-gonic/gin"
//	"github.com/jmoiron/sqlx"
//	"github.com/pavel-one/GoStarter/internal/appauth"
//	"github.com/pavel-one/GoStarter/internal/models"
//	"github.com/pavel-one/GoStarter/internal/resources/requests"
//	"net/http"
//	"os"
//	"strings"
//)
//
//type TelegramAuthController struct {
//	Database
//	BaseJwtAuthController
//}
//
//func (c *TelegramAuthController) Init(db *sqlx.DB) {
//	c.DB = db
//}
//
//func (c *TelegramAuthController) Login(ctx *gin.Context) {
//	tokenModel := new(models.AccessToken)
//	user := new(models.TelegramUser)
//	reqUser := new(requests.TelegramUser)
//	if err := ctx.ShouldBindJSON(user); err != nil {
//		c.ERROR(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	checkAuthData := strings.Join([]string{reqUser.AuthDate.String(), reqUser.FirstName, reqUser.Hash, string(reqUser.ID), reqUser.LastName, reqUser.PhotoURL, reqUser.Username}, "\n")
//	if !c.CheckHash(checkAuthData) {
//		c.ERROR(ctx, http.StatusBadRequest, errors.New("this not telegram request"))
//		return
//	}
//
//	user.TgID = reqUser.ID
//	user.FirstName = reqUser.FirstName
//	user.LastName = reqUser.LastName
//	user.Username = reqUser.Username
//	user.PhotoURL = reqUser.PhotoURL
//
//	db, ok := user.CheckAndUpdateDb(user.Username)
//	if db == nil || !ok {
//		_, err := user.Create(user.Username)
//		if err != nil {
//			c.ERROR(ctx, http.StatusBadRequest, err)
//			return
//		}
//	}
//
//	if user.ID == 0 {
//		c.ERROR(ctx, http.StatusBadRequest, errors.New("user not found"))
//		return
//	}
//
//	tokenStr, err := appauth.GenerateToken(user.ID, user.Username)
//	if err != nil {
//		c.ERROR(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	tokenModel.Token = tokenStr
//	tokenModel.UserID = user.ID
//	if err = tokenModel.Create(db); err != nil {
//		c.ERROR(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	c.setServiceCookie(ctx)
//
//	ctx.JSON(http.StatusOK, gin.H{"user": user, "token": tokenModel.Token})
//}
//
//func (c *TelegramAuthController) Logout(ctx *gin.Context) {
//	if err := c.LogoutFromApp(ctx, new(models.TelegramUser)); err != nil {
//		c.ERROR(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
//}
//
//func (c *TelegramAuthController) CheckHash(dataCheckString string) bool {
//	if hmac.Equal([]byte(dataCheckString), []byte(os.Getenv("TG_BOT_TOKEN"))) {
//		return true
//	}
//
//	return false
//}
//
//func (c *TelegramAuthController) setServiceCookie(ctx *gin.Context) {
//	path := "/api/v1/home"
//	ctx.SetCookie("service", "telegram", 3600, path, os.Getenv("DOMAIN"), false, true)
//}

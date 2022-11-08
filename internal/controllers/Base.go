package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type BaseController struct {
}

type DatabaseController struct {
	DB *sqlx.DB
}

func (c *BaseController) ERROR(ctx *gin.Context, code int, err error) {
	ctx.Header("Content-Type", "application/json")
	ctx.Status(code)

	ctx.JSON(code, map[string]interface{}{
		"status": http.StatusText(code),
		"error":  err.Error(),
	})
}

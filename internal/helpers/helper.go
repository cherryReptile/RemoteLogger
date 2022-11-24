package helpers

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func CheckAuthHeader(ctx *gin.Context) (string, bool) {
	authHeader := strings.Split(ctx.GetHeader("Authorization"), " ")

	if len(authHeader) < 2 || authHeader[1] == "" {
		return "", false
	}

	return authHeader[1], true
}

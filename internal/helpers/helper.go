package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckAuthHeader(ctx *gin.Context) (string, bool) {
	authHeader := strings.Split(ctx.GetHeader("Authorization"), " ")

	if len(authHeader) < 2 || authHeader[1] == "" {
		return "", false
	}

	return authHeader[1], true
}

func RequestToGithub(token string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func ServiceChecker(service string) error {
	var err error
	switch service {
	case "github":
		err = nil
	case "app":
		err = nil
	default:
		err = errors.New("unknown service")
	}

	return err
}

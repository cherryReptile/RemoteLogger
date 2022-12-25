package helpers

import (
	"bytes"
	"encoding/json"
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

func RequestToGoogle(token string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)

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

func GetAndCastToken(ctx *gin.Context) (string, error) {
	t, ok := ctx.Get("token")
	if !ok {
		return "", errors.New("cannot get token")
	}

	token, ok := t.(string)
	if !ok {
		return "", errors.New("invalid token")
	}

	return token, nil
}

func TrimJson(jsonBytes []byte) ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, jsonBytes); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func GetAndCastUserUUID(ctx *gin.Context) (string, error) {
	u, ok := ctx.Get("userUUID")
	if !ok {
		return "", errors.New("cannot user uuid")
	}

	uuid, ok := u.(string)
	if !ok {
		return "", errors.New("invalid uuid")
	}

	return uuid, nil
}

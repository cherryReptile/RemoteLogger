package appauth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type CustomClaims struct {
	UserID uint   `json:"id"`
	Login  string `json:"login"`
	jwt.RegisteredClaims
}

func ParseToken(authToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(authToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		key := []byte(os.Getenv("JWT_KEY"))

		return key, nil
	})

	return token, err
}

func GetClaims(ctx *gin.Context) (*CustomClaims, error) {
	authToken, ok := ctx.Get("token")

	if !ok {
		return nil, errors.New("auth token is missing")
	}

	token, err := ParseToken(authToken.(string))

	if err != nil {
		return nil, errors.New("error parsing token")
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		return nil, errors.New("failed to get claims")
	}

	return claims, nil
}

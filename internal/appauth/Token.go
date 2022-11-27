package appauth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pavel-one/GoStarter/internal/models"
	"os"
	"time"
)

type CustomClaims struct {
	UserID uint   `json:"id"`
	Login  string `json:"login"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.AppUser) (string, error) {
	claims := CustomClaims{
		UserID: user.ID,
		Login:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
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

func GetClaims(authToken string) (*CustomClaims, error) {
	token, err := ParseToken(authToken)

	if err != nil {
		return nil, errors.New("error parsing token")
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok || !token.Valid {
		return nil, errors.New("failed to get claims")
	}

	return claims, nil
}

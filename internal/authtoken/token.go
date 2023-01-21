package authtoken

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	//Unique   string `json:"unique"`
	//Service  string `json:"service"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		//Unique:  unique,
		//Service: service,
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
		if token != nil {
			claims, ok := token.Claims.(*CustomClaims)
			if !ok {
				return nil, err
			}
			return claims, err
		}
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok || !token.Valid {
		return nil, errors.New("failed to get claims")
	}

	return claims, nil
}

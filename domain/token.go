package domain

import (
	"time"
)

type AuthToken struct {
	ID        uint      `json:"id" db:"id""`
	Token     string    `json:"access_token" db:"token"`
	UserUUID  string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type AuthTokenRepo interface {
	Create(token *AuthToken) error
	GetByToken(token *AuthToken, tokenStr string) error
	Delete(token *AuthToken) error
}

type AuthTokenUsecase interface {
	Create(token *AuthToken) error
	GetByToken(token *AuthToken, tokenStr string) error
	Delete(token *AuthToken) error
}

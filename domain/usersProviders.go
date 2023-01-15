package domain

import "time"

type UsersProviders struct {
	ID         uint      `json:"id" db:"id"`
	UserID     string    `json:"user_id" db:"user_id"`
	ProviderID uint      `json:"provider_id" db:"provider_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type UsersProvidersRepo interface {
	Create(up *UsersProviders, userUUID string, providerID uint) error
}

type UsersProvidersUsecase interface {
	Create(up *UsersProviders, userUUID string, providerID uint) error
}

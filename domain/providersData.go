package domain

import (
	"encoding/json"
	"time"
)

type ProvidersData struct {
	ID         uint            `json:"id" db:"id"`
	UserData   json.RawMessage `json:"user_data" db:"user_data"`
	UserID     string          `json:"user_id" db:"user_id"`
	ProviderID uint            `json:"provider_id" db:"provider_id"`
	Username   string          `json:"username" db:"username"`
	CreatedAt  time.Time       `json:"create_at" db:"created_at"`
}

type ProvidersDataRepo interface {
	Create(pd *ProvidersData) error
	FindByUsernameAndProvider(pd *ProvidersData, username string, providerID uint) error
	GetAllByProvider(userUUID string, providerID uint) ([]ProvidersData, error)
}

type ProvidersDataUsecase interface {
	Create(pd *ProvidersData) error
	FindByUsernameAndProvider(pd *ProvidersData, username string, providerID uint) error
	GetAllByProvider(userUUID string, providerID uint) ([]ProvidersData, error)
}

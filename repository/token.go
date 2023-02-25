package repository

import (
	"errors"
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/jmoiron/sqlx"
	"time"
)

type tokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) domain.AuthTokenRepo {
	return &tokenRepository{
		db: db,
	}
}

func (r *tokenRepository) Create(token *domain.AuthToken) error {
	create := `INSERT INTO access_tokens (token, user_id, created_at) 
								VALUES (:token, :user_id, :created_at)`
	get := "SELECT * FROM access_tokens ORDER BY id DESC LIMIT 1"
	token.CreatedAt = time.Now()

	_, err := r.db.NamedExec(create, token)

	if err != nil {
		return errors.New("failed to create token " + err.Error())
	}

	// update model
	return r.db.Get(token, get)
}

func (r *tokenRepository) GetByToken(token *domain.AuthToken, tokenStr string) error {
	if err := r.db.Get(token, "SELECT * FROM access_tokens WHERE token=$1 ORDER BY id DESC LIMIT 1", tokenStr); err != nil {
		return err
	}

	return nil
}

func (r *tokenRepository) Delete(token *domain.AuthToken) error {
	_, err := r.db.NamedExec("DELETE FROM access_tokens WHERE id=:id", token)
	return err
}

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

func (t *tokenRepository) Create(token *domain.AuthToken) error {
	token.CreatedAt = time.Now()

	_, err := t.db.NamedExec(`INSERT INTO access_tokens (token, user_id, created_at) 
								VALUES (:token, :user_id, :created_at)`, token)

	if err != nil {
		return errors.New("failed to create token " + err.Error())
	}

	// update model
	if err = t.db.Get(token, "SELECT * FROM access_tokens ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}

func (t *tokenRepository) GetByToken(token *domain.AuthToken, tokenStr string) error {
	if err := t.db.Get(token, "SELECT * FROM access_tokens WHERE token=$1 ORDER BY id DESC LIMIT 1", tokenStr); err != nil {
		return err
	}

	return nil
}

func (t *tokenRepository) Delete(token *domain.AuthToken) error {
	_, err := t.db.NamedExec("DELETE FROM access_tokens WHERE id=:id", token)
	return err
}

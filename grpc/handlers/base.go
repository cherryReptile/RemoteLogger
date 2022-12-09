package handlers

import (
	"github.com/jmoiron/sqlx"
)

type BaseDB struct {
	DB *sqlx.DB
}

package controllers

import "github.com/jmoiron/sqlx"

type BaseController struct {
}

type DatabaseController struct {
	DB *sqlx.DB
}

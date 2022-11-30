package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/internal/sqlite"
	"os"
)

type BaseModel struct {
}

func (m *BaseModel) CheckDb(service, unique string) (*sqlx.DB, error) {
	path := "./storage/users/" + service + "/" + unique
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return nil, err
	}

	db, err := sqlite.GetDb(service, unique)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (m *BaseModel) createSubDir(service, unique string) (*sqlx.DB, error) {
	path := "./storage/users/" + service + "/" + unique
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, err
	}

	db, err := sqlite.GetDb(service, unique)
	if err != nil {
		return nil, err
	}

	err = sqlite.SetDefaultSchema(db, service)
	if err != nil {
		return nil, err
	}

	return db, nil
}

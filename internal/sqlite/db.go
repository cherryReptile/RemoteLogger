package sqlite

import (
	_ "embed"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed github.sql
var githubSchema string

//go:embed app.sql
var appSchema string

//go:embed google.sql
var googleSchema string

//go:embed telegram.sql
var telegramSchema string

//go:embed token.sql
var tokenSchema string

func GetDb(authService, userIdentifier string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("sqlite3", "./storage/users/"+authService+"/"+userIdentifier+"/"+userIdentifier+".sqlite3")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

func SetDefaultSchema(db *sqlx.DB, schema string) (err error) {
	switch schema {
	case "github":
		_, err = db.Exec(githubSchema + tokenSchema)
	case "app":
		_, err = db.Exec(appSchema + tokenSchema)
	case "google":
		_, err = db.Exec(googleSchema + tokenSchema)
	case "telegram":
		_, err = db.Exec(telegramSchema)
	default:
		err = errors.New("unknown service")
	}
	if err != nil {
		return err
	}

	return err
}

package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Rollback(err error, tx *sqlx.Tx) error {
	if back := tx.Rollback(); back != nil {
		logrus.Error(back)
		return fmt.Errorf("query error: %v\nrollback error: %v\n", err, back)
	}
	return err
}

package db

import (
	"database/sql"

	"github.com/Rellum/goku"
)

//go:generate glean -table=data --err_no_rows=github.com/Rellum/goku.ErrNotFound -scan -raw

type glean struct {
	goku.KV

	DeletedRef sql.NullInt64
	LeaseID    sql.NullInt64
}

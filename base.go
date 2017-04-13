package pqx

import (
	"database/sql"
	"github.com/maprost/pqx/pqarg"
)

type queryRowFunc func(sql string, args pqarg.Args) *sql.Row

type queryFunc func(sql string, args pqarg.Args) (*sql.Rows, error)

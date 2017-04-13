package pqx

import (
	"database/sql"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqutil"
)

func Query(sql string, args pqarg.Args) (*sql.Rows, error) {
	return LogQuery(sql, args, pqutil.DefaultLogger)
}

func LogQuery(sql string, args pqarg.Args, logger pqdep.Logger) (rows *sql.Rows, err error) {
	pqlib.LogQueryFunc(
		func(sql string, args ...interface{}) {
			rows, err = DB.Query(sql, args...)
		}, logger, sql, args)

	return
}

func QueryRow(sql string, args pqarg.Args) *sql.Row {
	return LogQueryRow(sql, args, pqutil.DefaultLogger)
}

func LogQueryRow(sql string, args pqarg.Args, logger pqdep.Logger) (row *sql.Row) {
	pqlib.LogQueryFunc(
		func(sql string, args ...interface{}) {
			row = DB.QueryRow(sql, args...)
		}, logger, sql, args)

	return
}

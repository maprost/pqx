package pqx

import (
	"database/sql"
	"time"

	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
)

type queryFunc func(sql string, args pqarg.Args) (*sql.Rows, error)

func logWrapper(queryFunc func(sql string, args ...interface{}), sql string, args pqarg.Args, logger pqdep.Logger) {
	// track duration
	start := time.Now()
	// execute
	queryFunc(sql, args.Get()...)
	// log sql + duration
	elapsed := time.Now().Sub(start)
	logger.Printf("[time: "+elapsed.String()+"] SQL: "+sql, args.Get()...)
}

func queryFuncWrapper(logger pqdep.Logger) queryFunc {
	return func(sql string, args pqarg.Args) (*sql.Rows, error) {
		return LogQuery(sql, args, logger)
	}
}

func closeRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}

package pqlib

import (
	"database/sql"
	"github.com/mleuth/pqlib/pqdep"
	"github.com/mleuth/pqlib/pqutil"
	"github.com/mleuth/timeutil"
)

func Query(sql string, args Args) (Result, error) {
	return queryFunc(db.Query, pqutil.DefaultLogger, sql, args)
}

func QueryLg(logger pqdep.Logger, sql string, args Args) (Result, error) {
	return queryFunc(db.Query, logger, sql, args)
}

func queryFunc(qFunc func(query string, args ...interface{}) (*sql.Rows, error), logger pqdep.Logger, sql string, args Args) (Result, error) {
	// track duration
	stopwatch := timeutil.NewStopwatch()
	// execute
	rows, e := qFunc(sql, args.get()...)
	// log sql + duration
	stopwatch.Stop()
	logger.Printf("[time: "+stopwatch.String()+"] SQL: "+sql, args.get()...)
	if e != nil {
		return Result{}, e
	}

	return Result{rows: rows, hasNext: false}, nil
}

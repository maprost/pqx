package pqx

import (
	"github.com/maprost/pqlib/pqarg"
	"github.com/maprost/pqlib/pqdep"
	"github.com/maprost/pqlib/pqutil"
	"time"
)

func QueryFunc(queryFunc func(sql string, args ...interface{}), sql string, args pqarg.Args) {
	LogQueryFunc(queryFunc, pqutil.DefaultLogger, sql, args)
}

func LogQueryFunc(queryFunc func(sql string, args ...interface{}), logger pqdep.Logger, sql string, args pqarg.Args) {
	// track duration
	start := time.Now()
	// execute
	queryFunc(sql, args.Get()...)
	// log sql + duration
	elapsed := time.Now().Sub(start)
	logger.Printf("[time: "+elapsed.String()+"] SQL: "+sql, args.Get()...)
}

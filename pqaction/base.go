package pqaction

import (
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqdep"
	"log"
	"os"
)

var defaultLogger = log.New(os.Stdout, "", 0)

type queryFunc func(sql string, args pqlib.Args) (pqlib.Result, error)

func queryFuncWrapper(logger pqdep.Logger) queryFunc {
	return func(sql string, args pqlib.Args) (pqlib.Result, error) {
		return pqlib.Query(logger, sql, args)
	}
}

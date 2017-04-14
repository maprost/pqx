package pqx

import (
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqutil"
	"github.com/maprost/pqx/pqutil/pqreflect"
)

// Register entities via pqx.LogQueryRow and use a default logger for logging.
//
func Register(entity ...interface{}) error {
	return LogRegister(pqutil.DefaultLogger, entity...)
}

// LogRegister register entities via pqx.LogQueryRow and use the given pqdep.Logger for logging.
func LogRegister(logger pqdep.Logger, entity ...interface{}) error {
	return registerListFunc(queryFuncWrapper(logger), entity)
}

// Register entities via tx.QueryRow and use the given tx.log for logging.
func (tx *Transaction) Register(entity ...interface{}) error {
	return registerListFunc(tx.Query, entity)
}

func registerListFunc(qFunc queryFunc, entities []interface{}) error {
	for _, entity := range entities {
		e := registerFunc(qFunc, entity)
		if e != nil {
			return e
		}
	}

	return nil
}

func registerFunc(qFunc queryFunc, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	exists, e := tableExists(qFunc, structInfo)
	if e != nil {
		return e
	}
	if exists {
		return nil
	}

	e = createFunc(qFunc, entity)
	return e
}

func tableExists(qFunc queryFunc, structInfo pqreflect.StructInfo) (bool, error) {
	args := pqarg.New()
	sql := "SELECT table_name FROM INFORMATION_SCHEMA.TABLES " +
		"WHERE TABLE_NAME = " + args.Next(structInfo.Name())
	rows, e := qFunc(sql, args)
	defer closeRows(rows)
	if e != nil {
		return false, e
	}

	// no table with that name is known
	if rows.Next() == false {
		return false, nil
	}

	// TODO: check columns

	// table exists and there is no error
	return true, nil
}

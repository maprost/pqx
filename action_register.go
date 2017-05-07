package pqx

import (
	"fmt"

	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqutil"
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
	table, err := pqtable.New(entity)
	if err != nil {
		return err
	}

	exists, err := tableExists(qFunc, table)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	err = createFunc(qFunc, entity)
	return err
}

func tableExists(qFunc queryFunc, table *pqtable.Table) (bool, error) {
	args := pqarg.New()
	sql := "SELECT table_name FROM INFORMATION_SCHEMA.TABLES " +
		"WHERE TABLE_NAME = " + args.Next(table.Name())
	rows, err := qFunc(sql, args)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// no table with that name is known
	if rows.Next() == false {
		return false, nil
	}

	// table exists -> check columns
	err = checkColumns(qFunc, table)
	return true, err
}

func checkColumns(qFunc queryFunc, table *pqtable.Table) error {
	sql := "Select * FROM " + table.Name() + ";"
	rows, err := qFunc(sql, pqarg.New())
	if err != nil {
		return err
	}
	defer rows.Close()

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	for _, c := range colTypes {
		fmt.Println("Column:", c.Name())
	}

	return nil
}

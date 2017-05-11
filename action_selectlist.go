package pqx

import (
	"errors"

	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqutil"
)

// Select a list of entities via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name WHERE key = value
func SelectListByKeyValue(key string, value interface{}, prototype interface{}, appendPrototypeToList func()) error {
	return LogSelectListByKeyValue(key, value, prototype, appendPrototypeToList, pqutil.DefaultLogger)
}

// LogSelect select a list of entities via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name WHERE key = value
func LogSelectListByKeyValue(key string, value interface{}, prototype interface{}, appendPrototypeToList func(), logger pqdep.Logger) error {
	return selectListByKeyValueFunc(queryFuncWrapper(logger), key, value, prototype, appendPrototypeToList)
}

// Select a list of entities via tx.LogQuery and use a tx.log for logging.
// SELECT column1, column2,... FROM table_name WHERE key = value
func (tx *Transaction) SelectListByKeyValue(key string, value interface{}, prototype interface{}, appendPrototypeToList func()) error {
	return selectListByKeyValueFunc(tx.Query, key, value, prototype, appendPrototypeToList)
}

// SELECT column1, column2,... FROM table_name WHERE key = value
func selectListByKeyValueFunc(qFunc queryFunc, key string, value interface{}, prototype interface{}, appendPrototypeToList func()) error {
	table, err := pqtable.New(prototype)
	if err != nil {
		return err
	}
	if table.IsPointer() == false {
		return errors.New("Struct must be given as pointer/reference.")
	}

	args := pqarg.New()
	sql := "SELECT " + selectRowList(table, "") +
		" FROM " + table.Name() +
		" WHERE " + key + " = " + args.Next(value)

	rows, err := qFunc(sql, args)
	defer closeRows(rows)
	if err != nil {
		return err
	}
	return ScanTableToList(rows, table, appendPrototypeToList)
}

// Select a list of all entities via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name WHERE key = value
func SelectList(prototype interface{}, appendPrototypeToList func()) error {
	return LogSelectList(prototype, appendPrototypeToList, pqutil.DefaultLogger)
}

// LogSelect select a list of all entities via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name
func LogSelectList(prototype interface{}, appendPrototypeToList func(), logger pqdep.Logger) error {
	return selectListFunc(queryFuncWrapper(logger), prototype, appendPrototypeToList)
}

// Select a list of all entities via tx.LogQuery and use a tx.log for logging.
// SELECT column1, column2,... FROM table_name
func (tx *Transaction) SelectList(prototype interface{}, appendPrototypeToList func()) error {
	return selectListFunc(tx.Query, prototype, appendPrototypeToList)
}

// SELECT column1, column2,... FROM table_name
func selectListFunc(qFunc queryFunc, prototype interface{}, appendPrototypeToList func()) error {
	table, err := pqtable.New(prototype)
	if err != nil {
		return err
	}
	if table.IsPointer() == false {
		return errors.New("Struct must be given as pointer/reference.")
	}

	args := pqarg.New()
	sql := "SELECT " + selectRowList(table, "") + " FROM " + table.Name()

	rows, err := qFunc(sql, args)
	defer closeRows(rows)
	if err != nil {
		return err
	}
	return ScanTableToList(rows, table, appendPrototypeToList)
}

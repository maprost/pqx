package pqx

import (
	"errors"

	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqutil"
)

// Select an entity via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name WHERE PK = valueX (with PK tag)
func Select(entity interface{}) (bool, error) {
	return LogSelect(entity, pqutil.DefaultLogger)
}

// LogSelect select an entity via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name WHERE PK = valueX (with PK tag)
func LogSelect(entity interface{}, logger pqdep.Logger) (bool, error) {
	return prepareSelect(queryFuncWrapper(logger), entity)
}

// Select an entity via tx.LogQuery and use a tx.log for logging.
// SELECT column1, column2,... FROM table_name WHERE PK = valueX (with PK tag)
func (tx *Transaction) Select(entity interface{}) (bool, error) {
	return prepareSelect(tx.Query, entity)
}

// SELECT column1, column2,... FROM table_name WHERE PK = valueX (with PK tag)
func prepareSelect(qFunc queryFunc, entity interface{}) (bool, error) {
	table, err := pqtable.New(entity)
	if err != nil {
		return false, err
	}

	// search for key
	for _, column := range table.Columns() {
		if column.PrimaryKeyTag() {
			return selectFunc(qFunc, table, column.Name(), column.GetValue())
		}
	}
	return false, errors.New("No primary key available.")

}

// Select an entity via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name WHERE key = value
func SelectByKeyValue(key string, value interface{}, entity interface{}) (bool, error) {
	return LogSelectByKeyValue(key, value, entity, pqutil.DefaultLogger)
}

// LogSelect select an entity via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name WHERE key = value
func LogSelectByKeyValue(key string, value interface{}, entity interface{}, logger pqdep.Logger) (bool, error) {
	return prepareSelectByKeyValue(queryFuncWrapper(logger), key, value, entity)
}

// Select an entity via tx.LogQuery and use a tx.log for logging.
// SELECT column1, column2,... FROM table_name WHERE key = value
func (tx *Transaction) SelectByKeyValue(key string, value interface{}, entity interface{}) (bool, error) {
	return prepareSelectByKeyValue(tx.Query, key, value, entity)
}

// SELECT column1, column2,... FROM table_name WHERE key = value
func prepareSelectByKeyValue(qFunc queryFunc, key string, value interface{}, entity interface{}) (bool, error) {
	table, err := pqtable.New(entity)
	if err != nil {
		return false, err
	}
	return selectFunc(qFunc, table, key, value)
}

// SELECT column1, column2,... FROM table_name WHERE key = value
func selectFunc(qFunc queryFunc, table *pqtable.Table, key string, value interface{}) (bool, error) {
	if table.IsPointer() == false {
		return false, errors.New("Struct must be given as pointer/reference.")
	}

	args := pqarg.New()
	sql := "SELECT " + selectRowList(table, "") +
		" FROM " + table.Name() +
		" WHERE " + key + " = " + args.Next(value)

	rows, err := qFunc(sql, args)
	defer closeRows(rows)
	if err != nil {
		return false, err
	}

	if rows.Next() == false {
		return false, errors.New("No row available to scan.")
	}

	err = ScanTable(rows, table)
	return err == nil, err
}

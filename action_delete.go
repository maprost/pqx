package pqx

import (
	"errors"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqutil"
)

// Delete an entity via pqx.LogQuery and use a default logger for logging.
// DELETE FROM table_name
// WHERE PK = value;
func Delete(entity interface{}) error {
	return LogDelete(entity, pqutil.DefaultLogger)
}

// LogDelete delete an entity via pqx.LogQuery and use the given pqdep.Logger for logging.
// DELETE FROM table_name
// WHERE PK = value;
func LogDelete(entity interface{}, logger pqdep.Logger) error {
	return prepareDelete(queryFuncWrapper(logger), entity)
}

// Delete an entity via tx.Query and use the tx.log for logging.
// DELETE FROM table_name
// WHERE PK = value;
func (tx *Transaction) Delete(entity interface{}) error {
	return prepareDelete(tx.Query, entity)
}

// DELETE FROM table_name
// WHERE PK = value;
func prepareDelete(qFunc queryFunc, entity interface{}) error {
	table, err := pqtable.NewCtx(entity, pqtable.Context{OnlyPrimaryKeyColumn: true})
	if err != nil {
		return err
	}

	// search for key
	for _, column := range table.Columns() {
		if column.PrimaryKeyTag() {
			return deleteFunc(qFunc, table, column.Name(), column.GetValue())
		}
	}
	return errors.New("No primary key available.")
}

// DELETE FROM table_name
// WHERE key = value;
func deleteFunc(qFunc queryFunc, table *pqtable.Table, key string, value interface{}) error {
	args := pqarg.New()
	sql := "DELETE FROM " + table.Name() + " WHERE " + key + " = " + args.Next(value)
	rows, err := qFunc(sql, args)
	closeRows(rows)
	return err
}

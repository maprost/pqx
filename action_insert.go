package pqx

import (
	"errors"

	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqtime"
	"github.com/maprost/pqx/pqutil"
)

// Insert an entity via pqx.LogQuery and use a default logger for logging.
// INSERT INTO table_name (AI, column1,column2,column3,...)
// VALUES (DEFAULT, value1,value2,value3,...) RETURNING AI;
func Insert(entity interface{}) error {
	return LogInsert(entity, pqutil.DefaultLogger)
}

// LogInsert insert an entity via pqx.LogQuery and use the given pqdep.Logger for logging.
// INSERT INTO table_name (AI, column1,column2,column3,...)
// VALUES (DEFAULT, value1,value2,value3,...) RETURNING AI;
func LogInsert(entity interface{}, logger pqdep.Logger) error {
	return insertFunc(queryFuncWrapper(logger), entity)
}

// Insert an entity via tx.Query and use the given tx.log for logging.
// INSERT INTO table_name (AI, column1,column2,column3,...)
// VALUES (DEFAULT, value1,value2,value3,...) RETURNING AI;
func (tx *Transaction) Insert(entity interface{}) error {
	return insertFunc(tx.Query, entity)
}

// INSERT INTO table_name (AI, column1,column2,column3,...)
// VALUES (DEFAULT, value1,value2,value3,...) RETURNING AI
func insertFunc(qfunc queryFunc, entity interface{}) (err error) {
	table, err := pqtable.New(entity)
	if err != nil {
		return err
	}

	if table.IsPointer() == false {
		err = errors.New("Struct must be given as pointer/reference.")
		return
	}

	columns := ""
	values := ""
	returning := ""
	args := pqarg.New()
	var autoIncrement pqtable.Column

	// preparation of the statement
	for _, column := range table.Columns() {
		columns = pqutil.Concate(columns, column.Name(), ",")

		if column.AutoIncrementTag() {
			returning = "RETURNING " + column.Name()
			values = pqutil.Concate(values, "DEFAULT", ",")
			autoIncrement = column
		} else {
			if column.CreateDateTag() || column.ChangeDateTag() {
				column.SetTime(pqtime.Now())
			}

			values = pqutil.Concate(values, args.Next(column.GetValue()), ",")
		}
	}

	// execute statement
	sql := "INSERT INTO " + table.Name() + " (" + columns + ") VALUES (" + values + ")" + returning
	rows, err := qfunc(sql, args)
	defer closeRows(rows)
	if err != nil {
		return err
	}
	// update pk with returning value (if needed)
	if len(returning) > 0 {
		if rows.Next() == false {
			return errors.New("No return element in insert statement.")
		}

		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return err
		}
		// add id to pk
		autoIncrement.SetInt64(id)
	}

	return nil
}

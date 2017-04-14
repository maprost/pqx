package pqx

import (
	"errors"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqutil"
	"github.com/maprost/pqx/pqutil/pqreflect"
	"github.com/maprost/timeutil"
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
	structInfo := pqreflect.NewStructInfo(entity)

	columns := ""
	values := ""
	returning := ""
	args := pqarg.New()
	var autoIncrement pqreflect.Field

	// preparation of the statement
	for _, field := range structInfo.Fields() {
		columns = pqutil.Concate(columns, field.Name(), ",")

		if isAutoIncrement(field) {
			returning = "RETURNING " + field.Name()
			values = pqutil.Concate(values, "DEFAULT", ",")
			autoIncrement = field
		} else {
			if isCreateDate(field) || isChangeDate(field) {
				field.SetTime(timeutil.Now())
			}

			values = pqutil.Concate(values, args.Next(field.GetValue()), ",")
		}
	}

	// execute statement
	sql := "INSERT INTO " + structInfo.Name() + " (" + columns + ") VALUES (" + values + ")" + returning
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
		autoIncrement.SetInt(id)
	}

	return nil
}

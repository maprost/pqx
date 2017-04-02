package pqaction

import (
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqutil"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
	"github.com/mleuth/timeutil"
)

// INSERT INTO table_name (AI, column1,column2,column3,...)
// VALUES (DEFAULT, value1,value2,value3,...) RETURNING AI;
func Insert(tx pqlib.Transaction, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	columns := ""
	values := ""
	returning := ""
	args := pqlib.NewArgs()
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
	result, e := tx.Query(sql, args)
	if e != nil {
		return e
	}

	// update pk with returning value (if needed)
	if len(returning) > 0 {
		var id int64
		e = result.Scan(&id)
		if e != nil {
			return e
		}
		// add id to pk
		autoIncrement.SetInt(id)
	}

	return nil
}

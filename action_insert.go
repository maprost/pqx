package pqx

import (
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqutil"
	"github.com/maprost/pqx/pqutil/pqreflect"
	"github.com/maprost/timeutil"
)

// INSERT INTO table_name (AI, column1,column2,column3,...)
// VALUES (DEFAULT, value1,value2,value3,...) RETURNING AI
func Insert(qfunc queryRowFunc, entity interface{}) (err error) {
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
	row := qfunc(sql, args)

	// update pk with returning value (if needed)
	if len(returning) > 0 {
		var id int64
		err = row.Scan(&id)
		if err != nil {
			return
		}
		// add id to pk
		autoIncrement.SetInt(id)
	}

	return nil
}

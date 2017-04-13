package pqx

import (
	"errors"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqutil"
	"github.com/maprost/pqx/pqutil/pqreflect"
	"github.com/maprost/timeutil"
)

// UPDATE table_name
// SET column1 = value1,
// column2 = value2,
// ...
// WHERE PK = valueX (with PK tag)
func Update(qFunc queryFunc, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	sets := ""
	args := pqarg.New()
	whereClause := ""

	// preparation of the statement
	for _, field := range structInfo.Fields() {
		if isCreateDate(field) {
			// don't change the create date
			continue
		}

		if isPrimaryKey(field) {
			whereClause = field.Name() + " = " + args.Next(field.GetValue())
		} else {
			if isChangeDate(field) {
				field.SetTime(timeutil.Now())
			}

			sets = pqutil.Concate(sets, field.Name()+" = "+args.Next(field.GetValue()), ",")
		}
	}

	if len(whereClause) == 0 {
		return errors.New("No primary key available.")
	}

	// execute statement
	sql := "UPDATE " + structInfo.Name() + " SET " + sets + " WHERE " + whereClause
	_, err := qFunc(sql, args)
	return err
}

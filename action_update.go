package pqx

import (
	"errors"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqutil"
	"github.com/maprost/pqx/pqutil/pqreflect"
	"github.com/maprost/timeutil"
)

// Update an entity via pqx.LogQuery and use a default logger for logging.
// UPDATE table_name
// SET column1 = value1,
// column2 = value2,
// ...
// WHERE PK = valueX (with PK tag)
func Update(entity interface{}) error {
	return LogUpdate(entity, pqutil.DefaultLogger)
}

// LogUpdate update an entity via pqx.LogQuery and use the given pqdep.Logger for logging.
// UPDATE table_name
// SET column1 = value1,
// column2 = value2,
// ...
// WHERE PK = valueX (with PK tag)
func LogUpdate(entity interface{}, logger pqdep.Logger) error {
	return updateFunc(queryFuncWrapper(logger), entity)
}

// Update an entity via tx.Query and use the tx.log for logging.
// UPDATE table_name
// SET column1 = value1,
// column2 = value2,
// ...
// WHERE PK = valueX (with PK tag)
func (tx *Transaction) Update(entity interface{}) error {
	return updateFunc(tx.Query, entity)
}

// UPDATE table_name
// SET column1 = value1,
// column2 = value2,
// ...
// WHERE PK = valueX (with PK tag)
func updateFunc(qFunc queryFunc, entity interface{}) error {
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
	rows, err := qFunc(sql, args)
	closeRows(rows)
	return err
}

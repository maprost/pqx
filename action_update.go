package pqx

import (
	"errors"
	"github.com/maprost/timeutil"

	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqutil"
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
	table, err := pqtable.New(entity)
	if err != nil {
		return err
	}

	if table.IsPointer() == false {
		return errors.New("Struct must be given as pointer/reference.")
	}

	sets := ""
	args := pqarg.New()
	whereClause := ""

	// preparation of the statement
	for _, column := range table.Columns() {
		if column.CreateDateTag() {
			// don't change the create date
			continue
		}

		if column.PrimaryKeyTag() {
			whereClause = column.Name() + " = " + args.Next(column.GetValue())
		} else {
			if column.ChangeDateTag() {
				column.SetTime(timeutil.Now())
			}

			sets = pqutil.Concate(sets, column.Name()+" = "+args.Next(column.GetValue()), ",")
		}
	}

	if len(whereClause) == 0 {
		return errors.New("No primary key available.")
	}

	// execute statement
	sql := "UPDATE " + table.Name() + " SET " + sets + " WHERE " + whereClause
	rows, err := qFunc(sql, args)
	closeRows(rows)
	return err
}

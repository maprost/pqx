package pqx

import (
	"errors"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqutil"
	"github.com/maprost/pqx/pqutil/pqreflect"
)

// Contains search for an entity via pqx.LogQuery and use a default logger for logging.
// SELECT PK FROM table_name WHERE PK = valueX (with PK tag)
func Contains(entity interface{}) (bool, error) {
	return LogContains(entity, pqutil.DefaultLogger)
}

// LogContains search for an entity via pqx.LogQueryRow and use the given pqdep.Logger for logging.
// SELECT PK FROM table_name WHERE PK = valueX (with PK tag)
func LogContains(entity interface{}, logger pqdep.Logger) (bool, error) {
	return prepareContains(queryFuncWrapper(logger), entity)
}

// Contains search for an entity via tx.QueryRow and use the given tx.log for logging.
// SELECT PK FROM table_name WHERE PK = valueX (with PK tag)
func (tx *Transaction) Contains(entity interface{}) (bool, error) {
	return prepareContains(tx.Query, entity)
}

// prepareContains looking for the PK of the entity
// SELECT PK FROM table_name WHERE PK = valueX (with PK tag)
func prepareContains(qFunc queryFunc, entity interface{}) (bool, error) {
	structInfo := pqreflect.NewStructInfo(entity)

	// search for key
	for _, field := range structInfo.Fields() {
		if isPrimaryKey(field) {
			return containsFunc(qFunc, structInfo, field.Name(), field.GetValue())
		}
	}
	return false, errors.New("No primary key available.")
}

// SELECT PK FROM table_name WHERE PK = valueX (with PK tag)
func containsFunc(qFunc queryFunc, structInfo pqreflect.StructInfo, key string, value interface{}) (bool, error) {
	args := pqarg.New()
	sql := "Select " + key +
		" FROM " + structInfo.Name() +
		" WHERE " + key + " = " + args.Next(value)

	rows, err := qFunc(sql, args)
	defer closeRows(rows)
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}

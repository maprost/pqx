package pqx

import (
	"errors"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqutil"
	"github.com/maprost/pqx/pqutil/pqreflect"
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
	structInfo := pqreflect.NewStructInfo(entity)

	// search for key
	for _, field := range structInfo.Fields() {
		if isPrimaryKey(field) {
			return selectFunc(qFunc, structInfo, entity, field.Name(), field.GetValue())
		}
	}
	return false, errors.New("No primary key available.")

}

// Select an entity via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name WHERE PK = valueX (with PK tag)
func SelectByKeyValue(key string, value interface{}, entity interface{}) (bool, error) {
	return LogSelectByKeyValue(key, value, entity, pqutil.DefaultLogger)
}

// LogSelect select an entity via pqx.LogQuery and use a default logger for logging.
// SELECT column1, column2,... FROM table_name WHERE PK = valueX (with PK tag)
func LogSelectByKeyValue(key string, value interface{}, entity interface{}, logger pqdep.Logger) (bool, error) {
	return prepareSelectByKeyValue(queryFuncWrapper(logger), key, value, entity)
}

// Select an entity via tx.LogQuery and use a tx.log for logging.
// SELECT column1, column2,... FROM table_name WHERE PK = valueX (with PK tag)
func (tx *Transaction) SelectByKeyValue(key string, value interface{}, entity interface{}) (bool, error) {
	return prepareSelectByKeyValue(tx.Query, key, value, entity)
}

// SELECT column1, column2,... FROM table_name WHERE key = value
func prepareSelectByKeyValue(qFunc queryFunc, key string, value interface{}, entity interface{}) (bool, error) {
	structInfo := pqreflect.NewStructInfo(entity)
	return selectFunc(qFunc, structInfo, entity, key, value)
}

// SELECT column1, column2,... FROM table_name WHERE key = value
func selectFunc(qFunc queryFunc, structInfo pqreflect.StructInfo, entity interface{}, key string, value interface{}) (bool, error) {
	args := pqarg.New()
	sql := "Select " + selectList(structInfo, "") +
		" FROM " + structInfo.Name() +
		" WHERE " + key + " = " + args.Next(value)

	rows, err := qFunc(sql, args)
	defer closeRows(rows)
	if err != nil {
		return false, err
	}

	if rows.Next() == false {
		return false, nil
	}

	err = ScanStruct(rows, entity)
	return err == nil, err
}

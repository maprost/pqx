package pqaction

import (
	"errors"
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqdep"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
)

// DELETE FROM table_name
// WHERE PK = value;
func Delete(entity interface{}) error {
	return DeleteLg(entity, defaultLogger)
}

// DELETE FROM table_name
// WHERE PK = value;
func DeleteLg(entity interface{}, logger pqdep.Logger) error {
	return delete(queryFuncWrapper(logger), entity)
}

// DELETE FROM table_name
// WHERE PK = value;
func DeleteTx(tx pqlib.Transaction, entity interface{}) error {
	return delete(tx.Query, entity)
}

// DELETE FROM table_name
// WHERE PK = value;
func delete(qFunc queryFunc, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	// search for key
	for _, field := range structInfo.Fields() {
		if isPrimaryKey(field) {
			return deleteFunc(qFunc, structInfo, field.Name(), field.GetValue())
		}
	}
	return errors.New("No primary key available.")
}

// DELETE FROM table_name
// WHERE key = value;
func DeleteByKeyValue(key string, entity interface{}) error {
	return DeleteByKeyValueLg(key, entity, defaultLogger)
}

// DELETE FROM table_name
// WHERE key = value;
func DeleteByKeyValueLg(key string, entity interface{}, logger pqdep.Logger) error {
	return deleteByKeyValue(queryFuncWrapper(logger), key, entity)
}

// DELETE FROM table_name
// WHERE key = value;
func DeleteByKeyValueTx(tx pqlib.Transaction, key string, entity interface{}) error {
	return deleteByKeyValue(tx.Query, key, entity)
}

// DELETE FROM table_name
// WHERE key = value;
func deleteByKeyValue(qFunc queryFunc, key string, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	// search for key
	for _, field := range structInfo.Fields() {
		if field.Name() == key {
			return deleteFunc(qFunc, structInfo, field.Name(), field.GetValue())
		}
	}
	return errors.New("Not supported field: " + key + ".")
}

// DELETE FROM table_name
// WHERE key = value;
func deleteFunc(qFunc queryFunc, structInfo pqreflect.StructInfo, key string, value interface{}) error {
	args := pqlib.NewArgs()
	sql := "DELETE FROM " + structInfo.Name() + " WHERE " + key + " = " + args.Next(value)
	_, e := qFunc(sql, args)
	return e
}

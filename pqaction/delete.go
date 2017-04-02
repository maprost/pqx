package pqaction

import (
	"errors"
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
)

// DELETE FROM table_name
// WHERE PK = value;
func Delete(tx pqlib.Transaction, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	// search for key
	for _, field := range structInfo.Fields() {
		if isPrimaryKey(field) {
			return deleteByKeyValue(tx, structInfo, field.Name(), field.GetValue())
		}
	}
	return errors.New("No primary key available.")
}

// DELETE FROM table_name
// WHERE key = value;
func DeleteByKeyValue(tx pqlib.Transaction, key string, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	// search for key
	for _, field := range structInfo.Fields() {
		if field.Name() == key {
			return deleteByKeyValue(tx, structInfo, field.Name(), field.GetValue())
		}
	}
	return errors.New("Not supported field: " + key + ".")
}

// DELETE FROM table_name
// WHERE key = value;
func deleteByKeyValue(tx pqlib.Transaction, structInfo pqreflect.StructInfo, key string, value interface{}) error {
	args := pqlib.NewArgs()
	sql := "DELETE FROM " + structInfo.Name() + " WHERE " + key + " = " + args.Next(value)
	_, e := tx.Query(sql, args)
	return e
}

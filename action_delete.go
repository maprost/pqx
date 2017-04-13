package pqx

import (
	"errors"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqutil/pqreflect"
)

// DELETE FROM table_name
// WHERE PK = value;
func Delete(qFunc queryFunc, entity interface{}) error {
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
func deleteFunc(qFunc queryFunc, structInfo pqreflect.StructInfo, key string, value interface{}) error {
	args := pqarg.New()
	sql := "DELETE FROM " + structInfo.Name() + " WHERE " + key + " = " + args.Next(value)
	_, err := qFunc(sql, args)
	return err
}

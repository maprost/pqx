package pqaction

import (
	"errors"
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
)

func SelectEntityById(tx pqlib.Transaction, entity interface{}, id int64) error {
	structInfo := pqreflect.NewStructInfo(entity)

	// search for key
	for _, field := range structInfo.Fields() {
		if isPrimaryKey(field) {
			return selectEntityByKeyValue(tx, structInfo, entity, field.Name(), id)
		}
	}
	return errors.New("No primary key available.")

}

func SelectEntityByKeyValue(tx pqlib.Transaction, entity interface{}, key string, value interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)
	return selectEntityByKeyValue(tx, structInfo, entity, key, value)
}

func selectEntityByKeyValue(tx pqlib.Transaction, structInfo pqreflect.StructInfo, entity interface{}, key string, value interface{}) error {
	args := pqlib.NewArgs()
	result, err := tx.Query(
		"Select "+selectList(structInfo, "")+
			" FROM "+structInfo.Name()+
			" WHERE "+key+" = "+args.Next(value), args)
	if err != nil {
		return err
	}

	return result.ScanStruct(entity)
}

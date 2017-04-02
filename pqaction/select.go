package pqaction

import (
	"errors"
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqutil"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
)

func SelectList(entity interface{}) string {
	structInfo := pqreflect.NewStructInfo(entity)
	return selectList(structInfo, "")
}

func SelectListWithAlias(entity interface{}, alias string) string {
	structInfo := pqreflect.NewStructInfo(entity)
	return selectList(structInfo, alias)
}

func GetSingleEntityById(tx pqlib.Transaction, entity interface{}, id int64) error {
	structInfo := pqreflect.NewStructInfo(entity)

	// search for key
	for _, field := range structInfo.Fields() {
		if isPrimaryKey(field) {
			return getSingleEntityByKeyValue(tx, structInfo, entity, field.Name(), id)
		}
	}
	return errors.New("No primary key available.")

}

func GetSingleEntityByKeyValue(tx pqlib.Transaction, entity interface{}, key string, value interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)
	return getSingleEntityByKeyValue(tx, structInfo, entity, key, value)
}

func getSingleEntityByKeyValue(tx pqlib.Transaction, structInfo pqreflect.StructInfo, entity interface{}, key string, value interface{}) error {
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

func selectList(structInfo pqreflect.StructInfo, alias string) string {
	list := ""
	if alias != "" {
		alias += "."
	}

	for _, field := range structInfo.Fields() {
		list = pqutil.Concate(list, alias+field.Name(), ",")
	}

	return list
}

func ContainsEntityById(tx pqlib.Transaction, entity interface{}, id int64) (bool, error) {
	structInfo := pqreflect.NewStructInfo(entity)

	// search for key
	for _, field := range structInfo.Fields() {
		if isPrimaryKey(field) {
			return containsEntityByKeyValue(tx, structInfo, field.Name(), id)
		}
	}
	return false, errors.New("No primary key available.")
}

func ContainsEntityByKeyValue(tx pqlib.Transaction, entity interface{}, key string, value interface{}) (bool, error) {
	structInfo := pqreflect.NewStructInfo(entity)
	return containsEntityByKeyValue(tx, structInfo, key, value)
}

func containsEntityByKeyValue(tx pqlib.Transaction, structInfo pqreflect.StructInfo, key string, value interface{}) (bool, error) {
	args := pqlib.NewArgs()
	result, e := tx.Query(
		"Select "+key+
			" FROM "+structInfo.Name()+
			" WHERE "+key+" = "+args.Next(value), args)
	if e != nil {
		return false, e
	}

	return result.HasNext(), nil
}

//func ConvertInt64ListAsArgument(list []int64, args *pqlib.Args) string {
//	return stringutil.JoinInt64WithConvertMethod(list, ",", func(elem int64) string {
//		return args.Next(elem)
//	})
//}

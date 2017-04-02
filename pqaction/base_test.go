package pqaction_test

import "github.com/mleuth/pqlib/pqutil/pqreflect"

func tableName(entity interface{}) string {
	structInfo := pqreflect.NewStructInfo(entity)
	return structInfo.Name()
}

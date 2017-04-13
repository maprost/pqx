package pqx_test

import "github.com/maprost/pqx/pqutil/pqreflect"

func tableName(entity interface{}) string {
	structInfo := pqreflect.NewStructInfo(entity)
	return structInfo.Name()
}

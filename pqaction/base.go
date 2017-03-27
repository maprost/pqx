package pqaction

import "github.com/mleuth/pqlib/pqutil/pqreflect"

func TableName(entity interface{}) string {
	structInfo := pqreflect.NewStructInfo(entity)
	return structInfo.Name()
}

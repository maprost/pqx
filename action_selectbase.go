package pqx

import (
	"github.com/maprost/pqx/pqutil"
	"github.com/maprost/pqx/pqutil/pqreflect"
)

func SelectList(entity interface{}) string {
	structInfo := pqreflect.NewStructInfo(entity)
	return selectList(structInfo, "")
}

func SelectListWithAlias(entity interface{}, alias string) string {
	structInfo := pqreflect.NewStructInfo(entity)
	return selectList(structInfo, alias)
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

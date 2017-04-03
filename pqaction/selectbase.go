package pqaction

import (
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

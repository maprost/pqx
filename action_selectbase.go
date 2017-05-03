package pqx

import (
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqutil"
)

func SelectList(entity interface{}) string {
	table, err := pqtable.New(entity)
	if err != nil {
		panic(err)
	}
	return selectList(table, "")
}

func SelectListWithAlias(entity interface{}, alias string) string {
	table, err := pqtable.New(entity)
	if err != nil {
		panic(err)
	}
	return selectList(table, alias)
}

func selectList(table *pqtable.Table, alias string) string {
	list := ""
	if alias != "" {
		alias += "."
	}

	for _, column := range table.Columns() {
		list = pqutil.Concate(list, alias+column.Name(), ",")
	}

	return list
}

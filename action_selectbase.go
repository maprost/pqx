package pqx

import (
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqutil"
)

func SelectRowList(entity interface{}) string {
	table, err := pqtable.New(entity)
	if err != nil {
		panic(err)
	}
	return selectRowList(table, "")
}

func SelectRowListWithAlias(entity interface{}, alias string) string {
	table, err := pqtable.New(entity)
	if err != nil {
		panic(err)
	}
	return selectRowList(table, alias)
}

func selectRowList(table *pqtable.Table, alias string) string {
	list := ""
	if alias != "" {
		alias += "."
	}

	for _, column := range table.Columns() {
		list = pqutil.Concate(list, alias+column.Name(), ",")
	}

	return list
}

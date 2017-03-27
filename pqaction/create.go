package pqaction

import (
	"github.com/matthiasleuthaeuser/pqlib"
	"github.com/matthiasleuthaeuser/pqlib/pqutil"
	"github.com/matthiasleuthaeuser/pqlib/pqutil/pqreflect"
)

/*
	type Example struct {	<- table name
		id int <- always PK
		att1 string `sql: UK` <- UK att is unique key
		att2 float `sql: UK[1]` <- UK group
		att2 string `sql: UK[1]`
		att3 int `sql: FK[table]` <- foreign key to table
 	}

 	entities contains only primitive types
*/

// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func Create(db pqlib.Transaction, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	lines := ""
	for _, field := range structInfo.Fields() {
		dbType, err := dbType(field)
		if err != nil {
			return err
		}

		line := "\t" + field.Name() + " " + dbType
		if isPrimaryKey(field) {
			line += " PRIMARY KEY"
		}

		lines = pqutil.Concate(lines, line, ",\n")
	}

	// TODO: insert unique and foreign keys
	sql := "CREATE TABLE " + structInfo.Name() + "(\n" + lines + "\n)"
	_, e := db.Query(sql, pqlib.NewArgs())
	return e
}

package pqx

import (
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqutil"
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

// Create an entity via pqx.LogQuery and use a default logger for logging.
// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func Create(entity interface{}) error {
	return LogCreate(entity, pqutil.DefaultLogger)
}

// LogCreate create an entity via pqx.LogQuery and use the given pqdep.Logger for logging.
// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func LogCreate(entity interface{}, logger pqdep.Logger) error {
	return createFunc(queryFuncWrapper(logger), entity)
}

// Create an entity via tx.Query and use the given tx.log for logging.
// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func (tx *Transaction) Create(entity interface{}) error {
	return createFunc(tx.Query, entity)
}

// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func createFunc(qFunc queryFunc, entity interface{}) error {
	table, err := pqtable.New(entity)
	if err != nil {
		return err
	}

	lines := ""
	for _, column := range table.Columns() {
		line := "\t" + column.Name() + " " + column.Type().String()
		if column.PrimaryKeyTag() {
			line += " PRIMARY KEY"
		}
		if column.Nullable() {
			line += " NULL"
		} else {
			line += " NOT NULL"
		}

		lines = pqutil.Concate(lines, line, ",\n")
	}

	// TODO: insert unique and foreign keys
	sql := "CREATE TABLE " + table.Name() + "(\n" + lines + "\n)"
	rows, err := qFunc(sql, pqarg.New())
	defer closeRows(rows)
	return err
}

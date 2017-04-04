package pqaction

import (
	"database/sql"
	"errors"
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqdep"
	"github.com/mleuth/pqlib/pqutil"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
	"reflect"
	"time"
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

// Create an entity via pqlib.Query method and use a default logger for logging.
// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func Create(entity interface{}) error {
	return CreateLg(entity, pqutil.DefaultLogger)
}

// CreateLg create an entity via pqlib.Query method and use a pqdep.Logger for logging.
// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func CreateLg(entity interface{}, logger pqdep.Logger) error {
	return createFunc(queryFuncWrapper(logger), entity)
}

// CreateTx create an entity over an active transaction.
// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func CreateTx(tx pqlib.Transaction, entity interface{}) error {
	return createFunc(tx.Query, entity)
}

// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func createFunc(qFunc queryFunc, entity interface{}) error {
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
	_, e := qFunc(
		"CREATE TABLE "+structInfo.Name()+"(\n"+lines+"\n)",
		pqlib.NewArgs())
	return e
}

func dbType(field pqreflect.Field) (dbType string, e error) {
	switch field.Kind() {
	case reflect.String:
		dbType = "text NOT NULL"

	case reflect.Int8, reflect.Int16, reflect.Uint8:
		if isAutoIncrement(field) {
			dbType = "smallserial"
		} else {
			dbType = "smallint NOT NULL"
		}
	case reflect.Int, reflect.Int32, reflect.Uint16:
		if isAutoIncrement(field) {
			dbType = "serial"
		} else {
			dbType = "integer NOT NULL"
		}
	case reflect.Int64, reflect.Uint32:
		if isAutoIncrement(field) {
			dbType = "bigserial"
		} else {
			dbType = "bigint NOT NULL"
		}
	case reflect.Uint64:
		dbType = "numeric NOT NULL"

	case reflect.Bool:
		dbType = "bool NOT NULL"

	case reflect.Float32:
		dbType = "real NOT NULL"
	case reflect.Float64:
		dbType = "double precision NOT NULL"

	default:
		// some important struct data
		switch field.TypeInterface().(type) {
		case time.Time:
			dbType = "timestamp with time zone NOT NULL"

		case sql.NullBool:
			dbType = "bool NULL"

		case sql.NullInt64:
			dbType = "bigint NULL"

		case sql.NullString:
			dbType = "text NULL"

		case sql.NullFloat64:
			dbType = "double precision NULL"

		default:
			e = errors.New("Not supported field type: " + field.Name() + " (" + field.Type() + ").")
		}
	}
	return
}

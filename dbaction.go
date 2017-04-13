package pqx

import (
	"database/sql"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqutil"
)

// Create an entity via pqdb.LogQuery and use a default logger for logging.
// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func Create(entity interface{}) error {
	return LogCreate(entity, pqutil.DefaultLogger)
}

// Create an entity via pqdb.LogQuery and use the given pqdep.Logger for logging.
// CREATE (
// 		id TYPE PRIMARY KEY,
// 		att1 TYPE,
// 		att2 TYPE
//		Unique(att1, att2)
// )
func LogCreate(entity interface{}, logger pqdep.Logger) error {
	return pqaction.Create(
		func(sql string, args pqarg.Args) (*sql.Rows, error) {
			return LogQuery(logger, sql, args)
		}, entity)
}

// Insert an entity via pqdb.LogQueryRow and use a default logger for logging.
// INSERT INTO table_name (AI, column1,column2,column3,...)
// VALUES (DEFAULT, value1,value2,value3,...) RETURNING AI;
func Insert(entity interface{}) error {
	return LogInsert(entity, pqutil.DefaultLogger)
}

// LogInsert insert an entity via pqdb.LogQueryRow and use the given pqdep.Logger for logging.
// INSERT INTO table_name (AI, column1,column2,column3,...)
// VALUES (DEFAULT, value1,value2,value3,...) RETURNING AI;
func LogInsert(entity interface{}, logger pqdep.Logger) error {
	return pqaction.Insert(
		func(sql string, args pqarg.Args) *sql.Row {
			return LogQueryRow(logger, sql, args)
		}, entity)
}

// Update an entity via pqdb.LogQuery and use a default logger for logging.
//UPDATE table_name
//SET column1 = value1,
//column2 = value2,
//...
//WHERE PK = valueX (with PK tag)
func Update(entity interface{}) error {
	return LogUpdate(entity, pqutil.DefaultLogger)
}

// LogUpdate update an entity via pqdb.LogQuery and use the given pqdep.Logger for logging.
//UPDATE table_name
//SET column1 = value1,
//column2 = value2,
//...
//WHERE PK = valueX (with PK tag)
func LogUpdate(entity interface{}, logger pqdep.Logger) error {
	return pqaction.Update(
		func(sql string, args pqarg.Args) (*sql.Rows, error) {
			return LogQuery(logger, sql, args)
		}, entity)
}

// Delete an entity via pqdb.LogQuery and use a default logger for logging.
// DELETE FROM table_name
// WHERE PK = value;
func Delete(entity interface{}) error {
	return LogDelete(entity, pqutil.DefaultLogger)
}

// LogDelete delete an entity via pqdb.LogQuery and use the given pqdep.Logger for logging.
// DELETE FROM table_name
// WHERE PK = value;
func LogDelete(entity interface{}, logger pqdep.Logger) error {
	return pqaction.Delete(
		func(sql string, args pqarg.Args) (*sql.Rows, error) {
			return LogQuery(logger, sql, args)
		}, entity)
}

package pqx

import (
	"database/sql"
	"github.com/maprost/pqx/pqarg"
)

// Insert an entity via pqtx.tx.QueryRow and use the given tx.log for logging.
// INSERT INTO table_name (AI, column1,column2,column3,...)
// VALUES (DEFAULT, value1,value2,value3,...) RETURNING AI;
func (tx *Transaction) Insert(entity interface{}) error {
	return pqaction.Insert(
		func(sql string, args pqarg.Args) *sql.Row {
			return tx.QueryRow(sql, args)
		}, entity)
}

// Update an entity via pqtx.tx.Query and use the tx.log for logging.
// UPDATE table_name
// SET column1 = value1,
// column2 = value2,
// ...
// WHERE PK = valueX (with PK tag)
func (tx *Transaction) Update(entity interface{}) error {
	return pqaction.Update(
		func(sql string, args pqarg.Args) (*sql.Rows, error) {
			return tx.Query(sql, args)
		}, entity)
}

// Delete an entity via pqtx.tx.Query and use the tx.log for logging.
// DELETE FROM table_name
// WHERE PK = value;
func (tx *Transaction) Delete(entity interface{}) error {
	return pqaction.Delete(
		func(sql string, args pqarg.Args) (*sql.Rows, error) {
			return tx.Query(sql, args)
		}, entity)
}

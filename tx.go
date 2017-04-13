package pqx

import (
	"database/sql"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqutil"
)

type Transaction struct {
	log pqdep.Logger
	tx  *sql.Tx
}

func New() (Transaction, error) {
	tx, err := pqdb.DB.Begin()

	return Transaction{
		log: pqutil.DefaultLogger,
		tx:  tx,
	}, err
}

func (tx *Transaction) AddLogger(logger pqdep.Logger) {
	tx.log = logger
}

func (tx *Transaction) Query(sql string, args pqarg.Args) (rows *sql.Rows, err error) {

	pqlib.LogQueryFunc(func(sql string, args ...interface{}) {
		rows, err = tx.Query(sql, args...)
	}, tx.log, sql, args)

	return
}

func (tx *Transaction) QueryRow(sql string, args pqarg.Args) (row *sql.Row) {

	pqlib.LogQueryFunc(func(sql string, args ...interface{}) {
		row = tx.QueryRow(sql, args...)
	}, tx.log, sql, args)

	return
}

func (tx *Transaction) Commit() error {
	if tx.tx == nil {
		return nil
	}

	err := tx.tx.Commit()
	if err != nil {
		tx.log.Printf("Fail to commit: %s", err.Error())
		return err
	}

	tx.tx = nil
	return nil
}

func (tx *Transaction) Rollback() error {
	if tx.tx == nil {
		return nil
	}

	err := tx.tx.Rollback()
	if err != nil {
		tx.log.Printf("Fail to rollback: %s", err.Error())
		return err
	}

	tx.tx = nil
	return nil
}

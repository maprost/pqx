package pqlib

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mleuth/pqlib/pqdep"
	"github.com/mleuth/pqlib/pqutil"
)

type Transaction interface {
	Query(sql string, args Args) (Result, error)
	Commit() error
	Rollback() error
}

type transaction struct {
	log      pqdep.Logger
	tx       *sql.Tx
	lastRows *Result
}

func NewTransaction() Transaction {
	return NewTransactionLg(pqutil.DefaultLogger)
}

func NewTransactionLg(logger pqdep.Logger) Transaction {
	return &transaction{log: logger, tx: nil, lastRows: nil}
}

func (pq *transaction) Query(sql string, args Args) (Result, error) {
	e := pq.begin()
	if e != nil {
		return Result{}, e
	}

	// first close the last rows of the last query
	pq.closeLastRow()

	// execute query
	result, e := queryFunc(pq.tx.Query, pq.log, sql, args)
	if e != nil {
		return Result{}, e
	}

	pq.lastRows = &result
	return result, nil
}

func (pq *transaction) Commit() error {
	if pq.tx == nil {
		return nil
	}

	pq.closeLastRow()
	e := pq.tx.Commit()
	if e != nil {
		return e
	}

	pq.tx = nil
	return nil
}

func (pq *transaction) Rollback() error {
	if pq.tx == nil {
		return nil
	}

	pq.closeLastRow()
	e := pq.tx.Rollback()
	if e != nil {
		return e
	}

	pq.tx = nil
	return nil
}

func (pq *transaction) begin() error {
	if pq.tx == nil {
		tx, e := db.Begin()
		if e != nil {
			return e
		}

		pq.tx = tx
	}
	return nil
}

func (pq *transaction) closeLastRow() error {
	if pq.lastRows != nil {
		return pq.lastRows.close()
	}
	return nil
}

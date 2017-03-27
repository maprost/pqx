package pqlib

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mleuth/pqlib/pqdep"
	"github.com/mleuth/timeutil"
)

var db *sql.DB = nil

func OpenDatabaseConnection(info pqdep.ConnectionInfo) error {
	var e error

	// close old db connection before open a new one
	if db != nil {
		e = db.Close()
		if e != nil {
			return e
		}
	}

	db, e = sql.Open(info.GetDatabaseDriver(),
		"user="+info.GetUserName()+
			" dbname="+info.GetDataBase()+
			" sslmode=disable"+
			" host="+info.GetHost()+
			" port="+info.GetPort())
	return e
}

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

func New(logger pqdep.Logger) Transaction {
	return &transaction{log: logger, tx: nil, lastRows: nil}
}

func (pq *transaction) Query(sql string, args Args) (Result, error) {
	e := pq.begin()
	if e != nil {
		return Result{}, e
	}

	// first close the last rows of the last query
	pq.closeLastRow()

	// track duration
	stopwatch := timeutil.NewStopwatch()
	// execute
	rows, e := pq.tx.Query(sql, args.get()...)
	// log sql + duration
	stopwatch.Stop()
	pq.log.Printf("[time: "+stopwatch.String()+"] SQL: "+sql, args.get()...)
	// check error
	if e != nil {
		return Result{}, e
	}

	pq.lastRows = &Result{rows: rows, hasNext: false}
	return *pq.lastRows, nil
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

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

	db, e = sql.Open(info.DatabaseDriver(),
		"user="+info.UserName()+
			" dbname="+info.DataBase()+
			" sslmode=disable"+
			" host="+info.Host()+
			" port="+info.Port())
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

func Query(logger pqdep.Logger, sql string, args Args) (Result, error) {
	rows, e := query(db.Query, logger, sql, args)
	// check error
	if e != nil {
		return Result{}, e
	}

	return Result{rows: rows, hasNext: false}, nil
}

func (pq *transaction) Query(sql string, args Args) (Result, error) {
	e := pq.begin()
	if e != nil {
		return Result{}, e
	}

	// first close the last rows of the last query
	pq.closeLastRow()

	// execute query
	rows, e := query(pq.tx.Query, pq.log, sql, args)
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

func query(queryFunc func(query string, args ...interface{}) (*sql.Rows, error), logger pqdep.Logger, sql string, args Args) (*sql.Rows, error) {
	// track duration
	stopwatch := timeutil.NewStopwatch()
	// execute
	rows, e := queryFunc(sql, args.get()...)
	// log sql + duration
	stopwatch.Stop()
	logger.Printf("[time: "+stopwatch.String()+"] SQL: "+sql, args.get()...)

	return rows, e
}

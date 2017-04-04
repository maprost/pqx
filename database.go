package pqlib

import (
	"database/sql"
	"github.com/mleuth/pqlib/pqdep"
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

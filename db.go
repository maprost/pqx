package pqx

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/maprost/pqx/pqdep"
)

var DB *sql.DB = nil

func OpenDatabaseConnection(info pqdep.ConnectionInfo) error {
	var e error

	// close old db connection before open a new one
	if DB != nil {
		e = DB.Close()
		if e != nil {
			return e
		}
	}

	DB, e = sql.Open(info.DatabaseDriver(),
		"user="+info.UserName()+
			" dbname="+info.DataBase()+
			" sslmode=disable"+
			" host="+info.Host()+
			" port="+info.Port())
	return e
}

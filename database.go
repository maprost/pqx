package pqx

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqdep"
	"github.com/maprost/pqx/pqutil"
)

var DB *sql.DB = nil

func OpenDatabaseConnection(info pqdep.ConnectionInfo) error {
	var err error

	// close old db connection before open a new one
	if DB != nil {
		err = DB.Close()
		if err != nil {
			return err
		}
	}

	DB, err = sql.Open(info.DatabaseDriver(),
		"user="+info.UserName()+
			" dbname="+info.DataBase()+
			" sslmode=disable"+
			" host="+info.Host()+
			" port="+info.Port())
	return err
}

func Query(sql string, args pqarg.Args) (*sql.Rows, error) {
	return LogQuery(sql, args, pqutil.DefaultLogger)
}

func LogQuery(sql string, args pqarg.Args, logger pqdep.Logger) (rows *sql.Rows, err error) {
	logWrapper(
		func(sql string, args ...interface{}) {
			rows, err = DB.Query(sql, args...)
		}, sql, args, logger)

	return
}

func QueryRow(sql string, args pqarg.Args) *sql.Row {
	return LogQueryRow(sql, args, pqutil.DefaultLogger)
}

func LogQueryRow(sql string, args pqarg.Args, logger pqdep.Logger) (row *sql.Row) {
	logWrapper(
		func(sql string, args ...interface{}) {
			row = DB.QueryRow(sql, args...)
		}, sql, args, logger)

	return
}

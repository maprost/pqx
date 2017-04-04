package pqtest

import (
	"github.com/mleuth/assertion"
	"github.com/mleuth/pqlib"
	"github.com/mleuth/timeutil"
	"log"
	"os"
)

type DataInfo struct{}

func (d DataInfo) DatabaseDriver() string {
	return "postgres"
}

func (d DataInfo) DataBase() string {
	return "testDB"
}

func (d DataInfo) Host() string {
	return "localhost"
}

func (d DataInfo) Port() string {
	return "5441"
}

func (d DataInfo) UserName() string {
	return "postgres"
}

func InitTransactionTest(t assertion.TestEnvironment) (pqlib.Transaction, assertion.Assert) {
	timeutil.Reset()

	assert := assertion.New(t)
	e := pqlib.OpenDatabaseConnection(DataInfo{})
	assert.Nil(e)

	tx := pqlib.NewTransaction(log.New(os.Stdout, "", 0))

	return tx, assert
}

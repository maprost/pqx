package pqtest

import (
	"github.com/maprost/assertion"
	"github.com/maprost/pqx"
	"github.com/maprost/timeutil"
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

func InitDatabaseTest(t assertion.TestEnvironment) assertion.Assert {
	timeutil.Reset()
	assert := assertion.New(t)

	err := pqx.OpenDatabaseConnection(DataInfo{})
	assert.Nil(err)

	return assert
}

func InitTransactionTest(t assertion.TestEnvironment) (pqx.Transaction, assertion.Assert) {
	assert := InitDatabaseTest(t)

	tx, err := pqx.New()
	assert.Nil(err)
	tx.AddLogger(log.New(os.Stdout, "", 0))

	return tx, assert
}

package pqtest

import (
	"log"
	"os"

	"github.com/maprost/assertion"
	"github.com/maprost/pqx"
	"github.com/maprost/timeutil"
)

type DataInfo struct{}

func (d DataInfo) DatabaseDriver() string {
	return "postgres"
}

func (d DataInfo) DataBase() string {
	return "test_pqx"
}

func (d DataInfo) Host() string {
	return "localhost"
}

func (d DataInfo) Port() string {
	return "5432"
}

func (d DataInfo) UserName() string {
	return "postgres"
}

func InitDatabase() error {
	return pqx.OpenDatabaseConnection(DataInfo{})
}

func InitDatabaseTest(t assertion.TestEnvironment) assertion.Assert {
	timeutil.Reset()
	assert := assertion.New(t)

	err := InitDatabase()
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

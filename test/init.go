package test

import (
	"github.com/mleuth/pqlib"
	"github.com/mleuth/timeutil"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

type DataInfo struct{}

func (d DataInfo) GetDatabaseDriver() string {
	return "postgres"
}

func (d DataInfo) GetDataBase() string {
	return "testDB"
}

func (d DataInfo) GetHost() string {
	return "localhost"
}

func (d DataInfo) GetPort() string {
	return "5432"
}

func (d DataInfo) GetUserName() string {
	return "postgres"
}

func InitTransactionTest(t *testing.T) (pqlib.Transaction, require.Assertions) {
	timeutil.Reset()

	check := require.New(t)
	e := pqlib.OpenDatabaseConnection(DataInfo{})
	check.Nil(e)

	tx := pqlib.New(log.New(os.Stdout, "", 0))

	return tx, check
}

package pqx_test

import (
	"testing"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqtest"
	"github.com/maprost/pqx/pqutil"
)

func TestTxCreate(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestTxCreateStruct struct {
		Id  int64
		Msg string
	}
	err := tx.Create(TestTxCreateStruct{})
	assert.Nil(err)
}

func TestDBCreate(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type TestDBCreateStruct struct {
		Id  int64
		Msg string
	}
	err := pqx.Create(TestDBCreateStruct{})
	assert.Nil(err)
}

func TestDBLogCreate(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type TestDBLogCreateStruct struct {
		Id  int64
		Msg string
	}
	err := pqx.LogCreate(TestDBLogCreateStruct{}, pqutil.DefaultLogger)
	assert.Nil(err)
}

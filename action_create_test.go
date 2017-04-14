package pqx_test

import (
	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqtest"
	"github.com/maprost/pqx/pqutil"
	"testing"
)

func TestTxCreate(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestTxCreateStruct struct {
		Id  int64
		Msg string
	}
	var testStruct TestTxCreateStruct
	err := tx.Create(&testStruct)
	assert.Nil(err)
}

func TestDBCreate(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type TestDBCreateStruct struct {
		Id  int64
		Msg string
	}
	var testStruct TestDBCreateStruct
	err := pqx.Create(&testStruct)
	assert.Nil(err)
}

func TestDBLogCreate(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type TestDBLogCreateStruct struct {
		Id  int64
		Msg string
	}
	var testStruct TestDBLogCreateStruct
	err := pqx.LogCreate(&testStruct, pqutil.DefaultLogger)
	assert.Nil(err)
}

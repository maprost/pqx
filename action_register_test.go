package pqx_test

import (
	"testing"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqnull"
	"github.com/maprost/pqx/pqtest"
)

func TestTableExists(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type TestTableExistsStruct struct {
		Id  int64
		Msg pqnull.String
	}

	err := pqx.Register(TestTableExistsStruct{})
	assert.Nil(err)

	err = pqx.Register(TestTableExistsStruct{})
	assert.Nil(err)
}

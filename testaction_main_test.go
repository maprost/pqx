package pqx_test

import (
	"os"
	"testing"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqnull"
	"github.com/maprost/pqx/pqtest"
)

type ThreeElementsStruct struct {
	Id     int64 `pqx:"PK AI"`
	Msg    string
	UserId pqnull.Int64
}

type noLogging struct {
}

func (b noLogging) Printf(format string, v ...interface{}) {
	// do nothing
}

func TestMain(m *testing.M) {
	pqtest.InitDatabase()
	registerBenchTables()
	retCode := m.Run()
	os.Exit(retCode)
}

func registerBenchTables() {
	pqx.Register(
		ThreeElementsStruct{},
	)
}

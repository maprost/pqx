package pqlib_test

import (
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqtest"
	"testing"
)

func TestScanStruct_simpleSelect(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestStruct struct {
		String string
		Int    int8
		Bool   bool
		Double float32
	}

	result, err := tx.Query("Select 'hello', 42, true, 12.4", pqlib.NewArgs())
	assert.Nil(err)

	counter := 0
	for result.HasNext() {
		counter++
		var toFill TestStruct
		err = result.ScanStruct(&toFill)
		assert.Nil(err)

		assert.Equal(toFill.String, "hello")
		assert.Equal(toFill.Int, int8(42))
		assert.Equal(toFill.Bool, true)
		assert.Equal(toFill.Double, float32(12.4))
	}
	assert.Equal(counter, 1)

	err = tx.Commit()
	assert.Nil(err)
}

func TestScanStruct_notSupportedField(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestStruct struct {
		Uintptr uintptr
	}

	result, err := tx.Query("Select 42", pqlib.NewArgs())
	assert.Nil(err)

	for result.HasNext() {
		var toFill TestStruct
		err = result.ScanStruct(&toFill)
		assert.NotNil(err)
	}
}

func TestScanStruct_wrongType(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestStruct struct {
		Int int64
	}

	result, err := tx.Query("Select 'hello'", pqlib.NewArgs())
	assert.Nil(err)

	for result.HasNext() {
		var toFill TestStruct
		err = result.ScanStruct(&toFill)
		assert.NotNil(err)
	}
}

func TestScanStruct_selectWith2RowsAsResult(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestStruct struct {
		Id  int
		Msg string
	}

	result, err := tx.Query("(Select 1, 'hello') UNION (SELECT 2, 'world')", pqlib.NewArgs())
	assert.Nil(err)

	assert.True(result.HasNext())
	var firstRow TestStruct
	err = result.ScanStruct(&firstRow)
	assert.Nil(err)
	assert.Equal(firstRow.Id, 1)
	assert.Equal(firstRow.Msg, "hello")

	assert.True(result.HasNext())
	var secondRow TestStruct
	err = result.ScanStruct(&secondRow)
	assert.Nil(err)
	assert.Equal(secondRow.Id, 2)
	assert.Equal(secondRow.Msg, "world")

	assert.False(result.HasNext())
}

func TestScanStruct_multiplyResults(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestStruct struct {
		Id  int
		Msg string
	}

	result, err := tx.Query("(Select 1, 'hello') UNION (SELECT 2, 'world')", pqlib.NewArgs())
	assert.Nil(err)

	resultList := []TestStruct{}
	var toFill TestStruct
	for result.HasNext() {
		err := result.ScanStruct(&toFill)
		assert.Nil(err)

		resultList = append(resultList, toFill)
	}

	assert.Equal(len(resultList), 2)

	assert.Equal(resultList[0].Id, 1)
	assert.Equal(resultList[0].Msg, "hello")

	assert.Equal(resultList[1].Id, 2)
	assert.Equal(resultList[1].Msg, "world")
}

func TestQuery_queryArguments(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestStruct struct {
		Msg string
	}

	args := pqlib.NewArgs()
	result, err := tx.Query("Select "+args.Next("hello")+"::text", args)
	assert.Nil(err)

	resultList := []TestStruct{}
	for result.HasNext() {
		var toFill TestStruct
		err := result.ScanStruct(&toFill)
		assert.Nil(err)

		resultList = append(resultList, toFill)
	}

	assert.Equal(len(resultList), 1)
	assert.Equal(resultList[0].Msg, "hello")
}

func TestQuery_injection(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestStruct struct {
		Msg string
	}

	args := pqlib.NewArgs()
	result, err := tx.Query("Select "+args.Next("blob'::text UNION SELECT 'badass")+"::text", args)
	assert.Nil(err)

	resultList := []TestStruct{}
	for result.HasNext() {
		var toFill TestStruct
		err := result.ScanStruct(&toFill)
		assert.Nil(err)

		resultList = append(resultList, toFill)
	}

	assert.Equal(len(resultList), 1)
	assert.Equal(resultList[0].Msg, "blob'::text UNION SELECT 'badass")
}

func TestScan_simpleSelect(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	result, err := tx.Query("Select 'hello', 42, true, 12.4", pqlib.NewArgs())
	assert.Nil(err)

	counter := 0
	for result.HasNext() {
		counter++
		var String string
		var Int int8
		var Bool bool
		var Double float32
		err = result.Scan(&String, &Int, &Bool, &Double)
		assert.Nil(err)

		assert.Equal(String, "hello")
		assert.Equal(Int, int8(42))
		assert.Equal(Bool, true)
		assert.Equal(Double, float32(12.4))
	}
	assert.Equal(counter, 1)
}

func TestScan_forgetColumn(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	result, err := tx.Query("Select 'hello', 42, true, 12.4", pqlib.NewArgs())
	assert.Nil(err)

	counter := 0
	for result.HasNext() {
		counter++
		var String string = "wrong"
		var Int int8 = -1
		var Bool bool = false
		err = result.Scan(&String, &Int, &Bool)
		assert.NotNil(err)

		assert.Equal(String, "wrong")
		assert.Equal(Int, int8(-1))
		assert.Equal(Bool, false)
	}
	assert.Equal(counter, 1)
}

func TestRollback_withNotOpenConnection(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	err := tx.Rollback()
	assert.Nil(err)

}

func TestCommit_withNotOpenConnection(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	err := tx.Commit()
	assert.Nil(err)
}

func TestRollback_withEmptyContent(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	result, err := tx.Query("Select 1;", pqlib.NewArgs())
	assert.Nil(err)

	var Int int8
	err = result.Scan(&Int)
	assert.Nil(err)
	assert.Equal(Int, int8(1))

	err = tx.Rollback()
	assert.Nil(err)
}

func TestCommit_withEmptyContent(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	result, err := tx.Query("Select 1;", pqlib.NewArgs())
	assert.Nil(err)
	assert.NotNil(result)

	err = tx.Commit()
	assert.Nil(err)
}

func TestMultiplyHasNextCommand(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	result, err := tx.Query("Select 1;", pqlib.NewArgs())
	assert.Nil(err)

	result.HasNext()
	result.HasNext()

	var Int int8
	err = result.Scan(&Int)
	assert.Nil(err)

	assert.Equal(Int, int8(1))
}

func TestScanStruct_moreArgsInQuery(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestStruct struct {
		String string
		Int    int8
		Bool   bool
		Double float32
	}

	args := pqlib.NewArgs()
	result, err := tx.Query("Select "+
		args.Next("hello")+"::text, "+
		args.Next(42)+"::int, "+
		args.Next(true)+"::bool, "+
		args.Next(12.4)+"::float", args)
	assert.Nil(err)

	var toFill TestStruct
	err = result.ScanStruct(&toFill)
	assert.Nil(err)

	assert.Equal(toFill.String, "hello")
	assert.Equal(toFill.Int, int8(42))
	assert.Equal(toFill.Bool, true)
	assert.Equal(toFill.Double, float32(12.4))

	err = tx.Rollback()
	assert.Nil(err)
}

func TestScanStruct_someQueries(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestStruct struct {
		String string
		Int    int8
	}

	for i := 0; i < 4; i++ {
		args := pqlib.NewArgs()
		result, err := tx.Query("Select "+args.Next("hello")+"::text, "+args.Next(i)+"::int", args)
		assert.Nil(err)

		var toFill TestStruct
		err = result.ScanStruct(&toFill)
		assert.Nil(err)

		assert.Equal(toFill.String, "hello")
		assert.Equal(toFill.Int, int8(i))
	}
	err := tx.Rollback()
	assert.Nil(err)
}

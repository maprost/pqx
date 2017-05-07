package pqx_test

import (
	"strconv"
	"testing"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqtest"
)

func TestMultiResultsToScan(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type TestMultiResultsToScanStruct struct {
		Id  int64 `pqx:"PK AI"`
		Msg string
	}
	pqx.Register(TestMultiResultsToScanStruct{})

	list := []TestMultiResultsToScanStruct{}
	for i := 1; i <= 5; i++ {
		// insert entity
		entity := TestMultiResultsToScanStruct{Msg: "hello_" + strconv.Itoa(i)}
		err := pqx.Insert(&entity)
		assert.Nil(err)
		assert.Equal(entity.Id, int64(i))

		list = append(list, entity)
	}

	// select entity -> 1, "hello", 42
	{
		var prototype TestMultiResultsToScanStruct
		selectedList := []TestMultiResultsToScanStruct{}
		err := pqx.SelectList(&prototype, func() {
			selectedList = append(selectedList, prototype)
		})
		assert.Nil(err)
		assert.Len(selectedList, 5)
		for i, e := range selectedList {
			assert.Equal(e.Id, int64(i+1))
			assert.Equal(e.Msg, "hello_"+strconv.Itoa(i+1))
		}
	}

	// select entity -> 1, "hello", 42
	{
		result, err := pqx.Query("SELECT "+pqx.SelectRowList(TestMultiResultsToScanStruct{})+
			" FROM "+pqtable.TableName(TestMultiResultsToScanStruct{}), pqarg.New())
		assert.Nil(err)

		var prototype TestMultiResultsToScanStruct
		selectedList := []TestMultiResultsToScanStruct{}
		for result.Next() {
			err := pqx.ScanStruct(result, &prototype)
			assert.Nil(err)
			selectedList = append(selectedList, prototype)
		}

		assert.Len(selectedList, 5)
		for i, e := range selectedList {
			assert.Equal(e.Id, int64(i+1))
			assert.Equal(e.Msg, "hello_"+strconv.Itoa(i+1))
		}
	}

	for _, e := range list {
		// delete entity
		err := pqx.Delete(&e)
		assert.Nil(err)
	}

}

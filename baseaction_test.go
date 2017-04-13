package pqx_test

import (
	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqarg"
	"github.com/maprost/pqx/pqtest"
	"github.com/maprost/timeutil"
	"strconv"
	"testing"
	"time"
)

func TestCreateSelect_simple(t *testing.T) {
	tx, _ := pqtest.InitTransactionTest(t)

	type TestCreateSelectStruct struct {
		Id  int64
		Msg string
	}
	var testStruct TestCreateSelectStruct
	pqx.Register(tx, &testStruct)
}

func TestSimpleWorkflow(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestSimpleWorkflowStruct struct {
		Id     int64 `sql:"PK AI"`
		Msg    string
		UserId int64
	}
	var testStruct TestSimpleWorkflowStruct
	pqlib.Register(tx, &testStruct)

	// insert entity
	entity := TestSimpleWorkflowStruct{Msg: "hello", UserId: 42}
	e := tx.Insert(&entity)
	assert.Nil(e)
	assert.Equal(entity.Id, int64(1))

	// select entity -> 1, "hello", 42
	{
		args := pqarg.New()
		result, e := tx.Query("SELECT "+pqlib.SelectList(&entity)+" FROM "+
			tableName(&testStruct)+" WHERE id = "+args.Next(1), args)
		assert.Nil(e)

		var checkSelect TestSimpleWorkflowStruct
		e = pqlib.ScanStruct(result, &checkSelect)
		assert.Nil(e)
		assert.Equal(checkSelect.Id, int64(1))
		assert.Equal(checkSelect.Msg, "hello")
		assert.Equal(checkSelect.UserId, int64(42))
	}

	// update entity
	entity.Msg = "world"
	e = tx.Update(&entity)
	assert.Nil(e)

	// select entity -> 1, "world", 42
	{
		args := pqarg.New()
		result, e := tx.Query("SELECT "+pqlib.SelectList(&entity)+" FROM "+
			tableName(&testStruct)+" WHERE id = "+args.Next(1), args)
		assert.Nil(e)

		var checkSelect TestSimpleWorkflowStruct
		e = pqlib.ScanStruct(result, &checkSelect)
		assert.Nil(e)
		assert.Equal(checkSelect.Id, int64(1))
		assert.Equal(checkSelect.Msg, "world")
		assert.Equal(checkSelect.UserId, int64(42))
	}

	// delete entity
	e = tx.Delete(&entity)
	assert.Nil(e)

	// select entity -> nothing found
	{
		args := pqarg.New()
		result, e := tx.Query("SELECT "+pqx.SelectList(&entity)+" FROM "+
			tableName(&testStruct)+" WHERE id = "+args.Next(1), args)
		assert.Nil(e)
		assert.False(result.Next())
	}

	e = tx.Commit()
	assert.Nil(e)
}

func TestUpdateWithoutID(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)

	type TestUpdateWithoutIDStruct struct {
		UserID int64
		Msg    string
	}
	var testStruct TestUpdateWithoutIDStruct
	pqx.Register(tx, &testStruct)

	// insert entity
	entity := TestUpdateWithoutIDStruct{Msg: "hello", UserID: 42}
	e := tx.Insert(&entity)
	assert.Nil(e)

	// select entity -> 42, "hello"
	{
		result, e := tx.Query("SELECT "+pqx.SelectList(&entity)+" FROM "+
			tableName(&testStruct), pqarg.New())
		assert.Nil(e)

		var checkSelect TestUpdateWithoutIDStruct
		e = pqx.ScanStruct(result, &checkSelect)
		assert.Nil(e)
		assert.Equal(checkSelect.UserID, int64(42))
		assert.Equal(checkSelect.Msg, "hello")
	}

	// try to update entity -> eor
	entity.Msg = "world"
	e = tx.Update(&entity)
	assert.NotNil(e)
	assert.Equal(e.Error(), "No primary key available.")

	// select entity -> 42, "hello"
	{
		result, e := tx.Query("SELECT "+pqaction.SelectList(&entity)+
			" FROM "+tableName(&testStruct), pqarg.New())
		assert.Nil(e)

		var checkSelect TestUpdateWithoutIDStruct
		e = pqres.ScanStruct(result, &checkSelect)
		assert.Nil(e)
		assert.Equal(checkSelect.UserID, int64(42))
		assert.Equal(checkSelect.Msg, "hello")
	}

	// try to delete entity -> eor
	e = tx.Delete(&entity)
	assert.NotNil(e)
	assert.Equal(e.Error(), "No primary key available.")

	// select entity -> 42, "hello"
	{
		result, e := tx.Query("SELECT "+pqaction.SelectList(&entity)+
			" FROM "+tableName(&testStruct), pqarg.New())
		assert.Nil(e)

		var checkSelect TestUpdateWithoutIDStruct
		e = pqres.ScanStruct(result, &checkSelect)
		assert.Nil(e)
		assert.Equal(checkSelect.UserID, int64(42))
		assert.Equal(checkSelect.Msg, "hello")
	}

	e = tx.Commit()
	assert.Nil(e)

}

func TestTimeColumn_workflow(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)
	type TestTimeColumnStruct struct {
		Id      int64 `sql:"PK AI"`
		Expired time.Time
	}
	var testStruct TestTimeColumnStruct
	pqaction.Register(tx, &testStruct)

	// insert time
	now := timeutil.Now()
	testStruct.Expired = now
	e := tx.Insert(&testStruct)
	assert.Nil(e)

	// select time
	{
		var toSelect TestTimeColumnStruct
		e := pqaction.SelectEntityById(tx, &toSelect, testStruct.Id)
		assert.Nil(e)
		assert.Equal(toSelect.Id, testStruct.Id)
		assert.Equal(toSelect.Expired, now)
	}

	// update time
	now = timeutil.AddDays(now, 1)
	testStruct.Expired = now
	e = tx.Update(&testStruct)
	assert.Nil(e)

	// select new time
	{
		var toSelect TestTimeColumnStruct
		e := pqaction.SelectEntityById(tx, &toSelect, testStruct.Id)
		assert.Nil(e)
		assert.Equal(toSelect.Id, testStruct.Id)
		assert.Equal(toSelect.Expired, now)
	}

}

func TestTimeColumn_withSelectOperations(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)
	type TestTimeColumnWithSelectOperationsStruct struct {
		Id   int64 `sql:"PK AI"`
		Time time.Time
	}
	var testStruct TestTimeColumnWithSelectOperationsStruct
	pqaction.Register(tx, &testStruct)

	// insert time
	now := timeutil.Now()
	testStruct.Time = now
	e := tx.Insert(&testStruct)
	assert.Nil(e)

	// select time -> equal
	{
		args := pqarg.New()
		result, e := tx.Query(
			"Select id FROM "+tableName(&testStruct)+" "+
				"WHERE time = "+args.Next(now), args)
		assert.Nil(e)

		assert.True(result.Next())
		var id int64
		e = result.Scan(&id)
		assert.Nil(e)
		assert.Equal(id, testStruct.Id)
		assert.False(result.Next())
	}

	after := timeutil.AddMinute(now, 1)
	// select time -> select smaller
	{
		args := pqarg.New()
		result, e := tx.Query(
			"Select id FROM "+tableName(&testStruct)+" "+
				"WHERE time < "+args.Next(after), args)
		assert.Nil(e)

		assert.True(result.Next())
		var id int64
		e = result.Scan(&id)
		assert.Nil(e)
		assert.Equal(id, testStruct.Id)
		assert.False(result.Next())
	}

	// select time -> select bigger -> no result
	{
		args := pqarg.New()
		result, e := tx.Query(
			"Select id FROM "+tableName(&testStruct)+" "+
				"WHERE time > "+args.Next(after), args)
		assert.Nil(e)
		assert.False(result.Next())
	}
}

func BenchmarkInsertStatement(b *testing.B) {
	tx, assert := pqtest.InitTransactionTest(b)

	type BenchmarkInsertStruct struct {
		Id     int64 `sql:"PK AI"`
		Msg    string
		UserId int64
	}

	err := pqaction.Register(tx, &BenchmarkInsertStruct{})
	assert.Nil(err)

	// run the insert function b.N times
	for n := 0; n < 10; n++ {
		toInsert := BenchmarkInsertStruct{
			Msg: "Msg" + strconv.Itoa(n),
		}
		tx.Insert(&toInsert)
	}
}

//func BenchmarkPlainInsert(b *testing.B) {
//	pqaction.OpenDatabaseConnection(configbase.GetConfig().Database[0])
//	pq := pqaction.New(log.Logger{})
//
//	type BenchmarkPlainInsertStruct struct {
//		Id     int64 `sql:"PK AI"`
//		Msg    string
//		UserId int64
//	}
//
//	testStruct := BenchmarkPlainInsertStruct{
//		Msg: "Blob",
//		UserId:42,
//	}
//	pqaction.Register(pq, &testStruct)
//
//	// run the insert function b.N times
//	for n := 0; n < b.N; n++ {
//		testStruct
//
//	}
//}

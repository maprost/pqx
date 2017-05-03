package pqx_test

import (
	"database/sql"
	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqtest"
	"github.com/maprost/pqx/pqutil"
	"testing"
)

func TestSimpleWorkflow_tx(t *testing.T) {
	tx, assert := pqtest.InitTransactionTest(t)
	tx.AddLogger(pqutil.DefaultLogger)

	type TestSimpleWorkflowTxStruct struct {
		Id     int64 `pqx:"PK AI"`
		Msg    string
		UserId int64
	}
	tx.Register(TestSimpleWorkflowTxStruct{})

	// insert entity
	entity := TestSimpleWorkflowTxStruct{Msg: "hello", UserId: 42}
	err := tx.Insert(&entity)
	assert.Nil(err)
	assert.Equal(entity.Id, int64(1))

	// select entity -> 1, "hello", 42
	{
		var checkSelect TestSimpleWorkflowTxStruct
		ok, err := tx.SelectByKeyValue("id", 1, &checkSelect)
		assert.Nil(err)
		assert.True(ok)
		assert.Equal(checkSelect.Id, int64(1))
		assert.Equal(checkSelect.Msg, "hello")
		assert.Equal(checkSelect.UserId, int64(42))
	}

	// update entity
	entity.Msg = "world"
	err = tx.Update(&entity)
	assert.Nil(err)

	// select entity -> 1, "world", 42
	{
		checkSelect := TestSimpleWorkflowTxStruct{Id: 1}
		ok, err := tx.Select(&checkSelect)
		assert.Nil(err)
		assert.True(ok)
		assert.Equal(checkSelect.Id, int64(1))
		assert.Equal(checkSelect.Msg, "world")
		assert.Equal(checkSelect.UserId, int64(42))
	}

	// delete entity
	err = tx.Delete(&entity)
	assert.Nil(err)

	// select entity -> nothing found
	{
		contains, err := tx.Contains(&TestSimpleWorkflowTxStruct{Id: 1})
		assert.Nil(err)
		assert.False(contains)
	}

	err = tx.Commit()
	assert.Nil(err)
}

func TestSimpleWorkflow_db(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type TestSimpleWorkflowDBStruct struct {
		Id     int64 `pqx:"PK AI"`
		Msg    string
		UserId int64
	}
	var testStruct TestSimpleWorkflowDBStruct
	pqx.Register(&testStruct)

	// insert entity
	entity := TestSimpleWorkflowDBStruct{Msg: "hello", UserId: 42}
	err := pqx.Insert(&entity)
	assert.Nil(err)
	assert.Equal(entity.Id, int64(1))

	// select entity -> 1, "hello", 42
	{
		var checkSelect TestSimpleWorkflowDBStruct
		ok, err := pqx.SelectByKeyValue("id", 1, &checkSelect)
		assert.Nil(err)
		assert.True(ok)
		assert.Equal(checkSelect.Id, int64(1))
		assert.Equal(checkSelect.Msg, "hello")
		assert.Equal(checkSelect.UserId, int64(42))
	}

	// update entity
	entity.Msg = "world"
	err = pqx.Update(&entity)
	assert.Nil(err)

	// select entity -> 1, "world", 42
	{
		checkSelect := TestSimpleWorkflowDBStruct{Id: 1}
		ok, err := pqx.Select(&checkSelect)
		assert.Nil(err)
		assert.True(ok)
		assert.Equal(checkSelect.Id, int64(1))
		assert.Equal(checkSelect.Msg, "world")
		assert.Equal(checkSelect.UserId, int64(42))
	}

	// delete entity
	err = pqx.Delete(&entity)
	assert.Nil(err)

	// select entity -> nothing found
	{
		contains, err := pqx.Contains(&TestSimpleWorkflowDBStruct{Id: 1})
		assert.Nil(err)
		assert.False(contains)
	}
}

func TestSimpleWorkflow_allTypes(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type TestSimpleWorkflowAllTypesStruct struct {
		Id     int64 `pqx:"PK AI"`
		Msg    string
		Blob   int
		UserId sql.NullInt64
	}

	err := pqx.Register(TestSimpleWorkflowAllTypesStruct{})
	assert.Nil(err)

	// insert entity
	entity := TestSimpleWorkflowAllTypesStruct{Msg: "hello", Blob: 42}
	err = pqx.Insert(&entity)
	assert.Nil(err)
	assert.Equal(entity.Id, int64(1))

	// select entity -> 1, "hello", nil
	{
		var checkSelect TestSimpleWorkflowAllTypesStruct
		ok, err := pqx.SelectByKeyValue("id", 1, &checkSelect)
		assert.Nil(err)
		assert.True(ok)
		assert.Equal(checkSelect.Id, int64(1))
		assert.Equal(checkSelect.Msg, "hello")
		assert.Equal(checkSelect.Blob, 42)
		assert.False(checkSelect.UserId.Valid)
	}

	// update entity
	entity.Msg = "world"
	entity.UserId = sql.NullInt64{Int64: 42, Valid: true}
	err = pqx.Update(&entity)
	assert.Nil(err)

	// select entity -> 1, "world", 42
	{
		checkSelect := TestSimpleWorkflowAllTypesStruct{Id: 1}
		ok, err := pqx.Select(&checkSelect)
		assert.Nil(err)
		assert.True(ok)
		assert.Equal(checkSelect.Id, int64(1))
		assert.Equal(checkSelect.Msg, "world")
		assert.True(checkSelect.UserId.Valid)
		assert.Equal(checkSelect.UserId.Int64, int64(42))
	}

	// delete entity
	err = pqx.Delete(&entity)
	assert.Nil(err)

	// select entity -> nothing found
	{
		contains, err := pqx.Contains(TestSimpleWorkflowAllTypesStruct{Id: 1})
		assert.Nil(err)
		assert.False(contains)
	}
}

//func TestUpdateWithoutID(t *testing.T) {
//	tx, assert := pqtest.InitTransactionTest(t)
//
//	type TestUpdateWithoutIDStruct struct {
//		UserID int64
//		Msg    string
//	}
//	var testStruct TestUpdateWithoutIDStruct
//	tx.Register(&testStruct)
//
//	// insert entity
//	entity := TestUpdateWithoutIDStruct{Msg: "hello", UserID: 42}
//	e := tx.Insert(&entity)
//	assert.Nil(e)
//
//	// select entity -> 42, "hello"
//	{
//		result, e := tx.Query("SELECT "+pqx.SelectList(&entity)+" FROM "+
//			tableName(&testStruct), pqarg.New())
//		assert.Nil(e)
//
//		var checkSelect TestUpdateWithoutIDStruct
//		e = pqx.ScanStruct(result, &checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.UserID, int64(42))
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//
//	// try to update entity -> eor
//	entity.Msg = "world"
//	e = tx.Update(&entity)
//	assert.NotNil(e)
//	assert.Equal(e.Error(), "No primary key available.")
//
//	// select entity -> 42, "hello"
//	{
//		result, e := tx.Query("SELECT "+pqx.SelectList(&entity)+
//			" FROM "+tableName(&testStruct), pqarg.New())
//		assert.Nil(e)
//
//		var checkSelect TestUpdateWithoutIDStruct
//		e = pqx.ScanStruct(result, &checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.UserID, int64(42))
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//
//	// try to delete entity -> eor
//	e = tx.Delete(&entity)
//	assert.NotNil(e)
//	assert.Equal(e.Error(), "No primary key available.")
//
//	// select entity -> 42, "hello"
//	{
//		result, e := tx.Query("SELECT "+pqx.SelectList(&entity)+
//			" FROM "+tableName(&testStruct), pqarg.New())
//		assert.Nil(e)
//
//		var checkSelect TestUpdateWithoutIDStruct
//		e = pqx.ScanStruct(result, &checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.UserID, int64(42))
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//
//	e = tx.Commit()
//	assert.Nil(e)
//
//}
//
//func TestTimeColumn_workflow(t *testing.T) {
//	tx, assert := pqtest.InitTransactionTest(t)
//	type TestTimeColumnStruct struct {
//		Id      int64 `sql:"PK AI"`
//		Expired time.Time
//	}
//	var testStruct TestTimeColumnStruct
//	tx.Register(&testStruct)
//
//	// insert time
//	now := timeutil.Now()
//	testStruct.Expired = now
//	e := tx.Insert(&testStruct)
//	assert.Nil(e)
//
//	// select time
//	{
//		toSelect := TestTimeColumnStruct{Id:testStruct.Id}
//		e := tx.Select(&toSelect)
//		assert.Nil(e)
//		assert.Equal(toSelect.Id, testStruct.Id)
//		assert.Equal(toSelect.Expired, now)
//	}
//
//	// update time
//	now = timeutil.AddDays(now, 1)
//	testStruct.Expired = now
//	e = tx.Update(&testStruct)
//	assert.Nil(e)
//
//	// select new time
//	{
//		toSelect := TestTimeColumnStruct{Id:testStruct.Id}
//		e := tx.Select(&toSelect)
//		assert.Nil(e)
//		assert.Equal(toSelect.Id, testStruct.Id)
//		assert.Equal(toSelect.Expired, now)
//	}
//
//}
//
//func TestTimeColumn_withSelectOperations(t *testing.T) {
//	tx, assert := pqtest.InitTransactionTest(t)
//	type TestTimeColumnWithSelectOperationsStruct struct {
//		Id   int64 `sql:"PK AI"`
//		Time time.Time
//	}
//	var testStruct TestTimeColumnWithSelectOperationsStruct
//	tx.Register(&testStruct)
//
//	// insert time
//	now := timeutil.Now()
//	testStruct.Time = now
//	e := tx.Insert(&testStruct)
//	assert.Nil(e)
//
//	// select time -> equal
//	{
//		args := pqarg.New()
//		result, e := tx.Query(
//			"Select id FROM "+tableName(&testStruct)+" "+
//				"WHERE time = "+args.Next(now), args)
//		assert.Nil(e)
//
//		assert.True(result.Next())
//		var id int64
//		e = result.Scan(&id)
//		assert.Nil(e)
//		assert.Equal(id, testStruct.Id)
//		assert.False(result.Next())
//	}
//
//	after := timeutil.AddMinute(now, 1)
//	// select time -> select smaller
//	{
//		args := pqarg.New()
//		result, e := tx.Query(
//			"Select id FROM "+tableName(&testStruct)+" "+
//				"WHERE time < "+args.Next(after), args)
//		assert.Nil(e)
//
//		assert.True(result.Next())
//		var id int64
//		e = result.Scan(&id)
//		assert.Nil(e)
//		assert.Equal(id, testStruct.Id)
//		assert.False(result.Next())
//	}
//
//	// select time -> select bigger -> no result
//	{
//		args := pqarg.New()
//		result, e := tx.Query(
//			"Select id FROM "+tableName(&testStruct)+" "+
//				"WHERE time > "+args.Next(after), args)
//		assert.Nil(e)
//		assert.False(result.Next())
//	}
//}
//
//func BenchmarkInsertStatement(b *testing.B) {
//	tx, assert := pqtest.InitTransactionTest(b)
//
//	type BenchmarkInsertStruct struct {
//		Id     int64 `sql:"PK AI"`
//		Msg    string
//		UserId int64
//	}
//
//	err := tx.Register(&BenchmarkInsertStruct{})
//	assert.Nil(err)
//
//	// run the insert function b.N times
//	for n := 0; n < 10; n++ {
//		toInsert := BenchmarkInsertStruct{
//			Msg: "Msg" + strconv.Itoa(n),
//		}
//		tx.Insert(&toInsert)
//	}
//}

//func BenchmarkPlainInsert(b *testing.B) {
//	pqx.OpenDatabaseConnection(configbase.GetConfig().Database[0])
//	pq := pqx.New(log.Logger{})
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
//	pqx.Register(pq, &testStruct)
//
//	// run the insert function b.N times
//	for n := 0; n < b.N; n++ {
//		testStruct
//
//	}
//}

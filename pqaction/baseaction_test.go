package pqaction_test
//
//import (
//	"log"
//	"rpp.de/BackendLib/base/configbase"
//	"rpp.de/BackendLib/base/dbbase/postgres"
//	"rpp.de/BackendLib/base/testbase/assert"
//	"rpp.de/BackendLib/util/timeutil"
//	"testing"
//	"time"
//)
//
//func TestCreateSelect_simple(t *testing.T) {
//	db := initTest(t)
//
//	type TestCreateSelectStruct struct {
//		Id  int64
//		Msg string
//	}
//	var testStruct TestCreateSelectStruct
//	postgres.Register(db, &testStruct)
//}
//
//func TestSimpleWorkflow(t *testing.T) {
//	db := initTest(t)
//
//	type TestSimpleWorkflowStruct struct {
//		Id     int64 `sql:"PK AI"`
//		Msg    string
//		UserId int64
//	}
//	var testStruct TestSimpleWorkflowStruct
//	postgres.Register(db, &testStruct)
//
//	// insert entity
//	entity := TestSimpleWorkflowStruct{Msg: "hello", UserId: 42}
//	e := postgres.Insert(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.Id, int64(1))
//
//	// select entity -> 1, "hello", 42
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE id = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestSimpleWorkflowStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.Id, int64(1))
//		assert.Equal(checkSelect.Msg, "hello")
//		assert.Equal(checkSelect.UserId, int64(42))
//	}
//
//	// update entity
//	entity.Msg = "world"
//	e = postgres.Update(db, &entity)
//	assert.Nil(e)
//
//	// select entity -> 1, "world", 42
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE id = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestSimpleWorkflowStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.Id, int64(1))
//		assert.Equal(checkSelect.Msg, "world")
//		assert.Equal(checkSelect.UserId, int64(42))
//	}
//
//	// delete entity
//	e = postgres.Delete(db, &entity)
//	assert.Nil(e)
//
//	// select entity -> nothing found
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE id = "+args.Next(1), args)
//		assert.Nil(e)
//		assert.False(result.HasNext())
//	}
//
//	e = db.Commit()
//	assert.Nil(e)
//}
//
//func TestUpdateWithoutID(t *testing.T) {
//	db := initTest(t)
//
//	type TestUpdateWithoutIDStruct struct {
//		UserID int64
//		Msg    string
//	}
//	var testStruct TestUpdateWithoutIDStruct
//	postgres.Register(db, &testStruct)
//
//	// insert entity
//	entity := TestUpdateWithoutIDStruct{Msg: "hello", UserID: 42}
//	e := postgres.Insert(db, &entity)
//	assert.Nil(e)
//
//	// select entity -> 42, "hello"
//	{
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct), postgres.NewArgs())
//		assert.Nil(e)
//
//		var checkSelect TestUpdateWithoutIDStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.UserID, int64(42))
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//
//	// try to update entity -> eor
//	entity.Msg = "world"
//	e = postgres.Update(db, &entity)
//	assert.NotNil(e)
//	assert.Equal(e, postgres.DatabaseNoPrimaryKeyAvailable())
//
//	// select entity -> 42, "hello"
//	{
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+
//			" FROM "+postgres.TableName(&testStruct), postgres.NewArgs())
//		assert.Nil(e)
//
//		var checkSelect TestUpdateWithoutIDStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.UserID, int64(42))
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//
//	// try to delete entity -> eor
//	e = postgres.Delete(db, &entity)
//	assert.NotNil(e)
//	assert.Equal(e, postgres.DatabaseNoPrimaryKeyAvailable())
//
//	// select entity -> 42, "hello"
//	{
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+
//			" FROM "+postgres.TableName(&testStruct), postgres.NewArgs())
//		assert.Nil(e)
//
//		var checkSelect TestUpdateWithoutIDStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.UserID, int64(42))
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//
//	e = db.Commit()
//	assert.Nil(e)
//
//}
//
//func TestTimeColumn_workflow(t *testing.T) {
//	db := initTest(t)
//	type TestTimeColumnStruct struct {
//		Id      int64 `sql:"PK AI"`
//		Expired time.Time
//	}
//	var testStruct TestTimeColumnStruct
//	postgres.Register(db, &testStruct)
//
//	// insert time
//	now := timeutil.Now()
//	testStruct.Expired = now
//	e := postgres.Insert(db, &testStruct)
//	assert.Nil(e)
//
//	// select time
//	{
//		var toSelect TestTimeColumnStruct
//		e := postgres.GetSingleEntityById(db, &toSelect, testStruct.Id)
//		assert.Nil(e)
//		assert.Equal(toSelect.Id, testStruct.Id)
//		assert.Equal(toSelect.Expired, now)
//	}
//
//	// update time
//	now = timeutil.AddDays(now, 1)
//	testStruct.Expired = now
//	e = postgres.Update(db, &testStruct)
//	assert.Nil(e)
//
//	// select new time
//	{
//		var toSelect TestTimeColumnStruct
//		e := postgres.GetSingleEntityById(db, &toSelect, testStruct.Id)
//		assert.Nil(e)
//		assert.Equal(toSelect.Id, testStruct.Id)
//		assert.Equal(toSelect.Expired, now)
//	}
//
//}
//
//func TestTimeColumn_withSelectOperations(t *testing.T) {
//	db := initTest(t)
//	type TestTimeColumnWithSelectOperationsStruct struct {
//		Id   int64 `sql:"PK AI"`
//		Time time.Time
//	}
//	var testStruct TestTimeColumnWithSelectOperationsStruct
//	postgres.Register(db, &testStruct)
//
//	// insert time
//	now := timeutil.Now()
//	testStruct.Time = now
//	e := postgres.Insert(db, &testStruct)
//	assert.Nil(e)
//
//	// select time -> equal
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query(
//			"Select id FROM "+postgres.TableName(&testStruct)+" "+
//				"WHERE time = "+args.Next(now), args)
//		assert.Nil(e)
//
//		assert.True(result.HasNext())
//		var id int64
//		e = result.Scan(&id)
//		assert.Nil(e)
//		assert.Equal(id, testStruct.Id)
//		assert.False(result.HasNext())
//	}
//
//	after := timeutil.AddMinute(now, 1)
//	// select time -> select smaller
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query(
//			"Select id FROM "+postgres.TableName(&testStruct)+" "+
//				"WHERE time < "+args.Next(after), args)
//		assert.Nil(e)
//
//		assert.True(result.HasNext())
//		var id int64
//		e = result.Scan(&id)
//		assert.Nil(e)
//		assert.Equal(id, testStruct.Id)
//		assert.False(result.HasNext())
//	}
//
//	// select time -> select bigger -> no result
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query(
//			"Select id FROM "+postgres.TableName(&testStruct)+" "+
//				"WHERE time > "+args.Next(after), args)
//		assert.Nil(e)
//		assert.False(result.HasNext())
//	}
//}
//
//func BenchmarkInsertStatement(b *testing.B) {
//	postgres.OpenDatabaseConnection(configbase.GetConfig().Database[0])
//	pq := postgres.New(log.Logger{})
//
//	type BenchmarkInsertStruct struct {
//		Id     int64 `sql:"PK AI"`
//		Msg    string
//		UserId int64
//	}
//
//	testStruct := BenchmarkInsertStruct{
//		Msg:    "Blob",
//		UserId: 42,
//	}
//	postgres.Register(pq, &testStruct)
//
//	// run the insert function b.N times
//	for n := 0; n < b.N; n++ {
//		testStruct
//		postgres.Insert(pq, &testStruct)
//	}
//}
//
////func BenchmarkPlainInsert(b *testing.B) {
////	postgres.OpenDatabaseConnection(configbase.GetConfig().Database[0])
////	pq := postgres.New(log.Logger{})
////
////	type BenchmarkPlainInsertStruct struct {
////		Id     int64 `sql:"PK AI"`
////		Msg    string
////		UserId int64
////	}
////
////	testStruct := BenchmarkPlainInsertStruct{
////		Msg: "Blob",
////		UserId:42,
////	}
////	postgres.Register(pq, &testStruct)
////
////	// run the insert function b.N times
////	for n := 0; n < b.N; n++ {
////		testStruct
////
////	}
////}

package pqaction_test
//
//import (
//	"rpp.de/BackendLib/base/dbbase/postgres"
//	"rpp.de/BackendLib/base/testbase/assert"
//	"rpp.de/BackendLib/util/timeutil"
//	"testing"
//	"time"
//)
//
//func TestPKWithAI_workflow(t *testing.T) {
//	db := initTest(t)
//
//	type TestPKWithAIWorkflowStruct struct {
//		PrimKey int `sql:"PK AI"`
//		Msg     string
//	}
//	var testStruct TestPKWithAIWorkflowStruct
//	postgres.Register(db, &testStruct)
//
//	// insert entity
//	entity := TestPKWithAIWorkflowStruct{Msg: "hello"}
//	e := postgres.Insert(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 1)
//
//	// select entity -> 1, "hello"
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestPKWithAIWorkflowStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 1)
//		assert.Equal(checkSelect.Msg, "hello")
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
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestPKWithAIWorkflowStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 1)
//		assert.Equal(checkSelect.Msg, "world")
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
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//		assert.False(result.HasNext())
//	}
//
//	e = db.Commit()
//	assert.Nil(e)
//}
//
//func TestPKWithoutAI(t *testing.T) {
//	db := initTest(t)
//
//	type TestPKWithoutAIStruct struct {
//		PrimKey int `sql:"PK"`
//		Msg     string
//	}
//	var testStruct TestPKWithoutAIStruct
//	postgres.Register(db, &testStruct)
//
//	// insert entity
//	entity := TestPKWithoutAIStruct{PrimKey: -1, Msg: "hello"}
//	e := postgres.Insert(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, -1)
//
//	// select entity -> -1, "hello"
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(-1), args)
//		assert.Nil(e)
//
//		var checkSelect TestPKWithoutAIStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, -1)
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//}
//
//func TestPKWithoutAIAndForgetToSetValue(t *testing.T) {
//	db := initTest(t)
//
//	type TestPKWithoutAIAndForgetToSetValueStruct struct {
//		PrimKey int `sql:"PK"`
//		Msg     string
//	}
//	var testStruct TestPKWithoutAIAndForgetToSetValueStruct
//	postgres.Register(db, &testStruct)
//
//	// insert entity
//	entity := TestPKWithoutAIAndForgetToSetValueStruct{Msg: "hello"}
//	e := postgres.Insert(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 0) // default value
//
//	// select entity -> 0, "hello"
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(0), args)
//		assert.Nil(e)
//
//		var checkSelect TestPKWithoutAIAndForgetToSetValueStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 0)
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//}
//
//func TestAIWithMultiplyInserts(t *testing.T) {
//	db := initTest(t)
//
//	type TestAIWithMultiplyInsertsStruct struct {
//		PrimKey int `sql:"AI"`
//	}
//	var testStruct TestAIWithMultiplyInsertsStruct
//	postgres.Register(db, &testStruct)
//
//	// insert entities and select them
//	for id := 1; id <= 10; id++ {
//		entity := TestAIWithMultiplyInsertsStruct{}
//		e := postgres.Insert(db, &entity)
//		assert.Nil(e)
//		assert.Equal(entity.PrimKey, id)
//
//		// select entity
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(id), args)
//		assert.Nil(e)
//
//		var checkSelect TestAIWithMultiplyInsertsStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, id)
//	}
//}
//
//func TestAIAndSetValue(t *testing.T) {
//	db := initTest(t)
//
//	type TestAIAndSetValueStruct struct {
//		PrimKey int `sql:"AI"`
//	}
//	var testStruct TestAIAndSetValueStruct
//	postgres.Register(db, &testStruct)
//
//	// insert entities
//	entity := TestAIAndSetValueStruct{PrimKey: 10}
//	e := postgres.Insert(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 1)
//
//	{
//		// select entity: 1
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestAIAndSetValueStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 1)
//	}
//
//	{
//		// select entity: 10
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(10), args)
//		assert.Nil(e)
//		assert.False(result.HasNext())
//	}
//}
//
//func TestCreated(t *testing.T) {
//	db := initTest(t)
//
//	type TestCreatedStruct struct {
//		PrimKey int       `sql:"PK AI"`
//		Created time.Time `sql:"CreateDate"`
//		Msg     string
//	}
//	var testStruct TestCreatedStruct
//	postgres.Register(db, &testStruct)
//
//	now := timeutil.Now()
//	yesterday := timeutil.AddDays(now, -1)
//	timeutil.Init(yesterday)
//
//	// insert entities
//	entity := TestCreatedStruct{Msg: "hello"}
//	e := postgres.Insert(db, &entity) // set 'Created' to timebase.now()
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 1)
//	assert.Equal(entity.Created, yesterday)
//	assert.Equal(entity.Msg, "hello")
//
//	{
//		// select entity: 1, yesterday, "hello"
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestCreatedStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 1)
//		assert.Equal(checkSelect.Created, yesterday)
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//
//	timeutil.Init(now)
//	entity.Msg = "world"
//	e = postgres.Update(db, &entity) // no change of 'Created'
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 1)
//	assert.Equal(entity.Created, yesterday)
//	assert.Equal(entity.Msg, "world")
//
//	{
//		// select entity: 1, yesterday, "world"
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestCreatedStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 1)
//		assert.Equal(checkSelect.Created, yesterday)
//		assert.Equal(checkSelect.Msg, "world")
//	}
//}
//
//func TestCreatedWithPreSetting(t *testing.T) {
//	db := initTest(t)
//
//	type TestCreatedWithPreSettingStruct struct {
//		PrimKey int       `sql:"PK AI"`
//		Created time.Time `sql:"CreateDate"`
//	}
//	var testStruct TestCreatedWithPreSettingStruct
//	postgres.Register(db, &testStruct)
//
//	today := timeutil.Now()
//	tomorrow := timeutil.AddDays(today, 1)
//	timeutil.Init(today)
//
//	// insert entities
//	entity := TestCreatedWithPreSettingStruct{Created: tomorrow}
//	e := postgres.Insert(db, &entity) // change 'Created' with timebase.now()
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 1)
//	assert.Equal(entity.Created, today)
//
//	// select entity: 1, today
//	args := postgres.NewArgs()
//	result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//		postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//	assert.Nil(e)
//
//	var checkSelect TestCreatedWithPreSettingStruct
//	e = result.ScanStruct(&checkSelect)
//	assert.Nil(e)
//	assert.Equal(checkSelect.PrimKey, 1)
//	assert.Equal(checkSelect.Created, today)
//}
//
//func TestChanged(t *testing.T) {
//	db := initTest(t)
//
//	type TestChangedStruct struct {
//		PrimKey int       `sql:"PK AI"`
//		Changed time.Time `sql:"ChangeDate"`
//		Msg     string
//	}
//	var testStruct TestChangedStruct
//	postgres.Register(db, &testStruct)
//
//	now := timeutil.Now()
//	yesterday := timeutil.AddDays(now, -1)
//	timeutil.Init(yesterday)
//
//	// insert entities
//	entity := TestChangedStruct{Msg: "hello"}
//	e := postgres.Insert(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 1)
//	assert.Equal(entity.Changed, yesterday)
//	assert.Equal(entity.Msg, "hello")
//
//	{
//		// select entity: 1, yesterday
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestChangedStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 1)
//		assert.Equal(checkSelect.Changed, yesterday)
//		assert.Equal(checkSelect.Msg, "hello")
//	}
//
//	timeutil.Init(now)
//	entity.Msg = "world"
//	e = postgres.Update(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 1)
//	assert.Equal(entity.Changed, now)
//	assert.Equal(entity.Msg, "world")
//
//	{
//		// select entity: 1, yesterday
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestChangedStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 1)
//		assert.Equal(checkSelect.Changed, now)
//		assert.Equal(checkSelect.Msg, "world")
//	}
//}
//
//func TestChangedWithPreSetting(t *testing.T) {
//	db := initTest(t)
//
//	type TestChangedWithPreSettingStruct struct {
//		PrimKey int       `sql:"PK AI"`
//		Changed time.Time `sql:"ChangeDate"`
//		Msg     string
//	}
//	var testStruct TestChangedWithPreSettingStruct
//	postgres.Register(db, &testStruct)
//
//	today := timeutil.Now()
//	tomorrow := timeutil.AddDays(today, 1)
//	timeutil.Init(today)
//
//	// insert entities
//	entity := TestChangedWithPreSettingStruct{Changed: tomorrow, Msg: "hello"}
//	e := postgres.Insert(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 1)
//	assert.Equal(entity.Changed, today)
//	assert.Equal(entity.Msg, "hello")
//
//	{
//		// select entity: 1, now
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestChangedWithPreSettingStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 1)
//		assert.Equal(checkSelect.Changed, today)
//		assert.Equal(entity.Msg, "hello")
//	}
//
//	entity.Changed = tomorrow
//	entity.Msg = "world"
//	e = postgres.Update(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.PrimKey, 1)
//	assert.Equal(entity.Changed, today)
//	assert.Equal(entity.Msg, "world")
//
//	{
//		// select entity: 1, now
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE primkey = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestChangedWithPreSettingStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.PrimKey, 1)
//		assert.Equal(checkSelect.Changed, today)
//		assert.Equal(entity.Msg, "world")
//	}
//}
//
//func TestAllTypesWorkflow(t *testing.T) {
//	db := initTest(t)
//
//	type TestAllTypesWorkflowStruct struct {
//		Id      int64     `sql:"PK AI"`
//		Created time.Time `sql:"CreateDate"`
//		Changed time.Time `sql:"ChangeDate"`
//		Msg     string
//	}
//	var testStruct TestAllTypesWorkflowStruct
//	postgres.Register(db, &testStruct)
//
//	today := timeutil.Now()
//	tomorrow := timeutil.AddDays(today, 1)
//	timeutil.Init(today)
//
//	// insert entity
//	entity := TestAllTypesWorkflowStruct{Msg: "hello"}
//	e := postgres.Insert(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.Id, int64(1))
//	assert.Equal(entity.Created, today)
//	assert.Equal(entity.Changed, today)
//
//	// select entity -> 1, "hello"
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE id = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestAllTypesWorkflowStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.Id, int64(1))
//		assert.Equal(checkSelect.Msg, "hello")
//		assert.Equal(checkSelect.Created, today)
//		assert.Equal(checkSelect.Changed, today)
//	}
//
//	// update entity
//	entity.Msg = "world"
//	timeutil.Init(tomorrow)
//	e = postgres.Update(db, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.Created, today)
//	assert.Equal(entity.Changed, tomorrow)
//
//	// select entity -> 1, "world"
//	{
//		args := postgres.NewArgs()
//		result, e := db.Query("SELECT "+postgres.SelectList(&entity)+" FROM "+
//			postgres.TableName(&testStruct)+" WHERE id = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestAllTypesWorkflowStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.Id, int64(1))
//		assert.Equal(checkSelect.Msg, "world")
//		assert.Equal(checkSelect.Created, today)
//		assert.Equal(checkSelect.Changed, tomorrow)
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

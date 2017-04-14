package pqx_test

//import (
//	"database/sql"
//	"github.com/maprost/pqx/pqtest"
//	"testing"
//	"time"
//)
//
//func TestNullableType(t *testing.T) {
//	tx, assert := pqtest.InitTransactionTest(t)
//
//	type ID int
//	type TestUnknownTypeStruct struct {
//		Id      ID        `sql:"PK AI"`
//		Created time.Time `sql:"CreateDate"`
//		Msg     sql.NullString
//		Flag    sql.NullBool
//		Sum     sql.NullFloat64
//		Counter sql.NullInt64
//	}
//	var testStruct TestUnknownTypeStruct
//	err := pqaction.Register(tx, &testStruct)
//	assert.Nil(err)
//
//	// insert entity
//	entity := TestUnknownTypeStruct{
//		Msg:     sql.NullString{Valid: false},
//		Flag:    sql.NullBool{Valid: false},
//		Sum:     sql.NullFloat64{Valid: false},
//		Counter: sql.NullInt64{Valid: false},
//	}
//	e := pqaction.InsertTx(tx, &entity)
//	assert.Nil(e)
//	assert.Equal(entity.Id, ID(1))
//
//	// select entity -> 1, all attributes are nil
//	{
//		args := pqlib.NewArgs()
//		result, e := tx.Query("SELECT "+pqaction.SelectList(&entity)+" FROM "+
//			tableName(&testStruct)+" WHERE id = "+args.Next(1), args)
//		assert.Nil(e)
//
//		var checkSelect TestUnknownTypeStruct
//		e = result.ScanStruct(&checkSelect)
//		assert.Nil(e)
//		assert.Equal(checkSelect.Id, ID(1))
//		assert.False(checkSelect.Msg.Valid)     // nil
//		assert.False(checkSelect.Flag.Valid)    // nil
//		assert.False(checkSelect.Sum.Valid)     // nil
//		assert.False(checkSelect.Counter.Valid) // nil
//	}
//}

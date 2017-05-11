package pqx_test

import (
	"github.com/maprost/assertion"
	"testing"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqnull"
)

// 10000	    137.949 ns/op
// 10000	    136.525 ns/op
func BenchmarkSelect_threeElements_pq(b *testing.B) {
	assert := assertion.New(b)

	entity := ThreeElementsStruct{
		Msg:    "hello world",
		UserId: pqnull.ValidInt64(42),
	}

	err := pqx.LogInsert(&entity, noLogging{})
	assert.Nil(err)

	for i := 0; i < b.N; i++ {
		sql := "SELECT id, msg, userid FROM ThreeElementsStruct WHERE id = $1;"
		rows, _ := pqx.DB.Query(sql, entity.Id)
		assert.True(rows.Next())

		var id int64
		var msg string
		var userId pqnull.Int64
		err := rows.Scan(&id, &msg, &userId)
		assert.Nil(err)
		assert.Equal(id, entity.Id)
		assert.Equal(msg, "hello world")
		assert.Equal(userId, pqnull.ValidInt64(42))

		rows.Close()
	}
}

// 10000	    142.683 ns/op
// 10000	    143.076 ns/op
func BenchmarkSelect_threeElements_pqx(b *testing.B) {
	assert := assertion.New(b)

	entity := ThreeElementsStruct{
		Msg:    "hello world",
		UserId: pqnull.ValidInt64(42),
	}

	err := pqx.LogInsert(&entity, noLogging{})
	assert.Nil(err)

	for i := 1; i < b.N; i++ {
		entityToSearch := ThreeElementsStruct{Id: entity.Id}
		ok, err := pqx.LogSelect(&entityToSearch, noLogging{})
		assert.Nil(err)
		assert.True(ok)
		assert.Equal(entityToSearch.Id, entity.Id)
		assert.Equal(entityToSearch.Msg, "hello world")
		assert.Equal(entityToSearch.UserId, pqnull.ValidInt64(42))
	}
}

// 10000	    137.949 ns/op
// 10000	    136.525 ns/op
//func BenchmarkSelectList_threeElements_pq(b *testing.B) {
//	log.Println("BenchmarkSelect_threeElements_pq")
//	assert := assertion.New(b)
//
//	type SelectListThreeElementsPqStruct struct {
//		Id     int64 `pqx:"PK AI"`
//		Msg    string
//		UserId pqnull.Int64
//	}
//	err := pqx.LogRegister(SelectListThreeElementsPqStruct{}, noLogging{})
//	assert.Nil(err)
//
//	entity := SelectListThreeElementsPqStruct{
//		Msg:    "hello world",
//		UserId: pqnull.ValidInt64(42),
//	}
//
//	err := pqx.LogInsert(&entity, noLogging{})
//	assert.Nil(err)
//
//	for i := 0; i < b.N; i++ {
//		sql := "SELECT id, msg, userid FROM ThreeElementsStruct WHERE id = $1;"
//		rows, _ := pqx.DB.Query(sql, entity.Id)
//		assert.True(rows.Next())
//
//		var id int64
//		var msg string
//		var userId pqnull.Int64
//		err := rows.Scan(&id, &msg, &userId)
//		assert.Nil(err)
//		assert.Equal(id, entity.Id)
//		assert.Equal(msg, "hello world")
//		assert.Equal(userId, pqnull.ValidInt64(42))
//
//		rows.Close()
//	}
//}
//
//// 10000	    142.683 ns/op
//// 10000	    143.076 ns/op
//func BenchmarkSelectList_threeElements_pqx(b *testing.B) {
//	log.Println("BenchmarkSelect_threeElements_pqx")
//	assert := assertion.New(b)
//
//	entity := SelectListThreeElementsPqxStruct{
//		Msg:    "hello world",
//		UserId: pqnull.ValidInt64(42),
//	}
//
//	err := pqx.LogInsert(&entity, noLogging{})
//	assert.Nil(err)
//
//	for i := 1; i < b.N; i++ {
//		entityToSearch := ThreeElementsStruct{Id: entity.Id}
//		ok, err := pqx.LogSelect(&entityToSearch, noLogging{})
//		assert.Nil(err)
//		assert.True(ok)
//		assert.Equal(entityToSearch.Id, entity.Id)
//		assert.Equal(entityToSearch.Msg, "hello world")
//		assert.Equal(entityToSearch.UserId, pqnull.ValidInt64(42))
//	}
//}

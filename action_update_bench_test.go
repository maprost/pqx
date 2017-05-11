package pqx_test

import (
	"github.com/maprost/assertion"
	"testing"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqnull"
)

// 1000	   2.063.516 ns/op
// 1000	   2.095.943 ns/op
func BenchmarkUpdate_threeElements_pq(b *testing.B) {
	assert := assertion.New(b)

	entity := ThreeElementsStruct{
		Msg:    "hello world",
		UserId: pqnull.ValidInt64(42),
	}

	err := pqx.LogInsert(&entity, noLogging{})
	assert.Nil(err)

	for i := 0; i < b.N; i++ {
		sql := "UPDATE ThreeElementsStruct SET msg = $1, userid = $2 WHERE id = $3;"
		rows, err := pqx.DB.Query(sql, entity.Msg, entity.UserId, entity.Id)
		assert.Nil(err)
		rows.Close()
	}
}

// 1000	   2.126.482 ns/op
// 1000	   2.065.972 ns/op
func BenchmarkUpdate_threeElements_pqx(b *testing.B) {
	assert := assertion.New(b)

	entity := ThreeElementsStruct{
		Msg:    "hello world",
		UserId: pqnull.ValidInt64(42),
	}

	err := pqx.LogInsert(&entity, noLogging{})
	assert.Nil(err)

	for i := 1; i < b.N; i++ {
		err := pqx.LogUpdate(&entity, noLogging{})
		assert.Nil(err)
	}
}

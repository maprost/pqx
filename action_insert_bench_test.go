package pqx_test

import (
	"github.com/maprost/assertion"
	"testing"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqnull"
)

// 500	   2.107.756 ns/op
// 1000	   1.947.991 ns/op
func BenchmarkInsert_threeElements_pq(b *testing.B) {
	assert := assertion.New(b)

	entity := ThreeElementsStruct{
		Msg:    "hello world",
		UserId: pqnull.ValidInt64(42),
	}

	b.Run("pq", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			sql := "INSERT INTO ThreeElementsStruct (id, msg, userid) VALUES (DEFAULT, $1, $2) RETURNING id"
			rows, err := pqx.DB.Query(sql, entity.Msg, entity.UserId)
			assert.Nil(err)
			assert.True(rows.Next())

			var id int64
			err = rows.Scan(&id)
			assert.Nil(err)

			entity.Id = id
			err = rows.Close()
			assert.Nil(err)
		}
	})
}

// 1000	   2.066.395 ns/op
// 1000	   1.966.197 ns/op
func BenchmarkInsert_threeElements_pqx(b *testing.B) {
	assert := assertion.New(b)

	entity := ThreeElementsStruct{
		Msg:    "hello world",
		UserId: pqnull.ValidInt64(42),
	}

	b.Run("pqx", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			err := pqx.LogInsert(&entity, noLogging{})
			assert.Nil(err)
		}
	})
}

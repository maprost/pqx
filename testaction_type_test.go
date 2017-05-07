package pqx_test

import (
	"database/sql"
	"github.com/lib/pq"
	"testing"
	"time"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqnull"
	"github.com/maprost/pqx/pqtest"
)

func TestNullableType_insertValid_updateNil(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type ID int
	type TestNullableTypeInsertValidUpdateNilStruct struct {
		Id            ID        `pqx:"PK AI"`
		Created       time.Time `pqx:"Created"`
		Closed_pq     pq.NullTime
		Closed_pqx    pqnull.Time
		Msg           string
		Msg_sql       sql.NullString
		Msg_pqx       pqnull.String
		Flag          bool
		Flag_sql      sql.NullBool
		Flag_pqx      pqnull.Bool
		Counter       int
		Counter_sql   sql.NullInt64
		Counter64_pqx pqnull.Int64
		Counter_pqx   pqnull.Int
		Counter32_pqx pqnull.Int32
		Counter16_pqx pqnull.Int16
		Counter8_pqx  pqnull.Int8
		Sum           float64
		Sum_sql       sql.NullFloat64
		Sum64_pqx     pqnull.Float64
		Sum32_pqx     pqnull.Float32
	}
	err := pqx.Register(TestNullableTypeInsertValidUpdateNilStruct{})
	assert.Nil(err)

	now := time.Now().UTC().Truncate(time.Millisecond)

	// init entity with valid values
	entity := TestNullableTypeInsertValidUpdateNilStruct{
		Closed_pq:     pq.NullTime{Time: now, Valid: true},
		Closed_pqx:    pqnull.ValidTime(now),
		Msg:           "blob",
		Msg_sql:       sql.NullString{String: "blob", Valid: true},
		Msg_pqx:       pqnull.ValidString("blob"),
		Flag:          true,
		Flag_sql:      sql.NullBool{Bool: true, Valid: true},
		Flag_pqx:      pqnull.ValidBool(true),
		Counter:       42,
		Counter_sql:   sql.NullInt64{Int64: 42, Valid: true},
		Counter64_pqx: pqnull.ValidInt64(42),
		Counter_pqx:   pqnull.ValidInt(42),
		Counter32_pqx: pqnull.ValidInt32(42),
		Counter16_pqx: pqnull.ValidInt16(42),
		Counter8_pqx:  pqnull.ValidInt8(42),
		Sum:           12.3,
		Sum_sql:       sql.NullFloat64{Float64: 12.3, Valid: true},
		Sum64_pqx:     pqnull.ValidFloat64(12.3),
		Sum32_pqx:     pqnull.ValidFloat32(12.3),
	}
	e := pqx.Insert(&entity)
	assert.Nil(e)
	assert.Equal(entity.Id, ID(1))

	// select entity -> 1
	{
		checkSelect := TestNullableTypeInsertValidUpdateNilStruct{
			Id: 1,
		}
		find, err := pqx.Select(&checkSelect)
		assert.Nil(err)
		assert.True(find)

		assert.Equal(checkSelect.Id, ID(1))
		assert.Equal(checkSelect.Closed_pq.Time, now)
		assert.True(checkSelect.Closed_pq.Valid)
		assert.Equal(checkSelect.Closed_pqx.Time, now)
		assert.True(checkSelect.Closed_pqx.Valid)

		assert.Equal(checkSelect.Msg, "blob")
		assert.Equal(checkSelect.Msg_sql.String, "blob")
		assert.True(checkSelect.Msg_sql.Valid)
		assert.Equal(checkSelect.Msg_pqx.String, "blob")
		assert.True(checkSelect.Msg_pqx.Valid)

		assert.Equal(checkSelect.Flag, true)
		assert.Equal(checkSelect.Flag_sql.Bool, true)
		assert.True(checkSelect.Flag_sql.Valid)
		assert.Equal(checkSelect.Flag_pqx.Bool, true)
		assert.True(checkSelect.Flag_pqx.Valid)

		assert.Equal(checkSelect.Counter, 42)
		assert.Equal(checkSelect.Counter_sql.Int64, int64(42))
		assert.True(checkSelect.Counter_sql.Valid)
		assert.Equal(checkSelect.Counter64_pqx.Int64, int64(42))
		assert.True(checkSelect.Counter64_pqx.Valid)
		assert.Equal(checkSelect.Counter_pqx.Int, int(42))
		assert.True(checkSelect.Counter_pqx.Valid)
		assert.Equal(checkSelect.Counter32_pqx.Int32, int32(42))
		assert.True(checkSelect.Counter32_pqx.Valid)
		assert.Equal(checkSelect.Counter16_pqx.Int16, int16(42))
		assert.True(checkSelect.Counter16_pqx.Valid)
		assert.Equal(checkSelect.Counter8_pqx.Int8, int8(42))
		assert.True(checkSelect.Counter8_pqx.Valid)

		assert.Equal(checkSelect.Sum, 12.3)
		assert.Equal(checkSelect.Sum_sql.Float64, 12.3)
		assert.True(checkSelect.Sum_sql.Valid)
		assert.Equal(checkSelect.Sum64_pqx.Float64, 12.3)
		assert.True(checkSelect.Sum64_pqx.Valid)
		assert.Equal(checkSelect.Sum32_pqx.Float32, float32(12.3))
		assert.True(checkSelect.Sum32_pqx.Valid)
	}

	// update entity with nil values
	entity.Closed_pq = pq.NullTime{Valid: false}
	entity.Closed_pqx = pqnull.NilTime()
	entity.Msg = "drop"
	entity.Msg_sql = sql.NullString{Valid: false}
	entity.Msg_pqx = pqnull.NilString()
	entity.Flag = false
	entity.Flag_sql = sql.NullBool{Valid: false}
	entity.Flag_pqx = pqnull.NilBool()
	entity.Counter = 21
	entity.Counter_sql = sql.NullInt64{Valid: false}
	entity.Counter64_pqx = pqnull.NilInt64()
	entity.Counter_pqx = pqnull.NilInt()
	entity.Counter32_pqx = pqnull.NilInt32()
	entity.Counter16_pqx = pqnull.NilInt16()
	entity.Counter8_pqx = pqnull.NilInt8()
	entity.Sum = 11.2
	entity.Sum_sql = sql.NullFloat64{Valid: false}
	entity.Sum64_pqx = pqnull.NilFloat64()
	entity.Sum32_pqx = pqnull.NilFloat32()

	err = pqx.Update(&entity)
	assert.Nil(err)

	// select entity -> 1
	{
		checkSelect := TestNullableTypeInsertValidUpdateNilStruct{
			Id: 1,
		}
		find, err := pqx.Select(&checkSelect)
		assert.Nil(err)
		assert.True(find)

		assert.Equal(checkSelect.Id, ID(1))
		assert.Equal(checkSelect.Closed_pq.Time, time.Time{})
		assert.False(checkSelect.Closed_pq.Valid)
		assert.Equal(checkSelect.Closed_pqx.Time, time.Time{})
		assert.False(checkSelect.Closed_pqx.Valid)

		assert.Equal(checkSelect.Msg, "drop")
		assert.Equal(checkSelect.Msg_sql.String, "")
		assert.False(checkSelect.Msg_sql.Valid)
		assert.Equal(checkSelect.Msg_pqx.String, "")
		assert.False(checkSelect.Msg_pqx.Valid)

		assert.Equal(checkSelect.Flag, false)
		assert.Equal(checkSelect.Flag_sql.Bool, false)
		assert.False(checkSelect.Flag_sql.Valid)
		assert.Equal(checkSelect.Flag_pqx.Bool, false)
		assert.False(checkSelect.Flag_pqx.Valid)

		assert.Equal(checkSelect.Counter, 21)
		assert.Equal(checkSelect.Counter_sql.Int64, int64(0))
		assert.False(checkSelect.Counter_sql.Valid)
		assert.Equal(checkSelect.Counter64_pqx.Int64, int64(0))
		assert.False(checkSelect.Counter64_pqx.Valid)
		assert.Equal(checkSelect.Counter_pqx.Int, int(0))
		assert.False(checkSelect.Counter_pqx.Valid)
		assert.Equal(checkSelect.Counter32_pqx.Int32, int32(0))
		assert.False(checkSelect.Counter32_pqx.Valid)
		assert.Equal(checkSelect.Counter16_pqx.Int16, int16(0))
		assert.False(checkSelect.Counter16_pqx.Valid)
		assert.Equal(checkSelect.Counter8_pqx.Int8, int8(0))
		assert.False(checkSelect.Counter8_pqx.Valid)

		assert.Equal(checkSelect.Sum, 11.2)
		assert.Equal(checkSelect.Sum_sql.Float64, 0.0)
		assert.False(checkSelect.Sum_sql.Valid)
		assert.Equal(checkSelect.Sum64_pqx.Float64, 0.0)
		assert.False(checkSelect.Sum64_pqx.Valid)
		assert.Equal(checkSelect.Sum32_pqx.Float32, float32(0.0))
		assert.False(checkSelect.Sum32_pqx.Valid)
	}
}

func TestNullableType_insertNil_updateValid(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)

	type ID int
	type TestNullableTypeInsertNilUpdateValidStruct struct {
		Id            ID        `pqx:"PK AI"`
		Created       time.Time `pqx:"Created"`
		Closed_pq     pq.NullTime
		Closed_pqx    pqnull.Time
		Msg           string
		Msg_sql       sql.NullString
		Msg_pqx       pqnull.String
		Flag          bool
		Flag_sql      sql.NullBool
		Flag_pqx      pqnull.Bool
		Counter       int
		Counter_sql   sql.NullInt64
		Counter64_pqx pqnull.Int64
		Counter_pqx   pqnull.Int
		Counter32_pqx pqnull.Int32
		Counter16_pqx pqnull.Int16
		Counter8_pqx  pqnull.Int8
		Sum           float64
		Sum_sql       sql.NullFloat64
		Sum64_pqx     pqnull.Float64
		Sum32_pqx     pqnull.Float32
	}
	err := pqx.Register(TestNullableTypeInsertNilUpdateValidStruct{})
	assert.Nil(err)

	now := time.Now().UTC().Truncate(time.Millisecond)

	// init entity with nil values
	entity := TestNullableTypeInsertNilUpdateValidStruct{
		Closed_pq:     pq.NullTime{Valid: false},
		Closed_pqx:    pqnull.NilTime(),
		Msg:           "drop",
		Msg_sql:       sql.NullString{Valid: false},
		Msg_pqx:       pqnull.NilString(),
		Flag:          false,
		Flag_sql:      sql.NullBool{Valid: false},
		Flag_pqx:      pqnull.NilBool(),
		Counter:       21,
		Counter_sql:   sql.NullInt64{Valid: false},
		Counter64_pqx: pqnull.NilInt64(),
		Counter_pqx:   pqnull.NilInt(),
		Counter32_pqx: pqnull.NilInt32(),
		Counter16_pqx: pqnull.NilInt16(),
		Counter8_pqx:  pqnull.NilInt8(),
		Sum:           11.2,
		Sum_sql:       sql.NullFloat64{Valid: false},
		Sum64_pqx:     pqnull.NilFloat64(),
		Sum32_pqx:     pqnull.NilFloat32(),
	}
	e := pqx.Insert(&entity)
	assert.Nil(e)
	assert.Equal(entity.Id, ID(1))

	// select entity -> 1
	{
		checkSelect := TestNullableTypeInsertNilUpdateValidStruct{
			Id: 1,
		}
		find, err := pqx.Select(&checkSelect)
		assert.Nil(err)
		assert.True(find)

		assert.Equal(checkSelect.Id, ID(1))
		assert.Equal(checkSelect.Closed_pq.Time, time.Time{})
		assert.False(checkSelect.Closed_pq.Valid)
		assert.Equal(checkSelect.Closed_pqx.Time, time.Time{})
		assert.False(checkSelect.Closed_pqx.Valid)

		assert.Equal(checkSelect.Msg, "drop")
		assert.Equal(checkSelect.Msg_sql.String, "")
		assert.False(checkSelect.Msg_sql.Valid)
		assert.Equal(checkSelect.Msg_pqx.String, "")
		assert.False(checkSelect.Msg_pqx.Valid)

		assert.Equal(checkSelect.Flag, false)
		assert.Equal(checkSelect.Flag_sql.Bool, false)
		assert.False(checkSelect.Flag_sql.Valid)
		assert.Equal(checkSelect.Flag_pqx.Bool, false)
		assert.False(checkSelect.Flag_pqx.Valid)

		assert.Equal(checkSelect.Counter, 21)
		assert.Equal(checkSelect.Counter_sql.Int64, int64(0))
		assert.False(checkSelect.Counter_sql.Valid)
		assert.Equal(checkSelect.Counter64_pqx.Int64, int64(0))
		assert.False(checkSelect.Counter64_pqx.Valid)
		assert.Equal(checkSelect.Counter_pqx.Int, int(0))
		assert.False(checkSelect.Counter_pqx.Valid)
		assert.Equal(checkSelect.Counter32_pqx.Int32, int32(0))
		assert.False(checkSelect.Counter32_pqx.Valid)
		assert.Equal(checkSelect.Counter16_pqx.Int16, int16(0))
		assert.False(checkSelect.Counter16_pqx.Valid)
		assert.Equal(checkSelect.Counter8_pqx.Int8, int8(0))
		assert.False(checkSelect.Counter8_pqx.Valid)

		assert.Equal(checkSelect.Sum, 11.2)
		assert.Equal(checkSelect.Sum_sql.Float64, 0.0)
		assert.False(checkSelect.Sum_sql.Valid)
		assert.Equal(checkSelect.Sum64_pqx.Float64, 0.0)
		assert.False(checkSelect.Sum64_pqx.Valid)
		assert.Equal(checkSelect.Sum32_pqx.Float32, float32(0.0))
		assert.False(checkSelect.Sum32_pqx.Valid)
	}

	// update entity with valid values
	entity.Closed_pq = pq.NullTime{Time: now, Valid: true}
	entity.Closed_pqx = pqnull.ValidTime(now)
	entity.Msg = "blob"
	entity.Msg_sql = sql.NullString{String: "blob", Valid: true}
	entity.Msg_pqx = pqnull.ValidString("blob")
	entity.Flag = true
	entity.Flag_sql = sql.NullBool{Bool: true, Valid: true}
	entity.Flag_pqx = pqnull.ValidBool(true)
	entity.Counter = 42
	entity.Counter_sql = sql.NullInt64{Int64: 42, Valid: true}
	entity.Counter64_pqx = pqnull.ValidInt64(42)
	entity.Counter_pqx = pqnull.ValidInt(42)
	entity.Counter32_pqx = pqnull.ValidInt32(42)
	entity.Counter16_pqx = pqnull.ValidInt16(42)
	entity.Counter8_pqx = pqnull.ValidInt8(42)
	entity.Sum = 12.3
	entity.Sum_sql = sql.NullFloat64{Float64: 12.3, Valid: true}
	entity.Sum64_pqx = pqnull.ValidFloat64(12.3)
	entity.Sum32_pqx = pqnull.ValidFloat32(12.3)

	err = pqx.Update(&entity)
	assert.Nil(err)

	// select entity -> 1
	{
		checkSelect := TestNullableTypeInsertNilUpdateValidStruct{
			Id: 1,
		}
		find, err := pqx.Select(&checkSelect)
		assert.Nil(err)
		assert.True(find)

		assert.Equal(checkSelect.Id, ID(1))
		assert.Equal(checkSelect.Closed_pq.Time, now)
		assert.True(checkSelect.Closed_pq.Valid)
		assert.Equal(checkSelect.Closed_pqx.Time, now)
		assert.True(checkSelect.Closed_pqx.Valid)

		assert.Equal(checkSelect.Msg, "blob")
		assert.Equal(checkSelect.Msg_sql.String, "blob")
		assert.True(checkSelect.Msg_sql.Valid)
		assert.Equal(checkSelect.Msg_pqx.String, "blob")
		assert.True(checkSelect.Msg_pqx.Valid)

		assert.Equal(checkSelect.Flag, true)
		assert.Equal(checkSelect.Flag_sql.Bool, true)
		assert.True(checkSelect.Flag_sql.Valid)
		assert.Equal(checkSelect.Flag_pqx.Bool, true)
		assert.True(checkSelect.Flag_pqx.Valid)

		assert.Equal(checkSelect.Counter, 42)
		assert.Equal(checkSelect.Counter_sql.Int64, int64(42))
		assert.True(checkSelect.Counter_sql.Valid)
		assert.Equal(checkSelect.Counter64_pqx.Int64, int64(42))
		assert.True(checkSelect.Counter64_pqx.Valid)
		assert.Equal(checkSelect.Counter_pqx.Int, int(42))
		assert.True(checkSelect.Counter_pqx.Valid)
		assert.Equal(checkSelect.Counter32_pqx.Int32, int32(42))
		assert.True(checkSelect.Counter32_pqx.Valid)
		assert.Equal(checkSelect.Counter16_pqx.Int16, int16(42))
		assert.True(checkSelect.Counter16_pqx.Valid)
		assert.Equal(checkSelect.Counter8_pqx.Int8, int8(42))
		assert.True(checkSelect.Counter8_pqx.Valid)

		assert.Equal(checkSelect.Sum, 12.3)
		assert.Equal(checkSelect.Sum_sql.Float64, 12.3)
		assert.True(checkSelect.Sum_sql.Valid)
		assert.Equal(checkSelect.Sum64_pqx.Float64, 12.3)
		assert.True(checkSelect.Sum64_pqx.Valid)
		assert.Equal(checkSelect.Sum32_pqx.Float32, float32(12.3))
		assert.True(checkSelect.Sum32_pqx.Valid)
	}
}

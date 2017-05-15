package pqtable_test

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/maprost/assertion"
	"testing"
	"time"

	"github.com/maprost/pqx/pqnull"
	"github.com/maprost/pqx/pqtable"
)

func BenchmarkTinyTable_oneElement(b *testing.B) {
	assert := assertion.New(b)

	type BenchmarkTinyTableOneElementStruct struct {
		Id int64
	}

	for i := 0; i < b.N; i++ {
		_, err := pqtable.New(BenchmarkTinyTableOneElementStruct{})
		assert.Nil(err)
	}
}

func BenchmarkSmallTable_threeElements(b *testing.B) {
	assert := assertion.New(b)

	type BenchmarkSmallTableThreeElements struct {
		Id     int64
		Msg    string
		UserId pqnull.Int64
	}

	for i := 0; i < b.N; i++ {
		_, err := pqtable.New(BenchmarkSmallTableThreeElements{})
		assert.Nil(err)
	}
}

func BenchmarkBigTable_twentyOneElements(b *testing.B) {
	assert := assertion.New(b)

	type ID int64
	type BenchmarkBigTableTwentyOneElementsStruct struct {
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

	for i := 0; i < b.N; i++ {
		_, err := pqtable.New(BenchmarkBigTableTwentyOneElementsStruct{})
		assert.Nil(err)
	}
}

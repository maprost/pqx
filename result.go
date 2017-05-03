package pqx

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/maprost/pqx/pqnull"
	"github.com/maprost/pqx/pqtable"
	"github.com/maprost/pqx/pqtype"
	"time"
)

type Result interface {
	Scan(dest ...interface{}) error
}

func ScanStruct(r Result, output interface{}) error {
	table, err := pqtable.New(output)
	if err != nil {
		return err
	}

	return ScanTable(r, table)
}

func ScanTable(r Result, table *pqtable.Table) (err error) {
	valueList := make([]interface{}, table.Len())
	afterAction := make([]func(), table.Len())

	for index, c := range table.Columns() {
		column := c
		switch column.ReflectType() {
		// =============== string =================
		case pqtype.ReflectString:
			var s string
			valueList[index] = &s
			afterAction[index] = func() {
				column.SetString(s)
			}
		case pqtype.ReflectString_sql:
			var s sql.NullString
			valueList[index] = &s
			afterAction[index] = func() {
				column.SetNullTypeStruct(s)
			}
		case pqtype.ReflectString_pqx:
			var s pqnull.String
			valueList[index] = &s
			afterAction[index] = func() {
				column.SetNullTypeStruct(s)
			}

		// =============== integer =================
		case pqtype.ReflectInt64:
			var i int64
			valueList[index] = &i
			afterAction[index] = func() {
				column.SetInt64(i)
			}
		case pqtype.ReflectInt64_sql:
			var i sql.NullInt64
			valueList[index] = &i
			afterAction[index] = func() {
				column.SetNullTypeStruct(i)
			}
		case pqtype.ReflectInt64_pqx:
			var i pqnull.Int64
			valueList[index] = &i
			afterAction[index] = func() {
				column.SetNullTypeStruct(i)
			}
		case pqtype.ReflectInt_pqx:
			var i pqnull.Int
			valueList[index] = &i
			afterAction[index] = func() {
				column.SetNullTypeStruct(i)
			}
		case pqtype.ReflectInt32_pqx:
			var i pqnull.Int32
			valueList[index] = &i
			afterAction[index] = func() {
				column.SetNullTypeStruct(i)
			}
		case pqtype.ReflectInt16_pqx:
			var i pqnull.Int16
			valueList[index] = &i
			afterAction[index] = func() {
				column.SetNullTypeStruct(i)
			}
		case pqtype.ReflectInt8_pqx:
			var i pqnull.Int8
			valueList[index] = &i
			afterAction[index] = func() {
				column.SetNullTypeStruct(i)
			}

		// =============== uinteger =================
		case pqtype.ReflectUInt64:
			var i uint64
			valueList[index] = &i
			afterAction[index] = func() {
				column.SetUInt64(i)
			}

		// =============== boolean =================
		case pqtype.ReflectBool:
			var b bool
			valueList[index] = &b
			afterAction[index] = func() {
				column.SetBool(b)
			}
		case pqtype.ReflectBool_sql:
			var b sql.NullBool
			valueList[index] = &b
			afterAction[index] = func() {
				column.SetNullTypeStruct(b)
			}
		case pqtype.ReflectBool_pqx:
			var b pqnull.Bool
			valueList[index] = &b
			afterAction[index] = func() {
				column.SetNullTypeStruct(b)
			}

		// =============== float =================
		case pqtype.ReflectFloat64:
			var f float64
			valueList[index] = &f
			afterAction[index] = func() {
				column.SetFloat64(f)
			}
		case pqtype.ReflectFloat64_sql:
			var f sql.NullFloat64
			valueList[index] = &f
			afterAction[index] = func() {
				column.SetNullTypeStruct(f)
			}
		case pqtype.ReflectFloat64_pqx:
			var f pqnull.Float64
			valueList[index] = &f
			afterAction[index] = func() {
				column.SetNullTypeStruct(f)
			}
		case pqtype.ReflectFloat32_pqx:
			var f pqnull.Float32
			valueList[index] = &f
			afterAction[index] = func() {
				column.SetNullTypeStruct(f)
			}

		// =============== time =================
		case pqtype.ReflectTime:
			var t time.Time
			valueList[index] = &t
			afterAction[index] = func() {
				column.SetTime(t)
			}
		case pqtype.ReflectTime_pq:
			var t pq.NullTime
			valueList[index] = &t
			afterAction[index] = func() {
				column.SetNullTypeStruct(t)
			}
		case pqtype.ReflectTime_pqx:
			var t pqnull.Time
			valueList[index] = &t
			afterAction[index] = func() {
				column.SetNullTypeStruct(t)
			}

		default:
			err = errors.New("Not supported column type: " + column.Name() + " (" + column.Type().String() + ").")
			return
		}
	}

	err = r.Scan(valueList...)
	if err != nil {
		return err
	}

	for _, action := range afterAction {
		action()
	}

	return nil
}

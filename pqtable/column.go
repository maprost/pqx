package pqtable

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/maprost/pqx/pqnull"
	"github.com/maprost/pqx/pqtype"
	"reflect"
	"strings"
	"time"
)

type Column struct {
	name             string
	columnType       pqtype.Type
	reflectType      pqtype.ReflectType
	nullable         bool
	primaryKeyTag    bool
	autoIncrementTag bool
	createDateTag    bool
	changeDateTag    bool
	value            reflect.Value
}

func newColumn(table reflect.Value, index int, ctx Context) (c Column, err error) {
	cValue := table.Field(index)
	cType := table.Type().Field(index)

	c = Column{
		name:             strings.ToLower(cType.Name),
		columnType:       pqtype.Text,          // default tag
		reflectType:      pqtype.ReflectString, // default tag
		nullable:         false,
		primaryKeyTag:    false,
		autoIncrementTag: false,
		createDateTag:    false,
		changeDateTag:    false,
		value:            cValue,
	}

	c.setTag(cType, ctx)

	// should be last, need tags
	err = c.setType(cType, ctx)

	return
}

func (c *Column) Name() string {
	return c.name
}

func (c *Column) Type() pqtype.Type {
	return c.columnType
}

func (c *Column) ReflectType() pqtype.ReflectType {
	return c.reflectType
}

func (c *Column) Nullable() bool {
	return c.nullable
}

func (c *Column) PrimaryKeyTag() bool {
	return c.primaryKeyTag
}

func (c *Column) AutoIncrementTag() bool {
	return c.autoIncrementTag
}

func (c *Column) CreateDateTag() bool {
	return c.createDateTag
}

func (c *Column) ChangeDateTag() bool {
	return c.changeDateTag
}

func (c *Column) GetValue() interface{} {
	return c.value.Interface()
}

func (c *Column) SetInt64(v int64) {
	c.value.SetInt(v)
}

func (c *Column) SetUInt64(v uint64) {
	c.value.SetUint(v)
}

func (c *Column) SetFloat64(v float64) {
	c.value.SetFloat(v)
}

func (c *Column) SetBool(v bool) {
	c.value.SetBool(v)
}

func (c *Column) SetString(v string) {
	c.value.SetString(v)
}

func (c *Column) SetTime(v time.Time) {
	c.value.Set(reflect.ValueOf(v))
}

func (c *Column) SetNullTypeStruct(v interface{}) {
	c.value.Set(reflect.ValueOf(v))
}

func (c *Column) setTag(cType reflect.StructField, ctx Context) {
	tagString := cType.Tag.Get("pqx")

	c.primaryKeyTag = strings.Contains(tagString, "PK")
	c.autoIncrementTag = strings.Contains(tagString, "AI")
	c.createDateTag = strings.Contains(tagString, "CreateDate")
	c.changeDateTag = strings.Contains(tagString, "ChangeDate")
}

func (c *Column) setType(cType reflect.StructField, ctx Context) (err error) {
	switch cType.Type.Kind() {
	case reflect.String:
		c.columnType = pqtype.Text
		c.reflectType = pqtype.ReflectString

	case reflect.Int8, reflect.Int16:
		if c.AutoIncrementTag() {
			c.columnType = pqtype.SmallSerial
		} else {
			c.columnType = pqtype.SmallInt
		}
		c.reflectType = pqtype.ReflectInt64

	case reflect.Uint8:
		if c.AutoIncrementTag() {
			c.columnType = pqtype.SmallSerial
		} else {
			c.columnType = pqtype.SmallInt
		}
		c.reflectType = pqtype.ReflectUInt64

	case reflect.Int, reflect.Int32:
		if c.AutoIncrementTag() {
			c.columnType = pqtype.Serial
		} else {
			c.columnType = pqtype.Integer
		}
		c.reflectType = pqtype.ReflectInt64

	case reflect.Uint16:
		if c.AutoIncrementTag() {
			c.columnType = pqtype.Serial
		} else {
			c.columnType = pqtype.Integer
		}
		c.reflectType = pqtype.ReflectUInt64

	case reflect.Int64:
		if c.AutoIncrementTag() {
			c.columnType = pqtype.BigSerial
		} else {
			c.columnType = pqtype.BigInt
		}
		c.reflectType = pqtype.ReflectInt64

	case reflect.Uint32:
		if c.AutoIncrementTag() {
			c.columnType = pqtype.BigSerial
		} else {
			c.columnType = pqtype.BigInt
		}
		c.reflectType = pqtype.ReflectUInt64

	case reflect.Uint64:
		c.columnType = pqtype.Numeric
		c.reflectType = pqtype.ReflectUInt64

	case reflect.Bool:
		c.columnType = pqtype.Bool
		c.reflectType = pqtype.ReflectBool

	case reflect.Float32:
		c.columnType = pqtype.Real
		c.reflectType = pqtype.ReflectFloat64

	case reflect.Float64:
		c.columnType = pqtype.DoublePrecision
		c.reflectType = pqtype.ReflectFloat64

	case reflect.Struct:
		switch c.GetValue().(type) {
		case time.Time:
			c.columnType = pqtype.Time
			c.reflectType = pqtype.ReflectTime

		case pq.NullTime:
			c.columnType = pqtype.Time
			c.reflectType = pqtype.ReflectTime_pq
			c.nullable = true

		case sql.NullBool:
			c.columnType = pqtype.Bool
			c.reflectType = pqtype.ReflectBool_sql
			c.nullable = true

		case sql.NullInt64:
			c.columnType = pqtype.BigInt
			c.reflectType = pqtype.ReflectInt64_sql
			c.nullable = true

		case sql.NullString:
			c.columnType = pqtype.Text
			c.reflectType = pqtype.ReflectString_sql
			c.nullable = true

		case sql.NullFloat64:
			c.columnType = pqtype.DoublePrecision
			c.reflectType = pqtype.ReflectFloat64_sql
			c.nullable = true

		case pqnull.String:
			c.columnType = pqtype.Text
			c.reflectType = pqtype.ReflectString_pqx
			c.nullable = true

		case pqnull.Int64:
			c.columnType = pqtype.BigInt
			c.reflectType = pqtype.ReflectInt64_pqx
			c.nullable = true

		case pqnull.Int:
			c.columnType = pqtype.BigInt
			c.reflectType = pqtype.ReflectInt_pqx
			c.nullable = true

		case pqnull.Int32:
			c.columnType = pqtype.Integer
			c.reflectType = pqtype.ReflectInt32_pqx
			c.nullable = true

		case pqnull.Int16:
			c.columnType = pqtype.BigInt
			c.reflectType = pqtype.ReflectInt16_pqx
			c.nullable = true

		case pqnull.Int8:
			c.columnType = pqtype.SmallInt
			c.reflectType = pqtype.ReflectInt8_pqx
			c.nullable = true

		case pqnull.Time:
			c.columnType = pqtype.Time
			c.reflectType = pqtype.ReflectTime_pqx
			c.nullable = true

		case pqnull.Bool:
			c.columnType = pqtype.Bool
			c.reflectType = pqtype.ReflectBool_pqx
			c.nullable = true

		case pqnull.Float64:
			c.columnType = pqtype.DoublePrecision
			c.reflectType = pqtype.ReflectFloat64_pqx
			c.nullable = true

		case pqnull.Float32:
			c.columnType = pqtype.Real
			c.reflectType = pqtype.ReflectFloat32_pqx
			c.nullable = true

		default:
			if ctx.IgnoreUnknownColumnType == false {
				err = errors.New("Not supported struct type: " + c.Name() + " (" + cType.Type.String() + ").")
			}
		}

	default:
		if ctx.IgnoreUnknownColumnType == false {
			err = errors.New("Not supported column type: " + c.Name() + " (" + cType.Type.String() + ").")
		}
	}

	return
}

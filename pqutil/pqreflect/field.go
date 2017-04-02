package pqreflect

import (
	"database/sql"
	"reflect"
	"strings"
	"time"
)

type Field struct {
	field       reflect.Value
	structField reflect.StructField
}

func NewField(elem reflect.Value, index int) Field {
	return Field{field: elem.Field(index), structField: elem.Type().Field(index)}
}

func (f Field) Name() string {
	return strings.ToLower(f.structField.Name)
}

func (f Field) Type() string {
	return f.structField.Type.Name()
}

func (f Field) Tag(key string) string {
	return f.structField.Tag.Get(key)
}

func (f Field) GetInt() int64 {
	return f.field.Int()
}

func (f Field) GetUint() uint64 {
	return f.field.Uint()
}

func (f Field) GetString() string {
	return f.field.String()
}

func (f Field) GetBool() bool {
	return f.field.Bool()
}

func (f Field) GetFloat() float64 {
	return f.field.Float()
}

func (f Field) GetTime() time.Time {
	return f.TypeInterface().(time.Time)
}

func (f Field) GetNullBool() sql.NullBool {
	return f.TypeInterface().(sql.NullBool)
}

func (f Field) GetNullString() sql.NullString {
	return f.TypeInterface().(sql.NullString)
}

func (f Field) GetNullInt64() sql.NullInt64 {
	return f.TypeInterface().(sql.NullInt64)
}

func (f Field) GetNullFloat64() sql.NullFloat64 {
	return f.TypeInterface().(sql.NullFloat64)
}

func (f Field) GetValue() interface{} {
	switch f.Kind() {
	case reflect.String:
		return f.GetString()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f.GetInt()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return f.GetUint()
	case reflect.Bool:
		return f.GetBool()
	case reflect.Float32, reflect.Float64:
		return f.GetFloat()
	default:
		switch f.TypeInterface().(type) {
		case time.Time:
			return f.GetTime()
		case sql.NullBool:
			return f.GetNullBool()
		case sql.NullString:
			return f.GetNullString()
		case sql.NullInt64:
			return f.GetNullInt64()
		case sql.NullFloat64:
			return f.GetNullFloat64()
		}
		return nil
	}
}

func (f *Field) SetInt(i int64) {
	f.field.SetInt(i)
}

func (f *Field) SetUint(i uint64) {
	f.field.SetUint(i)
}

func (f *Field) SetString(s string) {
	f.field.SetString(s)
}

func (f *Field) SetBool(b bool) {
	f.field.SetBool(b)
}

func (f *Field) SetFloat(fl float64) {
	f.field.SetFloat(fl)
}

func (f *Field) SetTime(t time.Time) {
	f.field.Set(reflect.ValueOf(t))
}

func (f *Field) SetNullBool(b sql.NullBool) {
	f.field.Set(reflect.ValueOf(b))
}

func (f *Field) SetNullString(s sql.NullString) {
	f.field.Set(reflect.ValueOf(s))
}

func (f *Field) SetNullInt64(i sql.NullInt64) {
	f.field.Set(reflect.ValueOf(i))
}

func (f *Field) SetNullFloat64(fl sql.NullFloat64) {
	f.field.Set(reflect.ValueOf(fl))
}

func (f Field) TypeInterface() interface{} {
	return f.field.Interface()
}

func (f Field) Kind() reflect.Kind {
	return f.field.Kind()
}

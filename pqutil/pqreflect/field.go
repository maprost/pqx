package pqreflect

import (
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

func (f Field) GetValue() interface{} {
	switch f.TypeInterface().(type) {
	case string:
		return f.GetString()
	case int, int8, int16, int32, int64:
		return f.GetInt()
	case bool:
		return f.GetBool()
	case float32, float64:
		return f.GetFloat()
	case time.Time:
		return f.GetTime()
	default:
		return nil
	}
}

func (f *Field) SetInt(i int64) {
	f.field.SetInt(i)
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

// that's only a prototype, should do in another way
func (f *Field) SetValue(value interface{}) {
	switch reflect.ValueOf(value).Interface().(type) {
	case string:
		if v, ok := value.(string); ok {
			f.SetString(v)
		}
	case int, int8, int16, int32, int64:
		if v, ok := value.(int64); ok {
			f.SetInt(v)
		}
	case bool:
		if v, ok := value.(bool); ok {
			f.SetBool(v)
		}
	case float32, float64:
		if v, ok := value.(float64); ok {
			f.SetFloat(v)
		}
	}
}

func (f Field) TypeInterface() interface{} {
	return f.field.Interface()
}

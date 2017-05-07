package pqnull_test

import (
	"fmt"
	"github.com/maprost/assertion"
	"math"
	"strconv"
	"testing"

	"github.com/maprost/pqx/pqnull"
)

func TestInt64(t *testing.T) {
	assert := assertion.New(t)

	i := pqnull.NilInt64()
	assert.Equal(i.Int64, int64(0))
	assert.False(i.Valid)

	var p *int64 = i.Ptr()
	assert.Nil(p)

	v, err := i.Value()
	assert.Nil(err)
	assert.Nil(v)

	err = i.Scan(math.MaxInt64)
	assert.Nil(err)
	assert.Equal(i.Int64, int64(math.MaxInt64))
	assert.True(i.Valid)

	v, err = i.Value()
	assert.Nil(err)
	assert.Equal(v, int64(math.MaxInt64))

	p = i.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, int64(math.MaxInt64))

	i = pqnull.PtrInt64(p)
	assert.Equal(i.Int64, int64(math.MaxInt64))
	assert.True(i.Valid)

	i = pqnull.ValidInt64(42)
	assert.Equal(i.Int64, int64(42))
	assert.True(i.Valid)

	i = pqnull.PtrInt64(nil)
	assert.Equal(i.Int64, int64(0))
	assert.False(i.Valid)
}

func TestInt(t *testing.T) {
	assert := assertion.New(t)

	i := pqnull.NilInt()
	assert.Equal(i.Int, int(0))
	assert.False(i.Valid)

	var p *int = i.Ptr()
	assert.Nil(p)

	v, err := i.Value()
	assert.Nil(err)
	assert.Nil(v)

	maxInt := math.MaxInt64
	if strconv.IntSize != 64 {
		fmt.Println("Use max int 32 instead of 64.")
		maxInt = math.MaxInt32
	}

	err = i.Scan(maxInt)
	assert.Nil(err)
	assert.Equal(i.Int, int(maxInt))
	assert.True(i.Valid)

	v, err = i.Value()
	assert.Nil(err)
	assert.Equal(v, int64(maxInt))

	p = i.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, int(maxInt))

	i = pqnull.PtrInt(p)
	assert.Equal(i.Int, int(maxInt))
	assert.True(i.Valid)

	i = pqnull.ValidInt(42)
	assert.Equal(i.Int, int(42))
	assert.True(i.Valid)

	i = pqnull.PtrInt(nil)
	assert.Equal(i.Int, int(0))
	assert.False(i.Valid)
}

func TestInt32(t *testing.T) {
	assert := assertion.New(t)

	i := pqnull.NilInt32()
	assert.Equal(i.Int32, int32(0))
	assert.False(i.Valid)

	var p *int32 = i.Ptr()
	assert.Nil(p)

	v, err := i.Value()
	assert.Nil(err)
	assert.Nil(v)

	err = i.Scan(math.MaxInt32)
	assert.Nil(err)
	assert.Equal(i.Int32, int32(math.MaxInt32))
	assert.True(i.Valid)

	v, err = i.Value()
	assert.Nil(err)
	assert.Equal(v, int64(math.MaxInt32))

	p = i.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, int32(math.MaxInt32))

	i = pqnull.PtrInt32(p)
	assert.Equal(i.Int32, int32(math.MaxInt32))
	assert.True(i.Valid)

	i = pqnull.ValidInt32(42)
	assert.Equal(i.Int32, int32(42))
	assert.True(i.Valid)

	i = pqnull.PtrInt32(nil)
	assert.Equal(i.Int32, int32(0))
	assert.False(i.Valid)
}

func TestInt16(t *testing.T) {
	assert := assertion.New(t)

	i := pqnull.NilInt16()
	assert.Equal(i.Int16, int16(0))
	assert.False(i.Valid)

	var p *int16 = i.Ptr()
	assert.Nil(p)

	v, err := i.Value()
	assert.Nil(err)
	assert.Nil(v)

	err = i.Scan(math.MaxInt16)
	assert.Nil(err)
	assert.Equal(i.Int16, int16(math.MaxInt16))
	assert.True(i.Valid)

	v, err = i.Value()
	assert.Nil(err)
	assert.Equal(v, int64(math.MaxInt16))

	p = i.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, int16(math.MaxInt16))

	i = pqnull.PtrInt16(p)
	assert.Equal(i.Int16, int16(math.MaxInt16))
	assert.True(i.Valid)

	i = pqnull.ValidInt16(42)
	assert.Equal(i.Int16, int16(42))
	assert.True(i.Valid)

	i = pqnull.PtrInt16(nil)
	assert.Equal(i.Int16, int16(0))
	assert.False(i.Valid)
}

func TestInt8(t *testing.T) {
	assert := assertion.New(t)

	i := pqnull.NilInt8()
	assert.Equal(i.Int8, int8(0))
	assert.False(i.Valid)

	var p *int8 = i.Ptr()
	assert.Nil(p)

	v, err := i.Value()
	assert.Nil(err)
	assert.Nil(v)

	err = i.Scan(math.MaxInt8)
	assert.Nil(err)
	assert.Equal(i.Int8, int8(math.MaxInt8))
	assert.True(i.Valid)

	v, err = i.Value()
	assert.Nil(err)
	assert.Equal(v, int64(math.MaxInt8))

	p = i.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, int8(math.MaxInt8))

	i = pqnull.PtrInt8(p)
	assert.Equal(i.Int8, int8(math.MaxInt8))
	assert.True(i.Valid)

	i = pqnull.ValidInt8(42)
	assert.Equal(i.Int8, int8(42))
	assert.True(i.Valid)

	i = pqnull.PtrInt8(nil)
	assert.Equal(i.Int8, int8(0))
	assert.False(i.Valid)
}

//func TestInt_validInt64(t *testing.T) {
//	assert := assertion.New(t)
//
//	i := pqnull.ValidInt64(42)
//	assert.Equal(i.Int64, int64(42))
//	assert.True(i.Valid)
//
//	p := i.Ptr()
//	assert.NotNil(p)
//	assert.Equal(*p, int64(42))
//
//	v, err := i.Value()
//	assert.Nil(err)
//	assert.Equal(v, int64(42))
//}
//
//func TestInt_ptr(t *testing.T) {
//	assert := assertion.New(t)
//	var in int64 = 42
//
//	i := pqnull.PtrInt64(&in)
//	assert.Equal(i.Int64, int64(42))
//	assert.True(i.Valid)
//
//	p := i.Ptr()
//	assert.NotNil(p)
//	assert.Equal(*p, int64(42))
//
//	v, err := i.Value()
//	assert.Nil(err)
//	assert.Equal(v, int64(42))
//}
//
//func TestInt_ptr_nil(t *testing.T) {
//	assert := assertion.New(t)
//
//	i := pqnull.PtrInt64(nil)
//	assert.Equal(i.Int64, int64(0))
//	assert.False(i.Valid)
//
//	p := i.Ptr()
//	if p != nil {
//		assert.Fail("Should be nil!")
//	}
//
//	v, err := i.Value()
//	assert.Nil(err)
//	assert.Nil(v)
//}
//
//func TestInt_init(t *testing.T) {
//	assert := assertion.New(t)
//
//	var i pqnull.Int64
//	assert.False(i.Valid)
//	assert.Equal(i.Int64, int64(0))
//
//	v, err := i.Value()
//	assert.Nil(err)
//	assert.Nil(v)
//
//	assert.Nil(i.Ptr8())
//	assert.Nil(i.Ptr16())
//	assert.Nil(i.Ptr32())
//	assert.Nil(i.Ptr())
//	assert.Nil(i.Ptr())
//}
//
//func TestInt_scan_value(t *testing.T) {
//	assert := assertion.New(t)
//
//	var i pqnull.Int64
//	err := i.Scan(42)
//	assert.Nil(err)
//
//	assert.True(i.Valid)
//	assert.Equal(i.Int64, int64(42))
//
//	v, err := i.Value()
//	assert.Nil(err)
//	assert.Equal(v, int64(42))
//}
//
//func TestInt_scanInt8_value(t *testing.T) {
//	assert := assertion.New(t)
//
//	var i pqnull.Int64
//	err := i.Scan(int8(42))
//	assert.Nil(err)
//
//	assert.True(i.Valid)
//	assert.Equal(i.Int64, int64(42))
//
//	v, err := i.Value()
//	assert.Nil(err)
//	assert.Equal(v, int64(42))
//}
//
//func TestInt_scan_nil(t *testing.T) {
//	assert := assertion.New(t)
//
//	var i pqnull.Int64
//	err := i.Scan(nil)
//	assert.Nil(err)
//
//	assert.False(i.Valid)
//	assert.Equal(i.Int64, int64(0))
//
//	v, err := i.Value()
//	assert.Nil(err)
//	assert.Nil(v)
//
//	assert.Nil(i.Ptr8())
//	assert.Nil(i.Ptr16())
//	assert.Nil(i.Ptr32())
//	assert.Nil(i.Ptr())
//	assert.Nil(i.Ptr())
//}
//
//func TestInt_kind(t *testing.T) {
//	assert := assertion.New(t)
//	var i pqnull.Int64
//
//	typ := reflect.TypeOf(i)
//	assert.Equal(typ.Kind(), reflect.Struct)
//
//	val := reflect.ValueOf(i)
//	switch val.Interface().(type) {
//	case sql.NullInt64:
//		assert.Fail("Wrong type.")
//	case pqnull.Int64:
//		// correct
//	default:
//		assert.Fail("Type not found")
//	}
//}
//
//func TestInt8(t *testing.T) {
//	assert := assertion.New(t)
//
//	var maxValue8 int8 = math.MaxInt8
//	var maxValue16 int16 = math.MaxInt8
//	var maxValue int = math.MaxInt8
//	var maxValue32 int32 = math.MaxInt8
//	var maxValue64 int64 = math.MaxInt8
//
//	i := pqnull.ValidInt8(maxValue8)
//	assert.Equal(i.Int64, int64(maxValue8))
//	assert.True(i.Valid)
//	assert.Equal(i.Ptr8(), &maxValue8)
//	assert.Equal(i.Ptr16(), &maxValue16)
//	assert.Equal(i.Ptr32(), &maxValue32)
//	assert.Equal(i.Ptr(), &maxValue)
//	assert.Equal(i.Ptr(), &maxValue64)
//
//	j := pqnull.PtrInt8(i.Ptr8())
//	assert.Equal(j.Int64, int64(maxValue8))
//	assert.True(j.Valid)
//	assert.Equal(j.Ptr8(), &maxValue8)
//	assert.Equal(j.Ptr16(), &maxValue16)
//	assert.Equal(j.Ptr32(), &maxValue32)
//	assert.Equal(j.Ptr(), &maxValue)
//	assert.Equal(j.Ptr(), &maxValue64)
//}
//
//func TestInt16(t *testing.T) {
//	assert := assertion.New(t)
//
//	var maxValue8 int8 = -1 // can't convert int16 to int8
//	var maxValue16 int16 = math.MaxInt16
//	var maxValue int = math.MaxInt16
//	var maxValue32 int32 = math.MaxInt16
//	var maxValue64 int64 = math.MaxInt16
//
//	i := pqnull.ValidInt16(maxValue16)
//	assert.Equal(i.Int64, int64(maxValue16))
//	assert.True(i.Valid)
//	assert.Equal(i.Ptr8(), &maxValue8)
//	assert.Equal(i.Ptr16(), &maxValue16)
//	assert.Equal(i.Ptr32(), &maxValue32)
//	assert.Equal(i.Ptr(), &maxValue)
//	assert.Equal(i.Ptr(), &maxValue64)
//
//	j := pqnull.PtrInt16(i.Ptr16())
//	assert.Equal(j.Int64, int64(maxValue16))
//	assert.True(j.Valid)
//	assert.Equal(j.Ptr8(), &maxValue8)
//	assert.Equal(j.Ptr16(), &maxValue16)
//	assert.Equal(j.Ptr32(), &maxValue32)
//	assert.Equal(j.Ptr(), &maxValue)
//	assert.Equal(j.Ptr(), &maxValue64)
//}
//
//func TestInt32(t *testing.T) {
//	assert := assertion.New(t)
//
//	var maxValue8 int8 = -1   // can't convert int32 to int8
//	var maxValue16 int16 = -1 // can't convert int32 to int16
//	var maxValue int = math.MaxInt32
//	var maxValue32 int32 = math.MaxInt32
//	var maxValue64 int64 = math.MaxInt32
//
//	i := pqnull.ValidInt32(maxValue32)
//	assert.Equal(i.Int64, int64(maxValue32))
//	assert.True(i.Valid)
//	assert.Equal(i.Ptr8(), &maxValue8)
//	assert.Equal(i.Ptr16(), &maxValue16)
//	assert.Equal(i.Ptr32(), &maxValue32)
//	assert.Equal(i.Ptr(), &maxValue)
//	assert.Equal(i.Ptr(), &maxValue64)
//
//	j := pqnull.PtrInt32(i.Ptr32())
//	assert.Equal(j.Int64, int64(maxValue32))
//	assert.True(j.Valid)
//	assert.Equal(j.Ptr8(), &maxValue8)
//	assert.Equal(j.Ptr16(), &maxValue16)
//	assert.Equal(j.Ptr32(), &maxValue32)
//	assert.Equal(j.Ptr(), &maxValue)
//	assert.Equal(j.Ptr(), &maxValue64)
//}
//
//func TestInt(t *testing.T) {
//	assert := assertion.New(t)
//
//	var maxValue8 int8 = -1   // can't convert int to int8
//	var maxValue16 int16 = -1 // can't convert int to int16
//	var maxValue int = math.MaxInt32
//	var maxValue32 int32 = math.MaxInt32
//	var maxValue64 int64 = math.MaxInt32
//
//	i := pqnull.ValidInt(maxValue)
//	assert.Equal(i.Int64, int64(maxValue))
//	assert.True(i.Valid)
//	assert.Equal(i.Ptr8(), &maxValue8)
//	assert.Equal(i.Ptr16(), &maxValue16)
//	assert.Equal(i.Ptr32(), &maxValue32)
//	assert.Equal(i.Ptr(), &maxValue)
//	assert.Equal(i.Ptr(), &maxValue64)
//
//	j := pqnull.PtrInt(i.Ptr())
//	assert.Equal(j.Int64, int64(maxValue))
//	assert.True(j.Valid)
//	assert.Equal(j.Ptr8(), &maxValue8)
//	assert.Equal(j.Ptr16(), &maxValue16)
//	assert.Equal(j.Ptr32(), &maxValue32)
//	assert.Equal(j.Ptr(), &maxValue)
//	assert.Equal(j.Ptr(), &maxValue64)
//}
//
//func TestInt64(t *testing.T) {
//	assert := assertion.New(t)
//
//	var maxValue8 int8 = -1   // can't convert int64 to int8
//	var maxValue16 int16 = -1 // can't convert int64 to int16
//	var maxValue32 int32 = -1 // can't convert int64 to int32
//	var maxValue int = math.MaxInt64
//	var maxValue64 int64 = math.MaxInt64
//
//	i := pqnull.ValidInt64(maxValue64)
//	assert.Equal(i.Int64, maxValue64)
//	assert.True(i.Valid)
//	assert.Equal(i.Ptr8(), &maxValue8)
//	assert.Equal(i.Ptr16(), &maxValue16)
//	assert.Equal(i.Ptr32(), &maxValue32)
//	assert.Equal(i.Ptr(), &maxValue)
//	assert.Equal(i.Ptr(), &maxValue64)
//
//	j := pqnull.PtrInt64(i.Ptr())
//	assert.Equal(j.Int64, maxValue64)
//	assert.True(j.Valid)
//	assert.Equal(j.Ptr8(), &maxValue8)
//	assert.Equal(j.Ptr16(), &maxValue16)
//	assert.Equal(j.Ptr32(), &maxValue32)
//	assert.Equal(j.Ptr(), &maxValue)
//	assert.Equal(j.Ptr(), &maxValue64)
//}
//
//func TestInt_nilPtr(t *testing.T) {
//	assert := assertion.New(t)
//
//	i8 := pqnull.PtrInt8(nil)
//	assert.Nil(i8.Ptr8())
//
//	i16 := pqnull.PtrInt16(nil)
//	assert.Nil(i16.Ptr16())
//
//	i32 := pqnull.PtrInt32(nil)
//	assert.Nil(i32.Ptr32())
//
//	i := pqnull.PtrInt(nil)
//	assert.Nil(i.Ptr())
//
//	i64 := pqnull.PtrInt64(nil)
//	assert.Nil(i64.Ptr())
//}

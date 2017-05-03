package pqnull_test

import (
	"database/sql"
	"github.com/maprost/assertion"
	"github.com/maprost/pqx/pqnull"
	"reflect"
	"testing"
)

func TestBool_nullBool(t *testing.T) {
	assert := assertion.New(t)

	b := pqnull.NilBool()
	assert.Equal(b.Bool, false)
	assert.False(b.Valid)

	var p *bool = b.Ptr()
	if p != nil {
		assert.Fail("Should be nil!")
	}

	v, err := b.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestBool_validBool(t *testing.T) {
	assert := assertion.New(t)

	b := pqnull.ValidBool(true)
	assert.Equal(b.Bool, true)
	assert.True(b.Valid)

	p := b.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, true)

	v, err := b.Value()
	assert.Nil(err)
	assert.Equal(v, true)
}

func TestBool_ptr(t *testing.T) {
	assert := assertion.New(t)
	var bo bool = true

	b := pqnull.PtrBool(&bo)
	assert.Equal(b.Bool, true)
	assert.True(b.Valid)

	p := b.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, true)

	v, err := b.Value()
	assert.Nil(err)
	assert.Equal(v, true)
}

func TestBool_ptr_nil(t *testing.T) {
	assert := assertion.New(t)

	b := pqnull.PtrBool(nil)
	assert.Equal(b.Bool, false)
	assert.False(b.Valid)

	p := b.Ptr()
	if p != nil {
		assert.Fail("Should be nil!")
	}

	v, err := b.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestBool_init(t *testing.T) {
	assert := assertion.New(t)

	var b pqnull.Bool
	assert.False(b.Valid)
	assert.Equal(b.Bool, false)

	v, err := b.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestBool_scan_value(t *testing.T) {
	assert := assertion.New(t)

	var b pqnull.Bool
	err := b.Scan(true)
	assert.Nil(err)

	assert.True(b.Valid)
	assert.Equal(b.Bool, true)

	v, err := b.Value()
	assert.Nil(err)
	assert.Equal(v, true)
}

func TestBool_scan_nil(t *testing.T) {
	assert := assertion.New(t)

	var b pqnull.Bool
	err := b.Scan(nil)
	assert.Nil(err)

	assert.False(b.Valid)
	assert.Equal(b.Bool, false)

	v, err := b.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestBool_kind(t *testing.T) {
	assert := assertion.New(t)
	var b pqnull.Bool

	typ := reflect.TypeOf(b)
	assert.Equal(typ.Kind(), reflect.Struct)

	val := reflect.ValueOf(b)
	switch val.Interface().(type) {
	case sql.NullBool:
		assert.Fail("Wrong type.")
	case pqnull.Bool:
		// correct
	default:
		assert.Fail("Type not found")
	}
}

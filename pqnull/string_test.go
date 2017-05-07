package pqnull_test

import (
	"database/sql"
	"github.com/maprost/assertion"
	"reflect"
	"testing"

	"github.com/maprost/pqx/pqnull"
)

func TestString_nullString(t *testing.T) {
	assert := assertion.New(t)

	s := pqnull.NilString()
	assert.Equal(s.String, "")
	assert.False(s.Valid)

	var p *string = s.Ptr()
	if p != nil {
		assert.Fail("Should be nil!")
	}

	v, err := s.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestString_validString(t *testing.T) {
	assert := assertion.New(t)

	s := pqnull.ValidString("blob")
	assert.Equal(s.String, "blob")
	assert.True(s.Valid)

	p := s.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, "blob")

	v, err := s.Value()
	assert.Nil(err)
	assert.Equal(v, "blob")
}

func TestString_ptr(t *testing.T) {
	assert := assertion.New(t)
	var str string = "blob"

	s := pqnull.PtrString(&str)
	assert.Equal(s.String, "blob")
	assert.True(s.Valid)

	p := s.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, "blob")

	v, err := s.Value()
	assert.Nil(err)
	assert.Equal(v, "blob")
}

func TestString_ptr_nil(t *testing.T) {
	assert := assertion.New(t)

	s := pqnull.PtrString(nil)
	assert.Equal(s.String, "")
	assert.False(s.Valid)

	p := s.Ptr()
	if p != nil {
		assert.Fail("Should be nil!")
	}

	v, err := s.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestString_init(t *testing.T) {
	assert := assertion.New(t)

	var s pqnull.String
	assert.False(s.Valid)
	assert.Equal(s.String, "")

	v, err := s.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestString_scan_value(t *testing.T) {
	assert := assertion.New(t)

	var s pqnull.String
	err := s.Scan("blob")
	assert.Nil(err)

	assert.True(s.Valid)
	assert.Equal(s.String, "blob")

	v, err := s.Value()
	assert.Nil(err)
	assert.Equal(v, "blob")
}

func TestString_scan_nil(t *testing.T) {
	assert := assertion.New(t)

	var s pqnull.String
	err := s.Scan(nil)
	assert.Nil(err)

	assert.False(s.Valid)
	assert.Equal(s.String, "")

	v, err := s.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestString_kind(t *testing.T) {
	assert := assertion.New(t)
	var s pqnull.String

	typ := reflect.TypeOf(s)
	assert.Equal(typ.Kind(), reflect.Struct)

	val := reflect.ValueOf(s)
	switch val.Interface().(type) {
	case sql.NullString:
		assert.Fail("Wrong type.")
	case pqnull.String:
		// correct
	default:
		assert.Fail("Type not found")
	}
}

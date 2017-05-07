package pqnull_test

import (
	"github.com/lib/pq"
	"github.com/maprost/assertion"
	"reflect"
	"testing"
	"time"

	"github.com/maprost/pqx/pqnull"
)

var defaultTime time.Time = time.Now()
var emptyTime = time.Time{}

func TestTime_NullTime(t *testing.T) {
	assert := assertion.New(t)

	tm := pqnull.NilTime()
	assert.Equal(tm.Time, emptyTime)
	assert.False(tm.Valid)

	var p *time.Time = tm.Ptr()
	if p != nil {
		assert.Fail("Should be nil!")
	}

	v, err := tm.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestTime_validTime(t *testing.T) {
	assert := assertion.New(t)

	tm := pqnull.ValidTime(defaultTime)
	assert.Equal(tm.Time, defaultTime)
	assert.True(tm.Valid)

	p := tm.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, defaultTime)

	v, err := tm.Value()
	assert.Nil(err)
	assert.Equal(v, defaultTime)
}

func TestTime_ptr(t *testing.T) {
	assert := assertion.New(t)
	var in time.Time = defaultTime

	tm := pqnull.PtrTime(&in)
	assert.Equal(tm.Time, defaultTime)
	assert.True(tm.Valid)

	p := tm.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, defaultTime)

	v, err := tm.Value()
	assert.Nil(err)
	assert.Equal(v, defaultTime)
}

func TestTime_ptr_nil(t *testing.T) {
	assert := assertion.New(t)

	tm := pqnull.PtrTime(nil)
	assert.Equal(tm.Time, emptyTime)
	assert.False(tm.Valid)

	p := tm.Ptr()
	if p != nil {
		assert.Fail("Should be nil!")
	}

	v, err := tm.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestTime_init(t *testing.T) {
	assert := assertion.New(t)

	var tm pqnull.Time
	assert.False(tm.Valid)
	assert.Equal(tm.Time, emptyTime)

	v, err := tm.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestTime_scan_value(t *testing.T) {
	assert := assertion.New(t)

	var tm pqnull.Time
	err := tm.Scan(defaultTime)
	assert.Nil(err)

	assert.True(tm.Valid)
	assert.Equal(tm.Time, defaultTime)

	v, err := tm.Value()
	assert.Nil(err)
	assert.Equal(v, defaultTime)
}

func TestTime_scan_nil(t *testing.T) {
	assert := assertion.New(t)

	var tm pqnull.Time
	err := tm.Scan(nil)
	assert.Nil(err)

	assert.False(tm.Valid)
	assert.Equal(tm.Time, emptyTime)

	v, err := tm.Value()
	assert.Nil(err)
	assert.Nil(v)
}

func TestTime_kind(t *testing.T) {
	assert := assertion.New(t)
	var tm pqnull.Time

	typ := reflect.TypeOf(tm)
	assert.Equal(typ.Kind(), reflect.Struct)

	val := reflect.ValueOf(tm)
	switch val.Interface().(type) {
	case pq.NullTime:
		assert.Fail("Wrong type.")
	case pqnull.Time:
	// correct
	default:
		assert.Fail("Type not found")
	}
}

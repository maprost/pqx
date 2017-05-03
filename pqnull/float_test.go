package pqnull_test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/pqx/pqnull"
	"math"
	"testing"
)

func TestFloat64(t *testing.T) {
	assert := assertion.New(t)

	i := pqnull.NilFloat64()
	assert.Equal(i.Float64, float64(0))
	assert.False(i.Valid)

	var p *float64 = i.Ptr()
	assert.Nil(p)

	v, err := i.Value()
	assert.Nil(err)
	assert.Nil(v)

	err = i.Scan(math.MaxFloat64)
	assert.Nil(err)
	assert.Equal(i.Float64, math.MaxFloat64)
	assert.True(i.Valid)

	v, err = i.Value()
	assert.Nil(err)
	assert.Equal(v, math.MaxFloat64)

	p = i.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, math.MaxFloat64)

	i = pqnull.PtrFloat64(p)
	assert.Equal(i.Float64, math.MaxFloat64)
	assert.True(i.Valid)

	i = pqnull.ValidFloat64(42.1)
	assert.Equal(i.Float64, float64(42.1))
	assert.True(i.Valid)

	i = pqnull.PtrFloat64(nil)
	assert.Equal(i.Float64, float64(0))
	assert.False(i.Valid)
}

func TestFloat32(t *testing.T) {
	assert := assertion.New(t)

	i := pqnull.NilFloat32()
	assert.Equal(i.Float32, float32(0))
	assert.False(i.Valid)

	var p *float32 = i.Ptr()
	assert.Nil(p)

	v, err := i.Value()
	assert.Nil(err)
	assert.Nil(v)

	err = i.Scan(math.MaxFloat32)
	assert.Nil(err)
	assert.Equal(i.Float32, float32(math.MaxFloat32))
	assert.True(i.Valid)

	v, err = i.Value()
	assert.Nil(err)
	assert.Equal(v, float64(math.MaxFloat32))

	p = i.Ptr()
	assert.NotNil(p)
	assert.Equal(*p, float32(math.MaxFloat32))

	i = pqnull.PtrFloat32(p)
	assert.Equal(i.Float32, float32(math.MaxFloat32))
	assert.True(i.Valid)

	i = pqnull.ValidFloat32(42)
	assert.Equal(i.Float32, float32(42))
	assert.True(i.Valid)

	i = pqnull.PtrFloat32(nil)
	assert.Equal(i.Float32, float32(0))
	assert.False(i.Valid)
}

package pqarg_test

import (
	"github.com/maprost/assertion"
	"testing"
	"time"

	"github.com/maprost/pqx/pqarg"
)

func TestSimple(t *testing.T) {
	assert := assertion.New(t)

	tm := time.Now()
	args := pqarg.New()

	assert.Equal(args.Next(12), "$1")
	assert.Equal(args.Next("Blob"), "$2")
	assert.Equal(args.Next(tm), "$3")
	assert.Equal(args.Next(12.4), "$4")
	assert.Equal(args.Next(true), "$5")

	values := args.Get()
	assert.Len(values, 5)
	assert.Equal(values, []interface{}{12, "Blob", tm, 12.4, true})
}

func TestEmpty(t *testing.T) {
	assert := assertion.New(t)

	args := pqarg.New()

	values := args.Get()
	assert.NotNil(values)
	assert.Len(values, 0)
}

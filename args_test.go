package pqlib

import (
	"github.com/mleuth/assertion"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	assert := assertion.New(t)

	tm := time.Now()
	args := NewArgs()

	assert.Equal(args.Next(12), "$1")
	assert.Equal(args.Next("Blob"), "$2")
	assert.Equal(args.Next(tm), "$3")
	assert.Equal(args.Next(12.4), "$4")
	assert.Equal(args.Next(true), "$5")

	values := args.get()
	assert.Len(values, 5)
	assert.Equal(values, []interface{}{12, "Blob", tm, 12.4, true})
}

func TestEmpty(t *testing.T) {
	assert := assertion.New(t)

	args := NewArgs()

	values := args.get()
	assert.NotNil(values)
	assert.Len(values, 0)
}

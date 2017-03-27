package pqlib

import (
	"github.com/mleuth/timeutil"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const time_format string = "2006-Jan-02 15:04"

func TestSimple(t *testing.T) {
	check := require.New(t)

	tm, e := time.Parse(time_format, "2017-Mar-27 16:00")
	check.Nil(e)
	timeutil.Init(tm)

	args := NewArgs()

	check.Equal(args.Next(12), "$1")
	check.Equal(args.Next("Blob"), "$2")
	check.Equal(args.Next(timeutil.Now()), "$3")
	check.Equal(args.Next(12.4), "$4")
	check.Equal(args.Next(true), "$5")

	values := args.get()
	check.Len(values, 5)
	check.Equal(values[0], 12)
	check.Equal(values[1], "Blob")
	check.Equal(values[2], tm)
	check.Equal(values[3], 12.4)
	check.Equal(values[4], true)

	timeutil.Reset()
}

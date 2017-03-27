package pqlib

import (
	"strconv"
)

type Args struct {
	counter int
	args    []interface{}
}

func NewArgs() Args {
	return Args{counter: 1}
}

func (a *Args) Next(s interface{}) string {
	result := "$" + strconv.Itoa(a.counter)
	a.counter++
	a.args = append(a.args, s)
	return result
}

func (a Args) get() []interface{} {
	return a.args
}

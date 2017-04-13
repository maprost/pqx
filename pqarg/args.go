package pqarg

import (
	"strconv"
)

type Args struct {
	counter int
	args    []interface{}
}

func New() Args {
	return Args{counter: 1}
}

func (a *Args) Next(s interface{}) string {
	result := "$" + strconv.Itoa(a.counter)
	a.counter++
	a.args = append(a.args, s)
	return result
}

func (a Args) Get() []interface{} {
	return a.args
}

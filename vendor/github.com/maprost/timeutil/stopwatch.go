package timeutil

import "time"

type Stopwatch struct {
	start   time.Time
	elapsed time.Duration
}

func NewStopwatch() Stopwatch {
	s := Stopwatch{}
	s.Start()
	return s
}

func (s *Stopwatch) Start() {
	s.start = Now()
}

func (s *Stopwatch) Stop() string {
	s.elapsed = Now().Sub(s.start)
	return s.String()
}

func (s *Stopwatch) String() string {
	return s.elapsed.String()
}

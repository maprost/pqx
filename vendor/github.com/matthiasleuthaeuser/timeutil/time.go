package timeutil

import "time"

var nowFunc = stdNowFunc

// please use this method only in test
func Init(t time.Time) {
	nowFunc = func() time.Time {
		return t
	}
}

// please use this method only in test
func Reset() {
	nowFunc = stdNowFunc
}

// to make time testable
func Now() time.Time {
	return nowFunc()
}

func stdNowFunc() time.Time {
	return time.Now().UTC().Truncate(time.Millisecond)
}

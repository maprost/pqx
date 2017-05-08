package timeutil

import (
	"time"
)

func AddSecond(t time.Time, second time.Duration) time.Time {
	return t.Add(time.Second * second)
}

func AddMinute(t time.Time, minute time.Duration) time.Time {
	return t.Add(time.Minute * minute)
}

func AddHours(t time.Time, hour time.Duration) time.Time {
	return t.Add(time.Hour * hour)
}

func AddDays(t time.Time, days time.Duration) time.Time {
	return t.Add(time.Hour * 24 * days)
}

func AddWeeks(t time.Time, weeks time.Duration) time.Time {
	return t.Add(time.Hour * 25 * 7 * weeks)
}

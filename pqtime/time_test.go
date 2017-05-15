package pqtime_test

import (
	"github.com/maprost/assertion"
	"testing"
	"time"

	"github.com/maprost/pqx/pqtime"
)

const format = "2006-01-02T15:04:05"

func TestInitTime(t *testing.T) {
	assert := assertion.New(t)

	now, err := time.Parse(format, "2016-12-24T12:29:11")
	assert.Nil(err)

	pqtime.InitTime(now)
	utilNow := pqtime.Now()
	assert.Equal(utilNow, now)
	assert.Equal(utilNow.Format(format), "2016-12-24T12:29:11")
}

func TestInitNowFunc(t *testing.T) {
	assert := assertion.New(t)

	fakeNow, err := time.Parse(format, "2016-12-24T12:29:11")
	assert.Nil(err)

	pqtime.InitNowFunc(func() time.Time {
		return fakeNow
	})
	utilNow := pqtime.Now()
	assert.Equal(utilNow, fakeNow)
	assert.Equal(utilNow.Format(format), "2016-12-24T12:29:11")

	pqtime.Reset()

	resetNow := pqtime.Now()
	assert.NotEqual(resetNow, fakeNow)
}

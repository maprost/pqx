package pqnull

import (
	"database/sql/driver"
	"github.com/lib/pq"
	"time"
)

// Time represents a time that may be null.
// Time implements the Scanner interface so
// it can be used as a scan destination.
type Time struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (v *Time) Scan(value interface{}) (err error) {
	i := pq.NullTime{}
	err = i.Scan(value)
	if err == nil {
		v.Time = i.Time
		v.Valid = i.Valid
	}
	return err
}

// Value implements the driver Valuer interface.
func (v Time) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return v.Time, nil
}

func (v Time) Ptr() *time.Time {
	if v.Valid == false {
		return nil
	}
	return &v.Time
}

func ValidTime(value time.Time) Time {
	return Time{Time: value, Valid: true}
}

func NilTime() Time {
	return Time{Valid: false}
}

func PtrTime(value *time.Time) Time {
	if value == nil {
		return NilTime()
	}
	return ValidTime(*value)
}

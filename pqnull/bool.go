package pqnull

import (
	"database/sql"
	"database/sql/driver"
)

// Bool represents a bool that may be null.
// Bool implements the Scanner interface so
// it can be used as a scan destination.
type Bool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

// Scan implements the Scanner interface.
func (v *Bool) Scan(value interface{}) (err error) {
	i := sql.NullBool{}
	err = i.Scan(value)
	if err == nil {
		v.Bool = i.Bool
		v.Valid = i.Valid
	}
	return err
}

// Value implements the driver Valuer interface.
func (v Bool) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return v.Bool, nil
}

func (v Bool) Ptr() *bool {
	if v.Valid == false {
		return nil
	}
	return &v.Bool
}

func ValidBool(value bool) Bool {
	return Bool{Bool: value, Valid: true}
}

func NilBool() Bool {
	return Bool{Valid: false}
}

func PtrBool(value *bool) Bool {
	if value == nil {
		return NilBool()
	}
	return ValidBool(*value)
}

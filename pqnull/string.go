package pqnull

import (
	"database/sql"
	"database/sql/driver"
)

// String represents a string that may be null.
// String implements the Scanner interface so
// it can be used as a scan destination.
type String struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (v *String) Scan(value interface{}) (err error) {
	s := sql.NullString{}
	err = s.Scan(value)
	if err == nil {
		v.String = s.String
		v.Valid = s.Valid
	}
	return err
}

// Value implements the driver Valuer interface.
func (v String) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return v.String, nil
}

// Prt returns a string pointer.
func (v String) Ptr() *string {
	if v.Valid == false {
		return nil
	}
	return &v.String
}

// ValidString creates a valid String
func ValidString(value string) String {
	return String{String: value, Valid: true}
}

// NilString creates an 'invalid' String
func NilString() String {
	return String{Valid: false}
}

// PtrString creates a String out of a string pointer
func PtrString(value *string) String {
	if value == nil {
		return NilString()
	}
	return ValidString(*value)
}

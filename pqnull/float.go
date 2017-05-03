package pqnull

import (
	"database/sql"
	"database/sql/driver"
)

// ================= Float64 =======================

// Float64 represents an float64 that may be null.
// Float64 implements the Scanner interface so
// it can be used as a scan destination.
type Float64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
}

// Scan implements the Scanner interface.
func (v *Float64) Scan(value interface{}) (err error) {
	v.Float64, v.Valid, err = scanFloat64(value)
	return
}

// Value implements the driver Valuer interface.
func (v Float64) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return v.Float64, nil
}

// Prt returns an float64 pointer.
func (v Float64) Ptr() *float64 {
	if v.Valid == false {
		return nil
	}
	return &v.Float64
}

// ValidFloat64 creates a valid Float64
func ValidFloat64(value float64) Float64 {
	return Float64{Float64: value, Valid: true}
}

// NilFloat64 creates an 'invalid' Float64
func NilFloat64() Float64 {
	return Float64{Valid: false}
}

// PtrFloat64 creates an Float64 out of a float64 pointer
func PtrFloat64(value *float64) Float64 {
	if value == nil {
		return NilFloat64()
	}
	return ValidFloat64(*value)
}

// ================= int32 =======================

// Float32 represents an float32 that may be null.
// Float32 implements the Scanner interface so
// it can be used as a scan destination.
type Float32 struct {
	Float32 float32
	Valid   bool // Valid is true if Float32 is not NULL
}

// Scan implements the Scanner interface.
// If the value (float64) doesn't fit into the type (float32),
// it stores -1 as value.
func (v *Float32) Scan(value interface{}) error {
	f64, valid, err := scanFloat64(value)
	v.Float32 = float32(f64)
	v.Valid = valid
	return err
}

// Value implements the driver Valuer interface.
func (v Float32) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return float64(v.Float32), nil
}

// Prt returns an float32 pointer.
func (v Float32) Ptr() *float32 {
	if v.Valid == false {
		return nil
	}
	return &v.Float32
}

// ValidFloat32 creates a valid Float32
func ValidFloat32(value float32) Float32 {
	return Float32{Float32: value, Valid: true}
}

// NilFloat32 creates an 'invalid' Float32
func NilFloat32() Float32 {
	return Float32{Valid: false}
}

// PtrFloat32 creates an Float32 out of a float32 pointer
func PtrFloat32(value *float32) Float32 {
	if value == nil {
		return NilFloat32()
	}
	return ValidFloat32(*value)
}

// ================= Helper ===========================

func scanFloat64(v interface{}) (value float64, valid bool, err error) {
	i := sql.NullFloat64{}
	err = i.Scan(v)
	if err == nil {
		value = i.Float64
		valid = i.Valid
	}
	return
}

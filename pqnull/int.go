package pqnull

import (
	"database/sql"
	"database/sql/driver"
)

// ================= int64 =======================

// Int64 represents an int64 that may be null.
// Int64 implements the Scanner interface so
// it can be used as a scan destination.
type Int64 struct {
	Int64 int64
	Valid bool // Valid is true if Int64 is not NULL
}

// Scan implements the Scanner interface.
func (v *Int64) Scan(value interface{}) (err error) {
	v.Int64, v.Valid, err = scanInt64(value)
	return
}

// Value implements the driver Valuer interface.
func (v Int64) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return v.Int64, nil
}

// Prt returns an int64 pointer.
func (v Int64) Ptr() *int64 {
	if v.Valid == false {
		return nil
	}
	return &v.Int64
}

// ValidInt64 creates a valid Int64
func ValidInt64(value int64) Int64 {
	return Int64{Int64: value, Valid: true}
}

// NilInt64 creates an 'invalid' Int64
func NilInt64() Int64 {
	return Int64{Valid: false}
}

// PtrInt64 creates an Int64 out of a int64 pointer
func PtrInt64(value *int64) Int64 {
	if value == nil {
		return NilInt64()
	}
	return ValidInt64(*value)
}

// ================= int =======================

// Int represents an int that may be null.
// Int implements the Scanner interface so
// it can be used as a scan destination.
type Int struct {
	Int   int
	Valid bool // Valid is true if Int is not NULL
}

// Scan implements the Scanner interface.
func (v *Int) Scan(value interface{}) (err error) {
	i64, valid, err := scanInt64(value)
	v.Int = int(i64)
	v.Valid = valid
	return err
}

// Value implements the driver Valuer interface.
func (v Int) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return int64(v.Int), nil
}

// Prt returns an int pointer.
func (v Int) Ptr() *int {
	if v.Valid == false {
		return nil
	}
	return &v.Int
}

// ValidInt creates a valid Int
func ValidInt(value int) Int {
	return Int{Int: value, Valid: true}
}

// NilInt creates an 'invalid' Int
func NilInt() Int {
	return Int{Valid: false}
}

// PtrInt creates an Int out of a int pointer
func PtrInt(value *int) Int {
	if value == nil {
		return NilInt()
	}
	return ValidInt(*value)
}

// ================= int32 =======================

// Int32 represents an int32 that may be null.
// Int32 implements the Scanner interface so
// it can be used as a scan destination.
type Int32 struct {
	Int32 int32
	Valid bool // Valid is true if Int32 is not NULL
}

// Scan implements the Scanner interface.
// If the value (int64) doesn't fit into the type (int32),
// it stores -1 as value.
func (v *Int32) Scan(value interface{}) error {
	i64, valid, err := scanInt64(value)
	v.Int32 = int32(i64)
	v.Valid = valid
	return err
}

// Value implements the driver Valuer interface.
func (v Int32) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return int64(v.Int32), nil
}

// Prt returns an int32 pointer.
func (v Int32) Ptr() *int32 {
	if v.Valid == false {
		return nil
	}
	return &v.Int32
}

// ValidInt32 creates a valid Int32
func ValidInt32(value int32) Int32 {
	return Int32{Int32: value, Valid: true}
}

// NilInt32 creates an 'invalid' Int32
func NilInt32() Int32 {
	return Int32{Valid: false}
}

// PtrInt32 creates an Int32 out of a int32 pointer
func PtrInt32(value *int32) Int32 {
	if value == nil {
		return NilInt32()
	}
	return ValidInt32(*value)
}

// ================= int16 =======================

// Int16 represents an int16 that may be null.
// Int16 implements the Scanner interface so
// it can be used as a scan destination.
type Int16 struct {
	Int16 int16
	Valid bool // Valid is true if Int16 is not NULL
}

// Scan implements the Scanner interface.
// If the value (int64) doesn't fit into the type (int16),
// it stores -1 as value.
func (v *Int16) Scan(value interface{}) error {
	i64, valid, err := scanInt64(value)
	v.Int16 = int16(i64)
	v.Valid = valid
	return err
}

// Value implements the driver Valuer interface.
func (v Int16) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return int64(v.Int16), nil
}

// Prt returns an int16 pointer.
func (v Int16) Ptr() *int16 {
	if v.Valid == false {
		return nil
	}
	return &v.Int16
}

// ValidInt16 creates a valid Int16
func ValidInt16(value int16) Int16 {
	return Int16{Int16: value, Valid: true}
}

// NilInt16 creates an 'invalid' Int16
func NilInt16() Int16 {
	return Int16{Valid: false}
}

// PtrInt16 creates an Int16 out of a int16 pointer
func PtrInt16(value *int16) Int16 {
	if value == nil {
		return NilInt16()
	}
	return ValidInt16(*value)
}

// ================= int8 =======================

// Int8 represents an int8 that may be null.
// Int8 implements the Scanner interface so
// it can be used as a scan destination.
type Int8 struct {
	Int8  int8
	Valid bool // Valid is true if Int8 is not NULL
}

// Scan implements the Scanner interface.
// If the value (int64) doesn't fit into the type (int8),
// it stores -1 as value.
func (v *Int8) Scan(value interface{}) error {
	i64, valid, err := scanInt64(value)
	v.Int8 = int8(i64)
	v.Valid = valid
	return err
}

// Value implements the driver Valuer interface.
func (v Int8) Value() (driver.Value, error) {
	if !v.Valid {
		return nil, nil
	}
	return int64(v.Int8), nil
}

// Prt returns an int8 pointer.
func (v Int8) Ptr() *int8 {
	if v.Valid == false {
		return nil
	}
	return &v.Int8
}

// ValidInt8 creates a valid Int8
func ValidInt8(value int8) Int8 {
	return Int8{Int8: value, Valid: true}
}

// NilInt8 creates an 'invalid' Int8
func NilInt8() Int8 {
	return Int8{Valid: false}
}

// PtrInt8 creates an Int8 out of a int8 pointer
func PtrInt8(value *int8) Int8 {
	if value == nil {
		return NilInt8()
	}
	return ValidInt8(*value)
}

// ================= Helper ===========================

func scanInt64(v interface{}) (value int64, valid bool, err error) {
	i := sql.NullInt64{}
	err = i.Scan(v)
	if err == nil {
		value = i.Int64
		valid = i.Valid
	}
	return
}

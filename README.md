[![GoDoc](https://godoc.org/github.com/mleuth/pqlib?status.svg)](https://godoc.org/github.com/mleuth/pqlib)

# pqx
Small lightweight postgres library.

## Actions
- insert
- update
- delete
- create
- register

## Supported Types
- int8, int16, int/int32, int64
- float32, float64
- bool
- string
- time.Time
- sql.NullBool, sql.NullString, sql.NullFloat64, sql.NullInt64

## Supported Tags
- `sql:"PK"` -> primary key
- `sql:"AI"` -> auto increment
    - available for: int8, uint8, int16, uint16, int/int32, uint/uint32, int64
- `sql:"CreateDate"` -> set initial time automatically
- `sql:"ChangeDate"` -> set initial and update time automatically
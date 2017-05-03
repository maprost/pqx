[![GoDoc](https://godoc.org/github.com/mleuth/pqlib?status.svg)](https://godoc.org/github.com/mleuth/pqlib)

# pqx v0.2
Small lightweight pq library extension. 

## Features
- log query feature
- simplify default actions (see actions)
- scan a struct 
- sql parameter without $1-$n

## Actions
- insert
- update
- delete
- create
- register
- contains (by id)
- select (by id)

## Supported Types
- Integer:
  - `int8`, `int16`, `int`, `int32`, `int64`
  - `sql.NullInt64` 
  - `pqnull.Int64`, `pqnull.Int`, `pqnull.Int32`, `pqnull.Int16`, `pqnull.Int8`
- Unsigned Integer:
  - `uint8`, `uint16`, `uint`, `uint32`, `uint64`
- Float:
  - `float32`, `float64`
  - `sql.NullFloat64`
  - `pqnull.Float64`, `pqnull.Float32`
- Boolean:
  - `bool`
  - `sql.NullBool`
  - `pqnull.Bool`
- String:
  - `string`
  - `sql.NullString`
  - `pqnull.String`
- Time:
  - `time.Time`
  - `pq.NullTime`
  - `pqnull.Time`

## Supported Tags
- `pqx:"PK"` -> primary key
- `pqx:"AI"` -> auto increment
    - available for: `int8`, `uint8`, `int16`, `uint16`, `int`, `int32`, `uint`, `uint32`, `int64`
- `pqx:"Created"` -> set initial time automatically
- `pqx:"Changed"` -> set initial and update time automatically

## Usage
```go


```
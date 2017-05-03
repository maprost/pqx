package pqtype

import "strconv"

type Type int

const (
	Text Type = iota
	SmallSerial
	Serial
	BigSerial
	SmallInt
	Integer
	BigInt
	Numeric
	Bool
	Real
	DoublePrecision
	Time
)

func (t Type) String() string {
	if int(t) < len(typeNames) {
		return typeNames[t]
	}
	return "pqtype" + strconv.Itoa(int(t))
}

var typeNames = []string{
	Text:            "text",
	SmallSerial:     "smallserial",
	Serial:          "serial",
	BigSerial:       "bigserial",
	SmallInt:        "smallint",
	Integer:         "integer",
	BigInt:          "bigint",
	Numeric:         "numeric",
	Bool:            "bool",
	Real:            "real",
	DoublePrecision: "double precision",
	Time:            "timestamp with time zone",
}

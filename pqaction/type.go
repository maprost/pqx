package pqaction

import (
	"errors"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
	"strings"
	"time"
)

func dbType(field pqreflect.Field) (dbType string, e error) {
	switch field.TypeInterface().(type) {
	case string:
		dbType = "text"

	case int8, int16:
		if isAutoIncrement(field) {
			dbType = "smallserial"
		} else {
			dbType = "smallint"
		}
	case int, int32:
		if isAutoIncrement(field) {
			dbType = "serial"
		} else {
			dbType = "integer"
		}
	case int64:
		if isAutoIncrement(field) {
			dbType = "bigserial"
		} else {
			dbType = "bigint"
		}

	case bool:
		dbType = "bool"

	case float32:
		dbType = "real"
	case float64:
		dbType = "double precision"

	case time.Time:
		dbType = "timestamp with time zone"

	default:
		e = errors.New("Not supported field type: " + field.Name() + " (" + field.Type() + ").")
		return
	}
	return
}

func isPrimaryKey(field pqreflect.Field) bool {
	return strings.Contains(field.Tag("sql"), "PK")
}

func isAutoIncrement(field pqreflect.Field) bool {
	return strings.Contains(field.Tag("sql"), "AI")
}

func isCreateDate(field pqreflect.Field) bool {
	return strings.Contains(field.Tag("sql"), "CreateDate")
}

func isChangeDate(field pqreflect.Field) bool {
	return strings.Contains(field.Tag("sql"), "ChangeDate")
}

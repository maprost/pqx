package pqaction

import (
	"github.com/mleuth/pqlib/pqutil/pqreflect"
	"strings"
)

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

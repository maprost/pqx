package pqreflect

import (
	"reflect"
	"strings"
)

type StructInfo struct {
	pkg       string
	name      string
	fieldList []Field
}

func NewStructInfo(s interface{}) StructInfo {
	fieldList := []Field{}

	value := reflect.ValueOf(s)
	elem := value.Elem()
	for i := 0; i < elem.NumField(); i++ {
		fieldList = append(fieldList, NewField(elem, i))
	}

	name := reflect.Indirect(value).Type().Name()
	pkg := ""
	structName := ""
	i := strings.Index(name, ".")
	if i > -1 {
		structName = name[i+1:]
		pkg = name[:i]
	} else {
		structName = name
	}
	return StructInfo{
		name:      strings.ToLower(structName),
		pkg:       strings.ToLower(pkg),
		fieldList: fieldList}
}

func (s *StructInfo) Name() string {
	return s.name
}

func (s *StructInfo) Package() string {
	return s.pkg
}

func (s *StructInfo) Fields() []Field {
	return s.fieldList
}

package pqtable

import (
	"errors"
	"reflect"
	"strings"
)

type Table struct {
	name    string
	columns []Column
}

func New(table interface{}) (*Table, error) {
	return NewCtx(table, defaultCtx)
}

func NewCtx(table interface{}, ctx Context) (*Table, error) {
	tableValue := reflect.Indirect(reflect.ValueOf(table))
	if tableValue.Kind() != reflect.Struct {
		return nil, errors.New("Value(" + tableValue.Kind().String() + ") is not a struct.")
	}

	// create columns
	columns := []Column{}
	for i := 0; i < tableValue.NumField(); i++ {
		c, err := newColumn(tableValue, i, ctx)
		if err != nil {
			return nil, err
		}

		// insert only PK column if OnlyPrimaryKeyColumn is active
		if ctx.OnlyPrimaryKeyColumn == false || c.PrimaryKeyTag() {
			columns = append(columns, c)
		}
	}

	return &Table{
		name:    tableName(&tableValue),
		columns: columns}, nil
}

func (s *Table) Name() string {
	return s.name
}

func (s *Table) Columns() []Column {
	return s.columns
}

func (t *Table) Len() int {
	return len(t.columns)
}

func TableName(entity interface{}) string {
	value := reflect.Indirect(reflect.ValueOf(entity))
	return tableName(&value)
}

func tableName(tableValue *reflect.Value) string {
	name := tableValue.Type().Name()
	i := strings.Index(name, ".")

	return strings.ToLower(name[i+1:])
}

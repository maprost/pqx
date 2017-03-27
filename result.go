package pqlib

import (
	"database/sql"
	"errors"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
	"time"
)

type Result struct {
	rows    *sql.Rows
	hasNext bool
}

func (r *Result) HasNext() bool {
	if r.hasNext == false {
		r.hasNext = r.rows.Next()
	}
	return r.hasNext
}

func (r *Result) Scan(output ...interface{}) error {
	defer r.resetHasNext()

	// is there something to scan?
	if r.HasNext() == false {
		return errors.New("No rows available.")
	}

	e := r.rows.Scan(output...)
	if e != nil {
		return e
	}

	return nil
}

func (r *Result) ScanStruct(output interface{}) error {
	defer r.resetHasNext()

	valueList, afterAction, err := r.init(output)
	if err != nil {
		return err
	}

	err = r.Scan(valueList...)
	if err != nil {
		return err
	}

	for _, action := range afterAction {
		action()
	}

	return nil
}

func (r *Result) init(output interface{}) (valueList []interface{}, afterAction []func(), e error) {
	structInfo := pqreflect.NewStructInfo(output)

	for _, f := range structInfo.Fields() {
		field := f
		switch field.TypeInterface().(type) {
		case string:
			var str string
			valueList = append(valueList, &str)
			afterAction = append(afterAction, func() {
				field.SetString(str)
			})
		case int, int8, int16, int32, int64:
			var i int64
			valueList = append(valueList, &i)
			afterAction = append(afterAction, func() {
				field.SetInt(i)
			})
		case bool:
			var b bool
			valueList = append(valueList, &b)
			afterAction = append(afterAction, func() {
				field.SetBool(b)
			})
		case float32, float64:
			var d float64
			valueList = append(valueList, &d)
			afterAction = append(afterAction, func() {
				field.SetFloat(d)
			})
		case time.Time:
			var t time.Time
			valueList = append(valueList, &t)
			afterAction = append(afterAction, func() {
				field.SetTime(t)
			})
		default:
			e = errors.New("Not supported field type: " + field.Name() + " (" + field.Type() + ").")
			return
		}
	}

	return
}

func (r *Result) close() error {
	return r.rows.Close()
}

func (r *Result) resetHasNext() {
	r.hasNext = false
}

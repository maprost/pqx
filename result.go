package pqx

import (
	"database/sql"
	"errors"
	"github.com/maprost/pqx/pqutil/pqreflect"
	"reflect"
	"time"
)

type Result interface {
	Scan(dest ...interface{}) error
}

func ScanStruct(r Result, output interface{}) error {
	valueList, afterAction, err := initScan(output)
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

func initScan(output interface{}) (valueList []interface{}, afterAction []func(), e error) {
	structInfo := pqreflect.NewStructInfo(output)

	for _, f := range structInfo.Fields() {
		field := f
		switch field.Kind() {
		case reflect.String:
			var str string
			valueList = append(valueList, &str)
			afterAction = append(afterAction, func() {
				field.SetString(str)
			})
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i int64
			valueList = append(valueList, &i)
			afterAction = append(afterAction, func() {
				field.SetInt(i)
			})
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i uint64
			valueList = append(valueList, &i)
			afterAction = append(afterAction, func() {
				field.SetUint(i)
			})
		case reflect.Bool:
			var b bool
			valueList = append(valueList, &b)
			afterAction = append(afterAction, func() {
				field.SetBool(b)
			})
		case reflect.Float32, reflect.Float64:
			var d float64
			valueList = append(valueList, &d)
			afterAction = append(afterAction, func() {
				field.SetFloat(d)
			})

		default:
			switch field.TypeInterface().(type) {
			case time.Time:
				var t time.Time
				valueList = append(valueList, &t)
				afterAction = append(afterAction, func() {
					field.SetTime(t)
				})
			case sql.NullBool:
				var b sql.NullBool
				valueList = append(valueList, &b)
				afterAction = append(afterAction, func() {
					field.SetNullBool(b)
				})
			case sql.NullString:
				var s sql.NullString
				valueList = append(valueList, &s)
				afterAction = append(afterAction, func() {
					field.SetNullString(s)
				})
			case sql.NullInt64:
				var i sql.NullInt64
				valueList = append(valueList, &i)
				afterAction = append(afterAction, func() {
					field.SetNullInt64(i)
				})
			case sql.NullFloat64:
				var f sql.NullFloat64
				valueList = append(valueList, &f)
				afterAction = append(afterAction, func() {
					field.SetNullFloat64(f)
				})
			default:
				e = errors.New("Not supported field type: " + field.Name() + " (" + field.Type() + ").")
				return
			}

		}
	}

	return
}

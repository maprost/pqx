package pqaction

import (
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
)

func RegisterList(tx pqlib.Transaction, entities []interface{}) error {
	for _, entity := range entities {
		e := Register(tx, entity)
		if e != nil {
			return e
		}
	}

	return nil
}

func Register(tx pqlib.Transaction, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	exists, e := tableExists(tx, structInfo)
	if e != nil {
		return e
	}
	if exists {
		return nil
	}

	e = Create(tx, entity)
	return e
}

func tableExists(tx pqlib.Transaction, structInfo pqreflect.StructInfo) (bool, error) {
	args := pqlib.NewArgs()
	result, e := tx.Query(
		"SELECT table_name FROM INFORMATION_SCHEMA.TABLES "+
			"WHERE TABLE_NAME = "+args.Next(structInfo.Name()), args)
	if e != nil {
		return false, e
	}
	// no table with that name is known
	if result.HasNext() == false {
		return false, nil
	}

	// TODO: check columns

	// table exists and there is no error
	return true, nil
}

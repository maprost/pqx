package pqaction

import (
	"github.com/mleuth/pqlib"
	"github.com/mleuth/pqlib/pqutil/pqreflect"
)

func RegisterList(db pqlib.Transaction, entities []interface{}) error {
	for _, entity := range entities {
		e := Register(db, entity)
		if e != nil {
			return e
		}
	}

	return nil
}

func Register(db pqlib.Transaction, entity interface{}) error {
	structInfo := pqreflect.NewStructInfo(entity)

	exists, e := tableExists(db, structInfo)
	if e != nil {
		return e
	}
	if exists {
		return nil
	}

	e = Create(db, entity)
	return e
}

func tableExists(db pqlib.Transaction, structInfo pqreflect.StructInfo) (bool, error) {
	args := pqlib.NewArgs()
	result, e := db.Query(
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

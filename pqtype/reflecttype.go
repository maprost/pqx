package pqtype

type ReflectType int

const (
	// go types
	ReflectString ReflectType = iota
	ReflectUInt64
	ReflectInt64
	ReflectBool
	ReflectFloat64
	ReflectTime

	// library (sql, pq) null types
	ReflectTime_pq
	ReflectString_sql
	ReflectInt64_sql
	ReflectFloat64_sql
	ReflectBool_sql

	// pqx null types
	ReflectString_pqx
	ReflectTime_pqx
	ReflectInt64_pqx
	ReflectInt_pqx
	ReflectInt32_pqx
	ReflectInt16_pqx
	ReflectInt8_pqx
	ReflectFloat64_pqx
	ReflectFloat32_pqx
	ReflectBool_pqx
)

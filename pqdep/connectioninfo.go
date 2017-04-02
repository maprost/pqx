package pqdep

type ConnectionInfo interface {
	DatabaseDriver() string
	DataBase() string
	Host() string
	Port() string
	UserName() string
}

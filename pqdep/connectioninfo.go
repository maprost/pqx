package pqdep

type ConnectionInfo interface {
	GetDatabaseDriver() string
	GetDataBase() string
	GetHost() string
	GetPort() string
	GetUserName() string
}

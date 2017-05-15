package pqdep

type ConnectionInfo interface {
	Driver() string
	Database() string
	Host() string
	Port() string
	Username() string
}

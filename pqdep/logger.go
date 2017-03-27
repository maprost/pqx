package pqdep

type Logger interface {
	Printf(format string, v ...interface{})
}

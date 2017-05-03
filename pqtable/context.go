package pqtable

var defaultCtx = Context{}

type Context struct {
	OnlyPrimaryKeyColumn    bool
	IgnoreUnknownColumnType bool
}

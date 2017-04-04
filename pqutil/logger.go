package pqutil

import (
	"log"
	"os"
)

var DefaultLogger = log.New(os.Stdout, "", 0)

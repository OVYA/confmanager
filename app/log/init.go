package log

import (
	"log"
	"os"
)

var (
	Info *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
}

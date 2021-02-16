package hub

import (
	"log"
)

func Trace(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Err(format string, v ...interface{}) {
	log.Printf(format, v...)
}

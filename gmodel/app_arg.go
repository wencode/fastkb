package gmodel

import (
	"flag"
)

var (
	DropInterval = 500
)

func init() {
	flag.IntVar(&DropInterval, "dropinter", 500, "letters drop interval time")
}

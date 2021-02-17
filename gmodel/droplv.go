package gmodel

import (
	"time"
)

type Char struct {
	C byte
	X float32
	Y float32
}

type _DropLv struct {
	GenInterval time.Duration
	MoveVel     float32

	CharList []Char
}

var (
	DropLv *_DropLv
)

func init() {
	DropLv = &_DropLv{
		CharList: make(Char, 0),
	}

}

func Add(c byte, x float32) {
	char := Char{
		C: c,
		X: x,
	}
	for i, v := range DropLv.CharList {
		if v.C == 0 {
			DropLv.CharList[i] = char
			return
		}
	}
	DropLv.CharList = append(DropLv.CharList, char)
}

func Remove(idx int) {
	DropLv.CharList[idx].C = 0
}

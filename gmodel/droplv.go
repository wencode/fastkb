package gmodel

import (
	"time"
)

type Char struct {
	C   byte
	Str string
	X   float64
	Y   float64
}

type _DropLv struct {
	GenX0, GenX1 int
	GenY         int
	DelY         float64
	GenInterval  time.Duration
	MoveVel      float64

	CharList []*Char

	Time     time.Duration
	LifeNum  int
	PointNum int
	Miss     int
}

var (
	DropLv *_DropLv
)

func init() {
	DropLv = &_DropLv{
		CharList: make([]*Char, 0),
	}

}

func (d *_DropLv) Add(c byte, x float64) {
	for _, v := range d.CharList {
		if v.C == 0 {
			v.C = c
			v.Str = string([]byte{c})
			v.X = x
			v.Y = float64(d.GenY)
			return
		}
	}
	char := &Char{
		C:   c,
		Str: string([]byte{c}),
		X:   x,
		Y:   float64(d.GenY),
	}
	d.CharList = append(d.CharList, char)
}

func (d *_DropLv) Remove(idx int) {
	d.CharList[idx].C = 0
}

func (d *_DropLv) CheckChar(c byte) int {
	idx := -1
	for k, char := range d.CharList {
		if char.C == c {
			if idx >= 0 && d.CharList[idx].Y >= char.Y {
				continue
			}
			idx = k
		}
	}
	return idx
}

func (d *_DropLv) RemoveByList(idxlist []int) {
	for _, idx := range idxlist {
		d.CharList[idx].C = 0
	}
}

package droplv

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/wencode/fastkb/audio"
	"github.com/wencode/fastkb/font"
	"github.com/wencode/fastkb/gmodel"
	"github.com/wencode/fastkb/hub"
)

var (
	dropableList = make([]byte, 0)
	mod          *Module
)

func init() {
	// for i := byte(0); i < 26; i++ {
	// 	dropableList = append(dropableList, 'A'+i)
	// }
	for i := byte(0); i < 26; i++ {
		dropableList = append(dropableList, 'a'+i)
	}
	// for i := byte(0); i < 10; i++ {
	// 	dropableList = append(dropableList, '0'+i)
	// }

	mod = &Module{}
	hub.Register(mod)
	hub.Listen(mod, "mainui", hub.Notify_mainui_LevelStart)
}

type Module struct {
	genCounter int64

	life      int
	lifeStr   string
	pointsStr string
}

func (m *Module) Name() string { return "droplv" }

func (m *Module) Init() error {
	data := gmodel.DropLv
	data.GenX0 = 50
	data.GenX1 = 950
	data.DelY = 500
	data.GenInterval = time.Second * 1
	data.MoveVel = 70

	m.life = 5
	m.lifeStr = strconv.Itoa(m.life)
	m.pointsStr = strconv.Itoa(gmodel.DropLv.Points)

	return nil
}

func (m *Module) Update(delta time.Duration) error {
	if m.life > 0 {
		m.checkIntput()

		m.genChar(delta)
		delList := m.updatePosition(delta)
		m.missChar(delList)
	}

	return nil
}

func (m *Module) Draw(screen *ebiten.Image) {
	for _, c := range gmodel.DropLv.CharList {
		if c.C == 0 {
			continue
		}
		font.DrawBigString(screen, c.Str, int(c.X), int(c.Y))
		//hub.Trace("draw %s [%f,%f]", c.Str, c.X, c.Y)
	}

	font.DrawString(screen, m.lifeStr, 10, 50, color.RGBA{255, 0, 0, 255})
	font.DrawString(screen, m.pointsStr, 950, 50, color.RGBA{0, 255, 255, 255})
}

func (m *Module) OnNotify(ntf hub.Notify, arg0, arg1 int, arg interface{}) {}

func (m *Module) checkIntput() {
	data := gmodel.DropLv
	for i := byte(0); i < 26; i++ {
		key := ebiten.Key(int(ebiten.KeyA) + int(i))
		if !inpututil.IsKeyJustReleased(key) {
			continue
		}
		idx := data.CheckChar('a' + i)
		if idx < 0 {
			continue
		}

		// hitted
		m.addPoint()
		audio.Hit()
		hub.Trace("hit %c at %d", 'a'+i, idx)
		data.Remove(idx)
	}
}

func (m *Module) updatePosition(delta time.Duration) []int {
	data := gmodel.DropLv
	distance := data.MoveVel * delta.Seconds()
	delList := make([]int, 0)
	for k, c := range data.CharList {
		if c.C == 0 {
			continue
		}
		c.Y += distance
		if c.Y >= data.DelY {
			delList = append(delList, k)
		}
	}
	return delList
}

func (m *Module) missChar(idxlist []int) {
	data := gmodel.DropLv
	for _, idx := range idxlist {
		data.CharList[idx].C = 0
		audio.Miss()
		data.Miss++
	}

	if data.Miss > 10 {
		m.decLife()
		hub.Trace("miss %d, decrease a lift", data.Miss)
		data.Miss -= 10
	}
}

func (m *Module) genChar(delta time.Duration) {
	data := gmodel.DropLv

	m.genCounter -= delta.Nanoseconds()
	if m.genCounter > 0 {
		return
	}
	m.genCounter = data.GenInterval.Nanoseconds()

	n := rand.Intn(len(dropableList))
	x := data.GenX0 + rand.Intn(data.GenX1-data.GenX0)
	data.Add(dropableList[n], float64(x))
}

func (m *Module) addPoint() {
	data := gmodel.DropLv
	data.Points++
	m.pointsStr = strconv.Itoa(data.Points)
}

func (m *Module) decLife() {
	m.life--
	if m.life >= 0 {
		m.lifeStr = strconv.Itoa(m.life)
	}
}

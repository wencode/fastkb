package mainui

import (
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/wencode/fastkb/hub"
	"github.com/wencode/fastkb/util"
)

var mod *Module

func init() {
	mod = &Module{}
	hub.Register(mod)
	hub.Listen(mod, "app", hub.Notif_app_Start)
	hub.Listen(mod, "droplv", hub.Notif_LevelOver)
	hub.Listen(mod, "droplv", hub.Notif_LevelEnd)
}

const (
	stateHidden = iota
	stateStart
	stateOver
	stateEnd
)

type Module struct {
	state    int
	startImg *ebiten.Image
	overImg  *ebiten.Image
	endImg   *ebiten.Image
}

func (m *Module) Name() string { return "mainui" }

func (m *Module) Init() error {
	var err error
	m.startImg, err = util.LoadImage("./start.png")
	if err != nil {
		hub.Err("load start.png error:%v", err)
	}
	m.overImg, err = util.LoadImage("./gameover.png")
	if err != nil {
		hub.Err("load gameover.png error:%v", err)
	}
	m.endImg, err = util.LoadImage("./end.png")
	if err != nil {
		hub.Err("load end.png error:%v", err)
	}
	return nil
}

func (m *Module) Update(delta time.Duration) error {
	if m.state > 0 {
		if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
			hub.LightNotify(mod, hub.Notif_mainui_LevelStart)
			m.state = stateHidden
		}
	}
	return nil
}

func (m *Module) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	op.GeoM.Translate(float64(sw/2-110), float64(sh/2-100))
	switch m.state {
	case stateStart:
		if m.startImg != nil {
			screen.DrawImage(m.startImg, op)
		}
	case stateOver:
		if m.overImg != nil {
			screen.DrawImage(m.overImg, op)
		}
	case stateEnd:
		if m.endImg != nil {
			screen.DrawImage(m.endImg, op)
		}
	}
}

func (m *Module) OnNotify(ntf hub.Notif, arg0, arg1 int, arg interface{}) {
	switch ntf {
	case hub.Notif_app_Start:
		m.state = stateStart
	case hub.Notif_LevelOver:
		m.state = stateOver
	case hub.Notif_LevelEnd:
		m.state = stateEnd
	}
}

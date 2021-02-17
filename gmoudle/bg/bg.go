package bg

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/wencode/fastkb/audio"
	"github.com/wencode/fastkb/hub"
	"github.com/wencode/fastkb/util"
)

var mod *Module

func init() {
	mod = &Module{}
	hub.Register(mod)
	hub.Listen(mod, "app", hub.Notify_app_Start)
}

type Module struct {
	img *ebiten.Image
}

func (m *Module) Name() string { return "bg" }

func (m *Module) Init() error {
	audio.PlayBG()
	return nil
}

func (m *Module) Update(delta time.Duration) error {
	return nil
}

func (m *Module) Draw(screen *ebiten.Image) {
	if img := m.img; img != nil {
		sw, sh := screen.Size()
		w, h := img.Size()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(sw)/float64(w), float64(sh)/float64(h))

		screen.DrawImage(m.img, op)
	}
}

func (m *Module) OnNotify(ntf hub.Notify, arg0, arg1 int, arg interface{}) {
	switch ntf {
	case hub.Notify_app_Start:
		m.onAppStart()
	}
}

func (m *Module) onAppStart() {
	img, err := util.LoadImage("./bg1.jpeg")
	if err != nil {
		hub.Err("load bg image error: %v", err)
	}
	m.img = img
	hub.Trace("[bg] load bg image ok")
}

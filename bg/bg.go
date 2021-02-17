package bg

import (
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/wencode/fastkb/hub"
	"github.com/wencode/fastkb/util"
)

var mod *Module

func init() {
	mod = &Module{}
	hub.Register(mod)
	hub.Listen(mod, "main", hub.Notify_Main_AppStart)
}

type Module struct {
	img *ebiten.Image
}

func (m *Module) Name() string { return "bg" }

func (m *Module) Update(delta time.Duration) error {
	return nil
}

func (m *Module) Draw(screen *ebiten.Image) {
	if m.img != nil {
		screen.DrawImage(m.img, &ebiten.DrawImageOptions{})
	}
}

func (m *Module) OnNotify(ntf hub.Notify, arg0, arg1 int, arg interface{}) {
	switch ntf {
	case hub.Notify_Main_AppStart:
		m.onAppStart()
	}
}

func (m *Module) onAppStart() {
	img, err := util.LoadImage("./morning_bg.jpeg")
	if err != nil {
		hub.Err("load bg image error: %v", err)
	}
	m.img = img
	hub.Trace("[bg] load bg image ok")
}

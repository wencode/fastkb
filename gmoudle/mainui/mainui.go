package mainui

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wencode/fastkb/hub"
)

var mod *Module

func init() {
	mod = &Module{}
	hub.Register(mod)
	hub.Listen(mod, "app", hub.Notify_app_Start)
}

type Module struct {
}

func (m *Module) Name() string { return "mainui" }

func (m *Module) Update(delta time.Duration) error {
	return nil
}

func (m *Module) Draw(screen *ebiten.Image) {
}

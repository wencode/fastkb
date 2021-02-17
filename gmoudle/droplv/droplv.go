package droplv

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wencode/fastkb/hub"
)

var mod *Module

func init() {
	mod = &Module{}
	hub.Register(mod)
	hub.Listen(mod, "mainui", hub.Notify_mainui_LevelStart)
}

type Module struct {
}

func (m *Module) Name() string { return "droplv" }

func (m *Module) Update(delta time.Duration) error {
	return nil
}

func (m *Module) Draw(screen *ebiten.Image) {
}

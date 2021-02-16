package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/wencode/fastkb/hub"

	// Modules
	_ "github.com/wencode/fastkb/bg"
)

type Game struct {
	lastTime time.Time
	screen   *ebiten.Image
}

func (g *Game) Update() error {
	// When IsDrawingSkipped is true, the rendered result is not adopted.
	// Skip rendering then.
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	now := time.Now()
	delta := now.Sub(g.lastTime)
	g.lastTime = now
	_, err := hub.Update(delta)
	return err
}

func (g *Game) callModDraw(mod hub.Module) {
	mod.Draw(g.screen)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screen = screen
	hub.SyncExec(g.callModDraw)
	g.screen = nil
}

func (g *Game) Layout(outsideWh, outsideHt int) (screenWidth, screenHt int) {
	return 400, 300
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Fast Keyboard")
	game.lastTime = time.Now()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"math/rand"
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/wencode/fastkb/hub"

	// Modules
	_ "github.com/wencode/fastkb/gmoudle/bg"
	_ "github.com/wencode/fastkb/gmoudle/droplv"
	_ "github.com/wencode/fastkb/gmoudle/mainui"
)

type Game struct {
	frameNum int
	lastTime time.Time
	screen   *ebiten.Image
}

func (g *Game) Update() error {
	g.frameNum++
	if g.frameNum == 1 {
		hub.SyncExec(func(mod hub.Module) {
			mod.Init()
		})
		hub.NotifyingDelegate("app", hub.Notify_app_Start, 0, 0, nil)
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHt int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := &Game{}
	ebiten.SetWindowSize(1000, 600)
	ebiten.SetWindowTitle("Fast Keyboard")
	game.lastTime = time.Now()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

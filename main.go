package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	// When IsDrawingSkipped is true, the rendered result is not adopted.
	// Skip rendering then.
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.Draw(screen)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWh, outsideHt int) (screenWidth, screenHt int) {
	return 400, 300
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Fast Keyboard")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

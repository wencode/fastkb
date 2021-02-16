package draw

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawing interface {
	Draw(screen *ebiten.Image)
}

package util

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	eimg := ebiten.NewImageFromImage(img)
	return eimg
}

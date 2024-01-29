package resource

import (
	"embed"
	"image"

	"fyne.io/fyne/v2/canvas"
)

//go:embed *
var Resource embed.FS

var Ebiten = GetResource("ebiten.png")

func GetResource(name string) *canvas.Image {
	f, err := Resource.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return canvas.NewImageFromImage(img)
}

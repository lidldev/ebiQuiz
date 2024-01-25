package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainPageUI(label *canvas.Text, button *widget.Button) *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		container.NewVBox(label),
		layout.NewSpacer(),
		container.NewVBox(button),
	)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(635, 475))

	text := canvas.NewText("Welcome To Ebitengine Quiz!!", color.NRGBA{255, 115, 22, 255})
	text.TextSize = 30
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Italic: false, Bold: true, Monospace: false}

	w.SetContent(MainPageUI(text, widget.NewButton("Click me!", func() {

	})))

	w.Show()
	a.Run()
}

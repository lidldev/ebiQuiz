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

func MainPageUI(label *canvas.Text, button *widget.Button, line *canvas.Line) *fyne.Container {
	return container.New(layout.NewBorderLayout(label, line, nil, nil),
		label,
		container.NewVBox(button),
		line,
	)
}

func main() {
	a := app.New()
	w := a.NewWindow("EbiQuiz")
	w.Resize(fyne.NewSize(635, 475))
	w.SetMainMenu(fyne.NewMainMenu())

	text := canvas.NewText("Welcome To Ebitengine Quiz!!", color.NRGBA{255, 115, 22, 255})
	text.TextSize = 30
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Italic: false, Bold: true, Monospace: false}

	line := canvas.NewLine(color.NRGBA{255, 115, 22, 255})

	if !w.FullScreen() {
		line.StrokeWidth = 750
	} else {
		line.StrokeWidth = 1000
	}

	w.SetContent(MainPageUI(text, widget.NewButton("Start Quiz!", func() {

	}), line))

	w.Show()
	a.Run()
}

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainPageUI(label *widget.Label, button *widget.Button) *fyne.Container {
	return container.New(layout.NewBorderLayout(label, button, nil, nil),
		container.NewVBox(label),
		layout.NewSpacer(),
		container.NewVBox(button),
	)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(635, 475))

	input := widget.NewLabel("Hello welcome to ebitengine quiz game!")

	w.SetContent(MainPageUI(input, widget.NewButton("Click me!", func() {

	})))

	w.Show()
	a.Run()
}

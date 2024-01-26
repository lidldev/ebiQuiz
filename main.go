package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeUI(tabs *container.AppTabs) *fyne.Container {
	return container.New(layout.NewGridLayout(1),
		tabs)
}

func SettingsUI() *fyne.Container {
	return container.New(layout.NewGridLayout(1),
		widget.NewLabel("Settings Page"))
}

func MainPageUI(label *canvas.Text, image *canvas.Image, button *widget.Button) *fyne.Container {
	return container.New(layout.NewBorderLayout(label, button, nil, nil),
		label,
		image,
		button,
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
		line.StrokeWidth = 300
	} else {
		line.StrokeWidth = 1000
	}

	ebiten := canvas.NewImageFromFile("resource/ebiten.png")
	ebiten.FillMode = canvas.ImageFillContain

	button := widget.NewButton("Start Quiz!", func() {
		dialog.ShowConfirm("Start Quiz", "Are You Ready?", func(b bool) {}, w)

	})

	button.Importance = widget.WarningImportance

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home Page", theme.HomeIcon(), MainPageUI(text, ebiten, button)),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), SettingsUI()),
	)

	tabs.Items[0].Icon = theme.HomeIcon()
	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(makeUI(tabs))

	w.Show()
	a.Run()
}

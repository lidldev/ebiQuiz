package main

import (
	"fyne.io/fyne/v2/app"
)

func makeUI() {

}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	w.Show()
	a.Run()
}

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	theApp := app.New()
	win := theApp.NewWindow("Fynodoro")
	win.Resize(fyne.NewSize(400, 200))

	hello := widget.NewLabel("25:00")
	win.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Start", func() {
			hello.SetText("Starting...")
		}),
	))

	win.ShowAndRun()
}

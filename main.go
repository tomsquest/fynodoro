package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func main() {
	theApp := app.NewWithID("com.tomquest.fynodoro")
	theApp.Settings().SetTheme(&myTheme{})

	win := theApp.NewWindow("Fynodoro")
	win.SetIcon(resourceIconPng)
	win.Resize(fyne.NewSize(400, 200))
	win.CenterOnScreen()

	red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
	text := canvas.NewText("25:00", red)
	text.TextStyle.Monospace = true
	text.TextStyle.Bold = true
	text.TextSize = 100

	content := container.New(layout.NewMaxLayout(), text)
	win.SetContent(content)

	win.ShowAndRun()
}

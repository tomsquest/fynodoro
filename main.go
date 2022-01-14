package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/tomsquest/fynodoro/ui"
)

func main() {
	myApp := app.NewWithID("com.tomsquest.fynodoro")
	myApp.Settings().SetTheme(&ui.Theme{})
	myApp.SetIcon(ui.AssetIconPng)

	ui.Display(myApp)
}

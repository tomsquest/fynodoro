//go:generate fyne bundle --package ui -o ui/assets.go --prefix Asset assets/Icon.png
package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/tomsquest/fynodoro/ui"
)

// Variables set by goreleaser
// We need to put them in a Struct to be able to access them in the UI
var (
	version    string
	commit     string
	commitDate string
)

func main() {
	myApp := app.NewWithID("com.tomsquest.fynodoro")
	myApp.SetIcon(ui.AssetIconPng)

	ui.Display(myApp, ui.BuildInfo{version, commit, commitDate})
}

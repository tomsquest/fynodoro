package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/tomsquest/fynodoro/pomodoro"
	"github.com/tomsquest/fynodoro/pref"
	"github.com/tomsquest/fynodoro/ui"
	"time"
)

func main() {
	myApp := app.NewWithID("com.tomsquest.fynodoro")
	myApp.Settings().SetTheme(&ui.Theme{})
	myApp.SetIcon(ui.AssetIconPng)

	myPref := pref.Load()

	myPomodoro := pomodoro.NewPomodoro(&pomodoro.Params{
		WorkDuration:       time.Duration(myPref.WorkDuration) * time.Minute,
		ShortBreakDuration: time.Duration(myPref.ShortBreakDuration) * time.Minute,
		LongBreakDuration:  time.Duration(myPref.LongBreakDuration) * time.Minute,
		WorkRounds:         myPref.WorkRounds,
	})

	myWin := myApp.NewWindow("Fynodoro")
	myWin.CenterOnScreen()
	myWin.SetMaster()
	myWin.SetContent(ui.MakeClassicView(myApp, myPomodoro))
	myWin.ShowAndRun()
}

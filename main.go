package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/tomsquest/fynodoro/pomodoro"
	"github.com/tomsquest/fynodoro/ui"
)

func main() {
	myApp := app.NewWithID("com.tomsquest.fynodoro")
	myApp.Settings().SetTheme(&ui.Theme{})

	myPomodoro := pomodoro.NewPomodoroWithDefault()
	//myPomodoro := pomodoro.NewPomodoro(&pomodoro.Params{
	//	WorkRound:          2,
	//	WorkDuration:       6 * time.Second,
	//	ShortBreakDuration: 2 * time.Second,
	//	LongBreakDuration:  4 * time.Second,
	//})

	myWin := myApp.NewWindow("Fynodoro")
	myWin.SetIcon(resourceIconPng)
	myWin.CenterOnScreen()
	myWin.SetMaster()
	myWin.SetContent(ui.MakeClassicView(myPomodoro))
	myWin.ShowAndRun()
}

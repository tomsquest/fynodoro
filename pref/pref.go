package pref

import (
	"fyne.io/fyne/v2"
)

type Pref struct {
	WorkDuration       int
	ShortBreakDuration int
	LongBreakDuration  int
	WorkRounds         int
}

func Load() Pref {
	app := fyne.CurrentApp()
	workDuration := app.Preferences().IntWithFallback("workDuration", 25)
	shortBreakDuration := app.Preferences().IntWithFallback("shortBreakDuration", 5)
	longBreakDuration := app.Preferences().IntWithFallback("longBreakDuration", 15)
	workRounds := app.Preferences().IntWithFallback("workRounds", 4)

	return Pref{
		workDuration,
		shortBreakDuration,
		longBreakDuration,
		workRounds,
	}
}

func Save(pref Pref) {
	app := fyne.CurrentApp()
	app.Preferences().SetInt("workDuration", pref.WorkDuration)
	app.Preferences().SetInt("shortBreakDuration", pref.ShortBreakDuration)
	app.Preferences().SetInt("longBreakDuration", pref.LongBreakDuration)
	app.Preferences().SetInt("workRounds", pref.WorkRounds)
}

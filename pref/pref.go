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
	return Pref{
		WorkDuration:       app.Preferences().IntWithFallback("workDuration", 25),
		ShortBreakDuration: app.Preferences().IntWithFallback("shortBreakDuration", 5),
		LongBreakDuration:  app.Preferences().IntWithFallback("longBreakDuration", 15),
		WorkRounds:         app.Preferences().IntWithFallback("workRounds", 4),
	}
}

func Save(pref Pref) {
	app := fyne.CurrentApp()
	app.Preferences().SetInt("workDuration", pref.WorkDuration)
	app.Preferences().SetInt("shortBreakDuration", pref.ShortBreakDuration)
	app.Preferences().SetInt("longBreakDuration", pref.LongBreakDuration)
	app.Preferences().SetInt("workRounds", pref.WorkRounds)
}

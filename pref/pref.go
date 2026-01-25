package pref

import (
	"fyne.io/fyne/v2"
)

type Pref struct {
	WorkDuration            int
	ShortBreakDuration      int
	LongBreakDuration       int
	WorkRounds              int
	TimerFontSize           int
	TimerFontColor          string
	ShowButtons             bool
	StartMinimized          bool
	EnableNotificationPopup bool
	NotificationScript      string
}

func Load() Pref {
	app := fyne.CurrentApp()
	return Pref{
		WorkDuration:            app.Preferences().IntWithFallback("workDuration", 25),
		ShortBreakDuration:      app.Preferences().IntWithFallback("shortBreakDuration", 5),
		LongBreakDuration:       app.Preferences().IntWithFallback("longBreakDuration", 15),
		WorkRounds:              app.Preferences().IntWithFallback("workRounds", 4),
		TimerFontSize:           app.Preferences().IntWithFallback("timerFontSize", 60),
		TimerFontColor:          app.Preferences().StringWithFallback("timerFontColor", "#555555"),
		ShowButtons:             app.Preferences().BoolWithFallback("showButtons", true),
		StartMinimized:          app.Preferences().BoolWithFallback("startMinimized", false),
		EnableNotificationPopup: app.Preferences().BoolWithFallback("enableNotificationPopup", true),
		NotificationScript:      app.Preferences().StringWithFallback("notificationScript", "/usr/share/fynodoro/notify.sh"),
	}
}

func Save(pref Pref) {
	app := fyne.CurrentApp()
	app.Preferences().SetInt("workDuration", pref.WorkDuration)
	app.Preferences().SetInt("shortBreakDuration", pref.ShortBreakDuration)
	app.Preferences().SetInt("longBreakDuration", pref.LongBreakDuration)
	app.Preferences().SetInt("workRounds", pref.WorkRounds)
	app.Preferences().SetInt("timerFontSize", pref.TimerFontSize)
	app.Preferences().SetString("timerFontColor", pref.TimerFontColor)
	app.Preferences().SetBool("showButtons", pref.ShowButtons)
	app.Preferences().SetBool("startMinimized", pref.StartMinimized)
	app.Preferences().SetBool("enableNotificationPopup", pref.EnableNotificationPopup)
	app.Preferences().SetString("notificationScript", pref.NotificationScript)
}

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
	app.Preferences().SetBool("startMinimized", pref.StartMinimized)
	app.Preferences().SetBool("enableNotificationPopup", pref.EnableNotificationPopup)
	app.Preferences().SetString("notificationScript", pref.NotificationScript)
}

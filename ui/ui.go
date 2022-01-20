package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/tomsquest/fynodoro/pomodoro"
	"github.com/tomsquest/fynodoro/pref"
	"time"
)

func Display(app fyne.App) {
	myPref := pref.Load()
	myPomodoro := pomodoro.NewPomodoro(&pomodoro.Params{
		WorkDuration:       time.Duration(myPref.WorkDuration) * time.Minute,
		ShortBreakDuration: time.Duration(myPref.ShortBreakDuration) * time.Minute,
		LongBreakDuration:  time.Duration(myPref.LongBreakDuration) * time.Minute,
		WorkRounds:         myPref.WorkRounds,
	})

	myWin := app.NewWindow("Fynodoro")
	myWin.CenterOnScreen()
	myWin.SetMaster()
	myWin.SetContent(MakeClassicLayout(app, myPomodoro))
	myWin.ShowAndRun()
}

func MakeClassicLayout(app fyne.App, myPomodoro *pomodoro.Pomodoro) fyne.CanvasObject {
	timer := canvas.NewText(formatDuration(myPomodoro.RemainingTime), nil)
	timer.TextSize = 42
	timerButton := widget.NewButton("", nil)
	timerPanel := container.NewHBox(layout.NewSpacer(), container.NewMax(timer, timerButton), layout.NewSpacer())

	playButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)
	stopButton := widget.NewButtonWithIcon("", theme.MediaStopIcon(), nil)
	nextButton := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), nil)
	settingsButton := widget.NewButtonWithIcon("", theme.SettingsIcon(), nil)
	buttons := container.NewHBox(layout.NewSpacer(), playButton, stopButton, nextButton, settingsButton, layout.NewSpacer())
	stopButton.Disable()

	onPlay := func() {
		if myPomodoro.Running {
			playButton.Icon = theme.MediaPlayIcon()
			playButton.Refresh()

			myPomodoro.Pause()
		} else {
			playButton.Icon = theme.MediaPauseIcon()
			playButton.Refresh()

			myPomodoro.Start()
		}
		stopButton.Enable()
	}
	onStop := func() {
		myPomodoro.Stop()

		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
		playButton.Icon = theme.MediaPlayIcon()
		playButton.Refresh()
		stopButton.Disable()
	}
	onNext := func() {
		myPomodoro.Next()

		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
		playButton.Icon = theme.MediaPlayIcon()
		playButton.Refresh()
		stopButton.Disable()
	}
	onSettings := func() {
		onPrefUpdated := func(newPref pref.Pref) {
			myPomodoro.SetWorkDuration(time.Duration(newPref.WorkDuration) * time.Minute)
			myPomodoro.SetShortBreakDuration(time.Duration(newPref.ShortBreakDuration) * time.Minute)
			myPomodoro.SetLongBreakDuration(time.Duration(newPref.LongBreakDuration) * time.Minute)
			myPomodoro.SetWorkRounds(newPref.WorkRounds)
			myPomodoro.SetRemainingTime()

			timer.Text = formatDuration(myPomodoro.RemainingTime)
			timer.Refresh()
		}

		win := app.NewWindow("Settings")
		settings := MakeSettings(win, onPrefUpdated)
		win.SetContent(settings)

		settingsButton.Disable()
		win.SetOnClosed(func() {
			settingsButton.Enable()
		})

		win.CenterOnScreen()
		win.Show()
	}

	playButton.OnTapped = onPlay
	timerButton.OnTapped = onPlay
	stopButton.OnTapped = onStop
	nextButton.OnTapped = onNext
	settingsButton.OnTapped = onSettings

	myPomodoro.OnTick = func() {
		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
	}
	myPomodoro.OnEnd = func(kind pomodoro.Kind) {
		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
		playButton.Icon = theme.MediaPlayIcon()
		playButton.Refresh()
		stopButton.Disable()

		notifyPomodoroDone(kind)
	}

	return container.NewVBox(timerPanel, buttons)
}

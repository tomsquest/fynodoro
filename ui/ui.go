package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/tomsquest/fynodoro/pomodoro"
)

func MakeClassicView(myPomodoro *pomodoro.Pomodoro) *fyne.Container {
	timer := canvas.NewText(formatDuration(myPomodoro.RemainingTime), nil)
	timer.TextSize = 42
	timerButton := widget.NewButton("", nil)
	timerPanel := container.NewHBox(layout.NewSpacer(), container.NewMax(timer, timerButton), layout.NewSpacer())

	startButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)
	stopButton := widget.NewButtonWithIcon("", theme.MediaStopIcon(), nil)
	nextButton := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), nil)
	buttons := container.NewHBox(layout.NewSpacer(), startButton, stopButton, nextButton, layout.NewSpacer())

	onPlay := func() {
		if myPomodoro.Running {
			startButton.Icon = theme.MediaPlayIcon()
			startButton.Refresh()

			myPomodoro.Pause()
		} else {
			startButton.Icon = theme.MediaPauseIcon()
			startButton.Refresh()

			myPomodoro.Start()
		}
	}
	onStop := func() {
		myPomodoro.Stop()

		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()
	}
	onNext := func() {
		myPomodoro.Next()

		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()
	}

	startButton.OnTapped = onPlay
	timerButton.OnTapped = onPlay
	stopButton.OnTapped = onStop
	nextButton.OnTapped = onNext

	myPomodoro.OnTick = func() {
		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
	}
	myPomodoro.OnEnd = func(kind pomodoro.Kind) {
		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()

		notifyPomodoroDone(kind)
	}

	return container.NewVBox(timerPanel, buttons)
}

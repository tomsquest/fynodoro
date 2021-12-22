package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gen2brain/beeep"
	"github.com/tomsquest/fynodoro/pomodoro"
)

func main() {
	myApp := app.NewWithID("com.tomquest.fynodoro")
	myApp.Settings().SetTheme(&myTheme{})

	myWin := myApp.NewWindow("Fynodoro")
	myWin.SetIcon(resourceIconPng)
	myWin.CenterOnScreen()

	myPomodoro := pomodoro.NewPomodoroWithDefault()

	timer := canvas.NewText(formatDuration(myPomodoro.Remaining), nil)
	timer.TextSize = 42
	timerPanel := container.NewHBox(layout.NewSpacer(), timer, layout.NewSpacer())

	startButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)
	startButton.OnTapped = func() {
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
	stopButton := widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() {
		myPomodoro.Stop()

		timer.Text = formatDuration(myPomodoro.Remaining)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()
	})
	skipButton := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() {
		myPomodoro.Next()

		timer.Text = formatDuration(myPomodoro.Remaining)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()
	})
	buttons := container.NewHBox(layout.NewSpacer(), startButton, stopButton, skipButton, layout.NewSpacer())

	myPomodoro.OnTick = func() {
		timer.Text = formatDuration(myPomodoro.Remaining)
		timer.Refresh()
	}
	myPomodoro.OnEnd = func(kind pomodoro.Kind) {
		timer.Text = formatDuration(myPomodoro.Remaining)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()

		notifTitle := fmt.Sprintf("%s done", kind)
		notifMessage := fmt.Sprintf("You just finished a %s pomodoro.", kind)
		_ = beeep.Notify(notifTitle, notifMessage, "")
	}

	myWin.SetContent(container.NewBorder(nil, buttons, nil, nil, timerPanel))
	myWin.ShowAndRun()
}

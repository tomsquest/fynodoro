package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	myApp := app.NewWithID("com.tomquest.fynodoro")
	myApp.Settings().SetTheme(&myTheme{})

	myWin := myApp.NewWindow("Fynodoro")
	myWin.SetIcon(resourceIconPng)
	myWin.CenterOnScreen()

	workDuration := 25 * 60 * time.Second
	shortBreakDuration := 5 * 60 * time.Second
	pomodoro := NewPomodoro(workDuration, shortBreakDuration)

	timer := canvas.NewText(formatDuration(pomodoro.remaining), nil)
	timer.TextSize = 42
	timerPanel := container.NewHBox(layout.NewSpacer(), timer, layout.NewSpacer())

	startButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)
	startButton.OnTapped = func() {
		if pomodoro.running {
			startButton.Icon = theme.MediaPlayIcon()
			startButton.Refresh()

			pomodoro.Pause()
		} else {
			startButton.Icon = theme.MediaPauseIcon()
			startButton.Refresh()

			pomodoro.Start()
		}
	}
	stopButton := widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() {
		pomodoro.Stop()

		timer.Text = formatDuration(pomodoro.remaining)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()
	})
	skipButton := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() {
		pomodoro.Next()

		timer.Text = formatDuration(pomodoro.remaining)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()
	})
	buttons := container.NewHBox(layout.NewSpacer(), startButton, stopButton, skipButton, layout.NewSpacer())

	pomodoro.onTick = func() {
		timer.Text = formatDuration(pomodoro.remaining)
		timer.Refresh()
	}
	pomodoro.onEnd = func(kind PomodoroKind) {
		fmt.Println("onEnd")

		timer.Text = formatDuration(pomodoro.remaining)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()

		notification := fyne.NewNotification(kind.String()+" done", "You just finished a "+kind.String()+" pomodoro.")
		myApp.SendNotification(notification)
	}

	myWin.SetContent(container.NewBorder(nil, buttons, nil, nil, timerPanel))
	myWin.ShowAndRun()
}

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/tomsquest/fynodoro/pomodoro"
	"github.com/tomsquest/fynodoro/pref"
	"time"
)

type BuildInfo struct {
	Version    string
	Commit     string
	CommitDate string
}

func Display(app fyne.App, buildInfo BuildInfo, cliStartMinimized bool) {
	myPref := pref.Load()
	thePomodoro := pomodoro.NewPomodoro(&pomodoro.Params{
		WorkDuration:       time.Duration(myPref.WorkDuration) * time.Minute,
		ShortBreakDuration: time.Duration(myPref.ShortBreakDuration) * time.Minute,
		LongBreakDuration:  time.Duration(myPref.LongBreakDuration) * time.Minute,
		WorkRounds:         myPref.WorkRounds,
	})

	mainWindow := app.NewWindow("Fynodoro")
	mainWindow.SetMaster()
	mainWindow.SetContent(MakeClassicLayout(app, mainWindow, thePomodoro))
	mainWindow.SetCloseIntercept(mainWindow.Hide)
	mainWindow.SetFixedSize(true)

	if desk, ok := app.(desktop.App); ok {
		aboutWindow := makeAboutWindow(app, buildInfo)
		trayMenu := fyne.NewMenu("Fynodoro",
			fyne.NewMenuItem("Show", mainWindow.Show),
			fyne.NewMenuItem("Hide", mainWindow.Hide),
			fyne.NewMenuItem("Center", mainWindow.CenterOnScreen),
			fyne.NewMenuItem("About", aboutWindow.Show),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Quit", func() { app.Quit() }))
		desk.SetSystemTrayMenu(trayMenu)
	}

	if cliStartMinimized || myPref.StartMinimized {
		app.Run()
	} else {
		mainWindow.ShowAndRun()
	}
}

func makeAboutWindow(app fyne.App, buildInfo BuildInfo) fyne.Window {
	aboutWindow := app.NewWindow("About Fynodoro")
	aboutWindow.SetFixedSize(true)

	img := canvas.NewImageFromResource(AssetIconPng)
	img.SetMinSize(fyne.NewSquareSize(64))
	imgContainer := container.NewHBox(img, layout.NewSpacer())

	markdownStr := "# Fynodoro" + "\n"
	markdownStr += "" + "\n"
	markdownStr += "Fynodoro is a tiny and cute Pomodoro Widget" + "\n"
	markdownStr += "" + "\n"
	markdownStr += "- `Version:     " + buildInfo.Version + "`" + "\n"
	markdownStr += "- `Commit date: " + buildInfo.CommitDate + " `" + "\n"
	markdownStr += "- `Commit:      " + buildInfo.Commit + "`" + "\n"
	markdown := widget.NewRichTextFromMarkdown(markdownStr)

	closeButton := &widget.Button{
		Text:     "Close",
		OnTapped: aboutWindow.Hide,
	}
	buttonsContainer := container.NewHBox(layout.NewSpacer(), closeButton)

	aboutWindow.SetContent(container.NewVBox(imgContainer, markdown, layout.NewSpacer(), buttonsContainer))
	return aboutWindow
}

func MakeClassicLayout(app fyne.App, mainWindow fyne.Window, thePomodoro *pomodoro.Pomodoro) fyne.CanvasObject {
	timer := NewTappableText(formatDuration(thePomodoro.RemainingTime), nil, nil)
	timer.Label.TextSize = 60
	timer.Label.TextStyle.Bold = true
	timer.Label.Alignment = fyne.TextAlignCenter
	timerPanel := container.NewCenter(container.NewHBox(timer))

	playButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)
	stopButton := widget.NewButtonWithIcon("", theme.MediaStopIcon(), nil)
	nextButton := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), nil)
	settingsButton := widget.NewButtonWithIcon("", theme.SettingsIcon(), nil)
	buttons := container.NewCenter(container.NewHBox(playButton, stopButton, nextButton, settingsButton))

	timer.OnTapped = func() {
		playPausePomodoro(thePomodoro, playButton)
	}
	timer.OnTappedSecondary = func(pe *fyne.PointEvent) {
		// Create menu items with dynamic play/pause label
		var playPauseLabel string
		var playPauseIcon fyne.Resource
		if thePomodoro.Running {
			playPauseLabel = "Pause"
			playPauseIcon = theme.MediaPauseIcon()
		} else {
			playPauseLabel = "Play"
			playPauseIcon = theme.MediaPlayIcon()
		}

		playPauseItem := fyne.NewMenuItem(playPauseLabel, func() {
			playPausePomodoro(thePomodoro, playButton)
		})
		playPauseItem.Icon = playPauseIcon

		stopItem := fyne.NewMenuItem("Stop", func() {
			stopPomodoro(thePomodoro, playButton, timer)
		})
		stopItem.Icon = theme.MediaStopIcon()

		nextItem := fyne.NewMenuItem("Next", func() {
			nextPomodoro(thePomodoro, playButton, timer)
		})
		nextItem.Icon = theme.MediaSkipNextIcon()

		settingsItem := fyne.NewMenuItem("Settings", func() {
			settings := NewSettings()
			settings.SetOnSubmit(func() {
				// Apply new preferences to current pomodoro
				newPref := pref.Load()
				thePomodoro.SetWorkDuration(time.Duration(newPref.WorkDuration) * time.Minute)
				thePomodoro.SetShortBreakDuration(time.Duration(newPref.ShortBreakDuration) * time.Minute)
				thePomodoro.SetLongBreakDuration(time.Duration(newPref.LongBreakDuration) * time.Minute)
				thePomodoro.SetWorkRounds(newPref.WorkRounds)
				thePomodoro.SetRemainingTime()

				// Display new duration
				setTimerRemainingTime(thePomodoro, timer)
			})
			settings.SetOnClosed(func() {
				settingsButton.Enable()
			})

			settingsButton.Disable()
			settings.Show()
		})
		settingsItem.Icon = theme.SettingsIcon()

		closeItem := fyne.NewMenuItem("Close", func() {
			mainWindow.Hide()
		})
		closeItem.Icon = theme.WindowCloseIcon()

		quitItem := fyne.NewMenuItem("Quit", func() {
			app.Quit()
		})
		quitItem.Icon = theme.CancelIcon()

		// Create and show popup menu
		menu := fyne.NewMenu("",
			playPauseItem,
			stopItem,
			nextItem,
			fyne.NewMenuItemSeparator(),
			settingsItem,
			fyne.NewMenuItemSeparator(),
			closeItem,
			quitItem,
		)

		widget.ShowPopUpMenuAtPosition(menu, mainWindow.Canvas(), pe.AbsolutePosition)
	}
	playButton.OnTapped = func() {
		playPausePomodoro(thePomodoro, playButton)
	}
	stopButton.OnTapped = func() {
		stopPomodoro(thePomodoro, playButton, timer)
	}
	nextButton.OnTapped = func() {
		nextPomodoro(thePomodoro, playButton, timer)
	}
	settingsButton.OnTapped = func() {
		settings := NewSettings()
		settings.SetOnSubmit(func() {
			// Apply new preferences to current pomodoro
			newPref := pref.Load()
			thePomodoro.SetWorkDuration(time.Duration(newPref.WorkDuration) * time.Minute)
			thePomodoro.SetShortBreakDuration(time.Duration(newPref.ShortBreakDuration) * time.Minute)
			thePomodoro.SetLongBreakDuration(time.Duration(newPref.LongBreakDuration) * time.Minute)
			thePomodoro.SetWorkRounds(newPref.WorkRounds)
			thePomodoro.SetRemainingTime()

			// Display new duration
			setTimerRemainingTime(thePomodoro, timer)
		})
		settings.SetOnClosed(func() {
			settingsButton.Enable()
		})

		settingsButton.Disable()
		settings.Show()
	}

	thePomodoro.OnTick = func() {
		setTimerRemainingTime(thePomodoro, timer)
	}
	thePomodoro.OnEnd = func(kind pomodoro.Kind) {
		setTimerRemainingTime(thePomodoro, timer)

		playButton.Icon = theme.MediaPlayIcon()
		playButton.Refresh()

		notifyPomodoroDone(kind)
	}

	return container.NewVBox(timerPanel, buttons)
}

func setTimerRemainingTime(thePomodoro *pomodoro.Pomodoro, timer *TappableText) {
	timer.SetText(formatDuration(thePomodoro.RemainingTime))
	timer.Refresh()
}

func playPausePomodoro(thePomodoro *pomodoro.Pomodoro, playButton *widget.Button) {
	if thePomodoro.Running {
		thePomodoro.Pause()
		playButton.Icon = theme.MediaPlayIcon()
	} else {
		thePomodoro.Start()
		playButton.Icon = theme.MediaPauseIcon()
	}
	playButton.Refresh()
}

func stopPomodoro(thePomodoro *pomodoro.Pomodoro, playButton *widget.Button, timer *TappableText) {
	thePomodoro.Stop()
	playButton.Icon = theme.MediaPlayIcon()
	playButton.Refresh()
	setTimerRemainingTime(thePomodoro, timer)
}

func nextPomodoro(thePomodoro *pomodoro.Pomodoro, playButton *widget.Button, timer *TappableText) {
	thePomodoro.Next()
	playButton.Icon = theme.MediaPlayIcon()
	playButton.Refresh()
	setTimerRemainingTime(thePomodoro, timer)
}

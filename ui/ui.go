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

func Display(app fyne.App, buildInfo BuildInfo) {
	myPref := pref.Load()
	myPomodoro := pomodoro.NewPomodoro(&pomodoro.Params{
		WorkDuration:       time.Duration(myPref.WorkDuration) * time.Minute,
		ShortBreakDuration: time.Duration(myPref.ShortBreakDuration) * time.Minute,
		LongBreakDuration:  time.Duration(myPref.LongBreakDuration) * time.Minute,
		WorkRounds:         myPref.WorkRounds,
	})

	mainWindow := app.NewWindow("Fynodoro")
	mainWindow.SetMaster()
	mainWindow.SetContent(MakeClassicLayout(myPomodoro))
	mainWindow.SetCloseIntercept(mainWindow.Hide)
	mainWindow.SetFixedSize(true)

	if desk, ok := app.(desktop.App); ok {
		aboutWindow := makeAboutWindow(app, buildInfo)
		trayMenu := fyne.NewMenu("Fynodoro",
			fyne.NewMenuItem("Show", mainWindow.Show),
			fyne.NewMenuItem("Hide", mainWindow.Hide),
			fyne.NewMenuItem("Center", mainWindow.CenterOnScreen),
			fyne.NewMenuItem("About", aboutWindow.Show))
		desk.SetSystemTrayMenu(trayMenu)
	}
	mainWindow.ShowAndRun()
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

func MakeClassicLayout(myPomodoro *pomodoro.Pomodoro) fyne.CanvasObject {
	timer := canvas.NewText(formatDuration(myPomodoro.RemainingTime), nil)
	timer.TextSize = 42
	timerPanel := container.NewHBox(layout.NewSpacer(), container.NewStack(timer), layout.NewSpacer())

	playButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)
	stopButton := widget.NewButtonWithIcon("", theme.MediaStopIcon(), nil)
	nextButton := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), nil)
	settingsButton := widget.NewButtonWithIcon("", theme.SettingsIcon(), nil)
	buttons := container.NewHBox(layout.NewSpacer(), playButton, stopButton, nextButton, settingsButton, layout.NewSpacer())

	onPlay := func() {
		if myPomodoro.Running {
			myPomodoro.Pause()

			playButton.Icon = theme.MediaPlayIcon()
			playButton.Refresh()
		} else {
			myPomodoro.Start()

			playButton.Icon = theme.MediaPauseIcon()
			playButton.Refresh()
		}
	}
	onStop := func() {
		myPomodoro.Stop()

		playButton.Icon = theme.MediaPlayIcon()
		playButton.Refresh()

		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
	}
	onNext := func() {
		myPomodoro.Next()

		playButton.Icon = theme.MediaPlayIcon()
		playButton.Refresh()

		timer.Text = formatDuration(myPomodoro.RemainingTime)
		timer.Refresh()
	}
	onSettings := func() {
		settings := NewSettings()
		settings.SetOnSubmit(func() {
			// Apply new preferences to current pomodoro
			newPref := pref.Load()
			myPomodoro.SetWorkDuration(time.Duration(newPref.WorkDuration) * time.Minute)
			myPomodoro.SetShortBreakDuration(time.Duration(newPref.ShortBreakDuration) * time.Minute)
			myPomodoro.SetLongBreakDuration(time.Duration(newPref.LongBreakDuration) * time.Minute)
			myPomodoro.SetWorkRounds(newPref.WorkRounds)
			myPomodoro.SetRemainingTime()

			// Display new duration
			timer.Text = formatDuration(myPomodoro.RemainingTime)
			timer.Refresh()
		})
		settings.SetOnClosed(func() {
			settingsButton.Enable()
		})

		settingsButton.Disable()
		settings.Show()
	}

	playButton.OnTapped = onPlay
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

		notifyPomodoroDone(kind)
	}

	return container.NewVBox(timerPanel, buttons)
}

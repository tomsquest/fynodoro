package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/tomsquest/fynodoro/pomodoro"
	"github.com/tomsquest/fynodoro/pref"
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
	pomodoroWidget := NewPomodoroWidget(thePomodoro)
	mainWindow.SetContent(pomodoroWidget)
	mainWindow.SetCloseIntercept(mainWindow.Hide)

	if desk, ok := app.(desktop.App); ok {
		aboutWindow := makeAboutWindow(app, buildInfo)
		trayMenu := fyne.NewMenu("Fynodoro",
			fyne.NewMenuItem("Play/Pause", pomodoroWidget.PlayPause),
			fyne.NewMenuItem("Stop", pomodoroWidget.Stop),
			fyne.NewMenuItem("Next", pomodoroWidget.Next),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Show", mainWindow.Show),
			fyne.NewMenuItem("Hide", mainWindow.Hide),
			fyne.NewMenuItem("Center", mainWindow.CenterOnScreen),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Preferences", func() {
				settings := NewSettings()
				settings.SetOnSubmit(func() {
					applyPreferencesToPomodoro(thePomodoro)
					pomodoroWidget.ApplyPreferences()
				})
				settings.Show()
			}),
			fyne.NewMenuItem("About", aboutWindow.Show),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Quit", app.Quit))
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

func applyPreferencesToPomodoro(p *pomodoro.Pomodoro) {
	newPref := pref.Load()
	p.SetWorkDuration(time.Duration(newPref.WorkDuration) * time.Minute)
	p.SetShortBreakDuration(time.Duration(newPref.ShortBreakDuration) * time.Minute)
	p.SetLongBreakDuration(time.Duration(newPref.LongBreakDuration) * time.Minute)
	p.SetWorkRounds(newPref.WorkRounds)
	p.SetRemainingTime()
}

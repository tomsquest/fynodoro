package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

type PomodoroKind int

const (
	Work PomodoroKind = iota
	ShortBreak
)

func (t PomodoroKind) String() string {
	names := [...]string{"Work", "Short Break"}
	return names[t]
}

type pomodoro struct {
	// Config
	workDuration       time.Duration
	shortBreakDuration time.Duration
	// External events
	onTick func()
	onEnd  func()
	// State
	kind      PomodoroKind
	remaining time.Duration
	running   bool
	// Internal
	ticker *time.Ticker
	timer  *time.Timer
}

func NewPomodoro(workDuration time.Duration, shortBreakDuration time.Duration) *pomodoro {
	p := new(pomodoro)
	p.workDuration = workDuration
	p.shortBreakDuration = shortBreakDuration

	p.kind = Work
	p.remaining = workDuration
	p.running = false
	return p
}
func (p *pomodoro) Start() {
	fmt.Println("Start", "Remaining:", p.remaining, "PomodoroKind:", p.kind)
	p.ticker = time.NewTicker(time.Second)
	p.timer = time.NewTimer(p.remaining)
	p.running = true

	go func() {
		for {
			select {
			case <-p.ticker.C:
				p.remaining -= time.Second

				if p.onTick != nil {
					p.onTick()
				}
			case <-p.timer.C:
				p.remaining = 0
				p.stop()
				p.next()

				if p.onEnd != nil {
					p.onEnd()
				}

				return
			}
		}
	}()
}
func (p *pomodoro) Pause() {
	fmt.Println("Pause", "Remaining:", p.remaining, "PomodoroKind:", p.kind)
	p.stop()
}
func (p *pomodoro) Stop() {
	fmt.Println("Stop", "Remaining:", p.remaining, "PomodoroKind:", p.kind)
	p.stop()

	switch p.kind {
	case ShortBreak:
		p.remaining = p.shortBreakDuration
	default:
		p.remaining = p.workDuration
	}
}
func (p *pomodoro) stop() {
	p.running = false

	if p.ticker != nil {
		p.ticker.Stop()
	}
	if p.timer != nil {
		p.timer.Stop()
	}
}
func (p *pomodoro) next() {
	switch p.kind {
	case ShortBreak:
		p.kind = Work
		p.remaining = p.workDuration
	default:
		p.kind = ShortBreak
		p.remaining = p.shortBreakDuration
	}
}

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
	buttons := container.NewHBox(layout.NewSpacer(), startButton, stopButton, layout.NewSpacer())

	pomodoro.onTick = func() {
		timer.Text = formatDuration(pomodoro.remaining)
		timer.Refresh()
	}
	pomodoro.onEnd = func() {
		fmt.Println("onEnd")

		timer.Text = formatDuration(pomodoro.remaining)
		timer.Refresh()
		startButton.Icon = theme.MediaPlayIcon()
		startButton.Refresh()
	}

	myWin.SetContent(container.NewVBox(timerPanel, buttons))
	myWin.ShowAndRun()
}

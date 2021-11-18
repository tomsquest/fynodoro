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

type pomodoro struct {
	period    time.Duration
	remaining time.Duration
	running   bool
	ticker    *time.Ticker
	timer     *time.Timer
	onTick    func(time.Duration)
	onEnd     func()
}

func NewPomodoro(period time.Duration) *pomodoro {
	p := new(pomodoro)
	p.period = period
	p.remaining = period
	p.running = false
	return p
}
func (p *pomodoro) Start() {
	fmt.Println("Start", p.remaining)
	p.ticker = time.NewTicker(time.Second)
	p.timer = time.NewTimer(p.remaining)
	p.running = true

	go func() {
		for {
			select {
			case <-p.ticker.C:
				p.remaining -= time.Second

				if p.onTick != nil {
					p.onTick(p.remaining)
				}
			case <-p.timer.C:
				p.stop()

				if p.onEnd != nil {
					p.onEnd()
				}

				return
			}
		}
	}()
}
func (p *pomodoro) Pause() {
	fmt.Println("Pause", p.remaining)
	p.stop()
}
func (p *pomodoro) Stop() {
	fmt.Println("Stop")
	p.stop()
	p.remaining = p.period
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

func main() {
	myApp := app.NewWithID("com.tomquest.fynodoro")
	myApp.Settings().SetTheme(&myTheme{})

	myWin := myApp.NewWindow("Fynodoro")
	myWin.SetIcon(resourceIconPng)
	myWin.CenterOnScreen()

	timerDuration := 25 * 60 * time.Second
	pomodoro := NewPomodoro(timerDuration)

	timer := canvas.NewText(formatDuration(pomodoro.period), nil)
	timer.TextSize = 72

	startButton := widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), nil)
	startButton.OnTapped = func() {
		if pomodoro.running {
			fmt.Println("Pause")
			startButton.Icon = theme.MediaPlayIcon()
			startButton.Text = "Start"
			startButton.Refresh()

			pomodoro.Pause()
		} else {
			fmt.Println("Start")
			startButton.Icon = theme.MediaPauseIcon()
			startButton.Text = "Pause"
			startButton.Refresh()

			pomodoro.Start()
		}
	}
	stopButton := widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), func() {
		pomodoro.Stop()
		timer.Text = formatDuration(pomodoro.period)
		timer.Refresh()

		startButton.Icon = theme.MediaPlayIcon()
		startButton.Text = "Start"
		startButton.Refresh()
	})
	buttons := container.NewHBox(layout.NewSpacer(), startButton, stopButton, layout.NewSpacer())

	pomodoro.onTick = func(remainingTime time.Duration) {
		fmt.Println("onTick", remainingTime)
		timer.Text = formatDuration(remainingTime)
		timer.Refresh()
	}
	pomodoro.onEnd = func() {
		fmt.Println("onEnd")

		startButton.Icon = theme.MediaPlayIcon()
		startButton.Text = "Start"
		startButton.Refresh()
	}

	myWin.SetContent(container.NewVBox(timer, buttons))
	myWin.ShowAndRun()
}

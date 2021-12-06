package main

import (
	"fmt"
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
	onEnd  func(kind PomodoroKind)
	// State
	kind      PomodoroKind
	remaining time.Duration
	running   bool
	// Internal
	ticker *time.Ticker
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
	p.running = true

	go func() {
		for {
			select {
			case <-p.ticker.C:
				p.remaining -= time.Second

				if p.remaining > 0 {
					if p.onTick != nil {
						p.onTick()
					}
				} else {
					currentKind := p.kind

					p.remaining = 0
					p.stop()
					p.next()

					if p.onEnd != nil {
						p.onEnd(currentKind)
					}
				}
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
func (p *pomodoro) Next() {
	fmt.Println("Next", "Remaining:", p.remaining, "PomodoroKind:", p.kind)
	p.stop()
	p.next()
}
func (p *pomodoro) stop() {
	p.running = false

	if p.ticker != nil {
		p.ticker.Stop()
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

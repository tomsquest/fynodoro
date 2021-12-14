package pomodoro

import (
	"time"
)

type Kind int

const (
	Work Kind = iota
	ShortBreak
)

func (t Kind) String() string {
	names := [...]string{"Work", "Short Break"}
	return names[t]
}

type pomodoro struct {
	// External events
	OnTick func()
	OnEnd  func(kind Kind)
	// State
	Kind      Kind
	Remaining time.Duration
	Running   bool
	// Config
	workDuration       time.Duration
	shortBreakDuration time.Duration
	// Internal
	ticker *time.Ticker
}

func NewPomodoro(workDuration time.Duration, shortBreakDuration time.Duration) *pomodoro {
	p := new(pomodoro)
	p.workDuration = workDuration
	p.shortBreakDuration = shortBreakDuration

	p.Kind = Work
	p.Remaining = workDuration
	p.Running = false
	return p
}

func (p *pomodoro) Start() {
	p.ticker = time.NewTicker(time.Second)
	p.Running = true

	go func() {
		for {
			select {
			case <-p.ticker.C:
				p.Remaining -= time.Second

				if p.Remaining > 0 {
					if p.OnTick != nil {
						p.OnTick()
					}
				} else {
					currentKind := p.Kind

					p.Remaining = 0
					p.stop()
					p.next()

					if p.OnEnd != nil {
						p.OnEnd(currentKind)
					}
				}
			}
		}
	}()
}

func (p *pomodoro) Pause() {
	p.stop()
}

func (p *pomodoro) Stop() {
	p.stop()

	switch p.Kind {
	case ShortBreak:
		p.Remaining = p.shortBreakDuration
	default:
		p.Remaining = p.workDuration
	}
}

func (p *pomodoro) Next() {
	p.stop()
	p.next()
}

func (p *pomodoro) stop() {
	p.Running = false

	if p.ticker != nil {
		p.ticker.Stop()
	}
}

func (p *pomodoro) next() {
	switch p.Kind {
	case ShortBreak:
		p.Kind = Work
		p.Remaining = p.workDuration
	default:
		p.Kind = ShortBreak
		p.Remaining = p.shortBreakDuration
	}
}

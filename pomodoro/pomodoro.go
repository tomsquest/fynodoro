package pomodoro

import (
	"fmt"
	"github.com/benbjohnson/clock"
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

type PomodoroParams struct {
	WorkDuration       time.Duration
	ShortBreakDuration time.Duration
	Clock              clock.Clock
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
	clock              clock.Clock
	// Internal
	ticker *clock.Ticker
}

func NewPomodoro(params *PomodoroParams) *pomodoro {
	p := &pomodoro{
		workDuration:       params.WorkDuration,
		shortBreakDuration: params.ShortBreakDuration,
		Kind:               Work,
		Remaining:          params.WorkDuration,
		Running:            false,
	}
	if params.Clock == nil {
		p.clock = clock.New()
	} else {
		p.clock = params.Clock
	}
	return p
}

func (p *pomodoro) Start() {
	fmt.Println("Start", p.Remaining, "Kind", p.Kind)
	p.ticker = p.clock.Ticker(time.Second)
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
	fmt.Println("Pause", p.Remaining, "Kind", p.Kind)
	p.stop()
}

func (p *pomodoro) Stop() {
	fmt.Println("Stop", p.Remaining, "Kind", p.Kind)
	p.stop()

	switch p.Kind {
	case ShortBreak:
		p.Remaining = p.shortBreakDuration
	default:
		p.Remaining = p.workDuration
	}
}

func (p *pomodoro) Next() {
	fmt.Println("Next", p.Remaining, "Kind", p.Kind)
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

package pomodoro

import (
	"github.com/benbjohnson/clock"
	"time"
)

type Kind int

const (
	Work Kind = iota
	ShortBreak
	LongBreak
)

func (t Kind) String() string {
	names := [...]string{"Work", "Short Break", "Long Break"}
	return names[t]
}

type Params struct {
	WorkDuration       time.Duration
	ShortBreakDuration time.Duration
	LongBreakDuration  time.Duration
	WorkRound          uint8
	// use for testing, do not set it yourself
	Clock clock.Clock
}

type pomodoro struct {
	// Config
	workDuration       time.Duration
	shortBreakDuration time.Duration
	longBreakDuration  time.Duration
	workRound          uint8
	clock              clock.Clock
	// External events
	OnTick func()
	OnEnd  func(kind Kind)
	// State
	Kind           Kind
	RemainingTime  time.Duration
	RemainingRound uint8
	Running        bool
	// Internal
	ticker *clock.Ticker
}

func NewPomodoroWithDefault() *pomodoro {
	p := &pomodoro{
		workDuration:       25 * time.Minute,
		shortBreakDuration: 5 * time.Minute,
		longBreakDuration:  15 * time.Minute,
		workRound:          4,
		Kind:               Work,
		Running:            false,
	}
	p.RemainingTime = p.workDuration
	p.RemainingRound = p.workRound
	p.clock = clock.New()
	return p
}

func NewPomodoro(params *Params) *pomodoro {
	p := NewPomodoroWithDefault()
	if params.WorkDuration > 0 {
		p.workDuration = params.WorkDuration
	}
	if params.ShortBreakDuration > 0 {
		p.shortBreakDuration = params.ShortBreakDuration
	}
	if params.LongBreakDuration > 0 {
		p.longBreakDuration = params.LongBreakDuration
	}
	if params.WorkRound > 0 {
		p.workRound = params.WorkRound
	}
	if params.Clock != nil {
		p.clock = params.Clock
	}
	p.RemainingTime = p.workDuration
	p.RemainingRound = p.workRound
	return p
}

func (p *pomodoro) Start() {
	p.ticker = p.clock.Ticker(time.Second)
	p.Running = true

	go func() {
		for {
			select {
			case <-p.ticker.C:
				p.RemainingTime -= time.Second

				if p.RemainingTime > 0 {
					if p.OnTick != nil {
						p.OnTick()
					}
				} else {
					endedKind := p.Kind

					p.stop()
					p.next()

					if p.OnEnd != nil {
						p.OnEnd(endedKind)
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
		p.RemainingTime = p.shortBreakDuration
	case LongBreak:
		p.RemainingTime = p.longBreakDuration
	default:
		p.RemainingTime = p.workDuration
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
	case ShortBreak, LongBreak:
		p.Kind = Work
		p.RemainingTime = p.workDuration
	default:
		p.RemainingRound--
		if p.RemainingRound == 0 {
			p.Kind = LongBreak
			p.RemainingTime = p.longBreakDuration
			p.RemainingRound = p.workRound
		} else {
			p.Kind = ShortBreak
			p.RemainingTime = p.shortBreakDuration
		}
	}
}

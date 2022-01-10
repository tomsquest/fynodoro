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
	WorkRounds         int
	// use for testing, do not set it yourself
	Clock clock.Clock
}

type Pomodoro struct {
	// Config
	workDuration       time.Duration
	shortBreakDuration time.Duration
	longBreakDuration  time.Duration
	workRounds         int
	clock              clock.Clock
	// External events
	OnTick func()
	OnEnd  func(kind Kind)
	// State
	Kind           Kind
	RemainingTime  time.Duration
	RemainingRound int
	Running        bool
	// Internal
	ticker *clock.Ticker
}

func NewPomodoro(params *Params) *Pomodoro {
	p := &Pomodoro{}
	p.workDuration = params.WorkDuration
	p.shortBreakDuration = params.ShortBreakDuration
	p.longBreakDuration = params.LongBreakDuration
	p.workRounds = params.WorkRounds
	p.RemainingTime = p.workDuration
	p.RemainingRound = p.workRounds
	p.Kind = Work
	p.Running = false
	if params.Clock != nil {
		p.clock = params.Clock
	} else {
		p.clock = clock.New()
	}
	return p
}

func (p *Pomodoro) Start() {
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

func (p *Pomodoro) Pause() {
	p.stop()
}

func (p *Pomodoro) Stop() {
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

func (p *Pomodoro) Next() {
	p.stop()
	p.next()
}

func (p *Pomodoro) stop() {
	p.Running = false

	if p.ticker != nil {
		p.ticker.Stop()
	}
}

func (p *Pomodoro) next() {
	switch p.Kind {
	case ShortBreak, LongBreak:
		p.Kind = Work
		p.RemainingTime = p.workDuration
	default:
		if p.workRounds == 0 || p.longBreakDuration == 0 {
			// LongBreak disabled, only do ShortBreak
			p.Kind = ShortBreak
			p.RemainingTime = p.shortBreakDuration
		} else {
			p.RemainingRound--
			if p.RemainingRound == 0 {
				p.Kind = LongBreak
				p.RemainingTime = p.longBreakDuration
				p.RemainingRound = p.workRounds
			} else {
				p.Kind = ShortBreak
				p.RemainingTime = p.shortBreakDuration
			}
		}
	}
}

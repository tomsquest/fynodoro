package pomodoro

import (
	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPomodoro(t *testing.T) {
	p := NewPomodoro(&Params{
		WorkDuration:       1 * time.Second,
		ShortBreakDuration: 2 * time.Minute,
		LongBreakDuration:  3 * time.Hour,
		WorkRounds:         42,
	})

	assert.Equal(t, 1*time.Second, p.workDuration)
	assert.Equal(t, 2*time.Minute, p.shortBreakDuration)
	assert.Equal(t, 3*time.Hour, p.longBreakDuration)
	assert.Equal(t, 42, p.workRounds)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 1*time.Second, p.RemainingTime)
	assert.False(t, p.Running)
}

func TestNewPomodoro_disableLongBreak_zeroDurationLongBreak(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		LongBreakDuration: 0,

		WorkDuration:       1,
		ShortBreakDuration: 1,
		WorkRounds:         1,
		Clock:              clockMock,
	})

	for i := 0; i < 99; i++ {
		assert.NotEqual(t, LongBreak, p.Kind)
		p.Next()
	}
}

func TestNewPomodoro_disableLongBreak_zeroWorkRounds(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		WorkRounds: 0,

		WorkDuration:       1,
		ShortBreakDuration: 1,
		LongBreakDuration:  1,
		Clock:              clockMock,
	})

	for i := 0; i < 99; i++ {
		assert.NotEqual(t, LongBreak, p.Kind)
		p.Next()
	}
}

func TestPomodoro_OnTick(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		WorkDuration:       25 * time.Second,
		ShortBreakDuration: 5 * time.Second,
		LongBreakDuration:  10 * time.Second,
		WorkRounds:         2,
		Clock:              clockMock,
	})

	// Capture tick events
	tickedCount := 0
	p.OnTick = func() {
		tickedCount++
	}

	// Start Work
	p.Start()

	clockMock.Add(1 * time.Second)
	assert.Equal(t, 1, tickedCount)
	clockMock.Add(1 * time.Second)
	assert.Equal(t, 2, tickedCount)

	// Start ShortBreak
	p.Next()
	assert.Equal(t, ShortBreak, p.Kind)
	tickedCount = 0 // Reset count

	p.Start()

	clockMock.Add(1 * time.Second)
	assert.Equal(t, 1, tickedCount)
	clockMock.Add(1 * time.Second)
	assert.Equal(t, 2, tickedCount)

	// Skip to Work
	p.Next()

	// Start LongBreak
	p.Next()
	assert.Equal(t, LongBreak, p.Kind)
	tickedCount = 0 // Reset count

	p.Start()

	clockMock.Add(1 * time.Second)
	assert.Equal(t, 1, tickedCount)
	clockMock.Add(1 * time.Second)
	assert.Equal(t, 2, tickedCount)
}

func TestPomodoro_OnEnd(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		WorkDuration:       25 * time.Second,
		ShortBreakDuration: 5 * time.Second,
		LongBreakDuration:  10 * time.Second,
		WorkRounds:         2,
		Clock:              clockMock,
	})

	// Capture end event
	endCalled := false
	var endKind Kind
	p.OnEnd = func(kind Kind) {
		endCalled = true
		endKind = kind
	}

	// Start and finish Work
	p.Start()
	clockMock.Add(25 * time.Second)

	assert.True(t, endCalled)
	assert.Equal(t, Work, endKind)
	assert.False(t, p.Running)

	// Start and finish ShortBreak
	assert.Equal(t, 5*time.Second, p.RemainingTime)
	p.Start()
	clockMock.Add(5 * time.Second)

	assert.True(t, endCalled)
	assert.Equal(t, ShortBreak, endKind)
	assert.False(t, p.Running)

	// Skip Work
	p.Start()
	p.Next()

	// Start and finish LongBreak
	assert.Equal(t, 10*time.Second, p.RemainingTime)
	p.Start()
	clockMock.Add(10 * time.Second)

	assert.True(t, endCalled)
	assert.Equal(t, LongBreak, endKind)
	assert.False(t, p.Running)
}

func TestPomodoro_Stop(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		WorkDuration:       25 * time.Second,
		ShortBreakDuration: 5 * time.Second,
		LongBreakDuration:  10 * time.Second,
		WorkRounds:         2,
		Clock:              clockMock,
	})

	// Start Work
	p.Start()

	// Stop Work
	p.Stop()
	assert.False(t, p.Running)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.RemainingTime)

	// Start ShortBreak
	p.Next()
	p.Start()

	// Stop ShortBreak
	p.Stop()
	assert.False(t, p.Running)
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 5*time.Second, p.RemainingTime)

	// Skip Work
	p.Next()

	// Start LongBreak
	p.Next()
	p.Start()

	// Stop LongBreak
	p.Stop()
	assert.False(t, p.Running)
	assert.Equal(t, LongBreak, p.Kind)
	assert.Equal(t, 10*time.Second, p.RemainingTime)
}

func TestPomodoro_Pause(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		WorkDuration:       25 * time.Second,
		ShortBreakDuration: 5 * time.Second,
		Clock:              clockMock,
	})

	// Start Work
	p.Start()
	clockMock.Add(1 * time.Second)

	// Pause
	p.Pause()
	assert.False(t, p.Running)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 24*time.Second, p.RemainingTime)

	// Finish Work
	p.Start()
	clockMock.Add(25 * time.Second)
	// Start ShortBreak
	p.Start()
	clockMock.Add(1 * time.Second)

	// Pause
	p.Pause()
	assert.False(t, p.Running)
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 4*time.Second, p.RemainingTime)
}

func TestPomodoro_Next(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		WorkDuration:       25 * time.Second,
		ShortBreakDuration: 5 * time.Second,
		LongBreakDuration:  10 * time.Second,
		WorkRounds:         3,
		Clock:              clockMock,
	})

	// Work
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.RemainingTime)
	assert.False(t, p.Running)

	// ShortBreak
	p.Next()
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 5*time.Second, p.RemainingTime)
	assert.False(t, p.Running)

	// Work
	p.Next()
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.RemainingTime)
	assert.False(t, p.Running)

	// ShortBreak
	p.Next()
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 5*time.Second, p.RemainingTime)
	assert.False(t, p.Running)

	// Work
	p.Next()
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.RemainingTime)
	assert.False(t, p.Running)

	// LongBreak
	p.Next()
	assert.Equal(t, LongBreak, p.Kind)
	assert.Equal(t, 10*time.Second, p.RemainingTime)
	assert.False(t, p.Running)

	// Work
	p.Next()
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.RemainingTime)
	assert.False(t, p.Running)

	// While running

	// ShortBreak
	p.Next()
	p.Start()
	clockMock.Add(1 * time.Second)
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 4*time.Second, p.RemainingTime)
	assert.True(t, p.Running)

	// Work
	p.Next()
	p.Start()
	clockMock.Add(1 * time.Second)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 24*time.Second, p.RemainingTime)
	assert.True(t, p.Running)

	// ShortBreak
	p.Next()
	p.Start()
	clockMock.Add(1 * time.Second)
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 4*time.Second, p.RemainingTime)
	assert.True(t, p.Running)

	// Work
	p.Next()
	p.Start()
	clockMock.Add(1 * time.Second)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 24*time.Second, p.RemainingTime)
	assert.True(t, p.Running)

	// LongBreak
	p.Next()
	p.Start()
	clockMock.Add(1 * time.Second)
	assert.Equal(t, LongBreak, p.Kind)
	assert.Equal(t, 9*time.Second, p.RemainingTime)
	assert.True(t, p.Running)

	// Work
	p.Next()
	p.Start()
	clockMock.Add(1 * time.Second)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 24*time.Second, p.RemainingTime)
	assert.True(t, p.Running)
}

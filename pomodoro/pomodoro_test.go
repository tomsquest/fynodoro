package pomodoro

import (
	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPomodoro_OnTick(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		WorkDuration:       3 * time.Second,
		ShortBreakDuration: 3 * time.Second,
		Clock:              clockMock,
	})

	// Capture tick events
	tickedCount := 0
	p.OnTick = func() {
		tickedCount++
	}

	p.Start()

	clockMock.Add(1 * time.Second)
	assert.Equal(t, 1, tickedCount)
	clockMock.Add(1 * time.Second)
	assert.Equal(t, 2, tickedCount)

	// Finish Work, start ShortBreak
	clockMock.Add(3 * time.Second)
	assert.Equal(t, ShortBreak, p.Kind)
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
	assert.Equal(t, 5*time.Second, p.Remaining)

	// Start and finish ShortBreak
	p.Start()
	clockMock.Add(5 * time.Second)

	assert.True(t, endCalled)
	assert.Equal(t, ShortBreak, endKind)
	assert.False(t, p.Running)
	assert.Equal(t, 25*time.Second, p.Remaining)
}

func TestPomodoro_Stop(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		WorkDuration:       25 * time.Second,
		ShortBreakDuration: 5 * time.Second,
		Clock:              clockMock,
	})

	// Start Work
	p.Start()
	clockMock.Add(1 * time.Second)

	// Stop Work
	p.Stop()
	assert.False(t, p.Running)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.Remaining)

	// Finish Work, start ShortBreak
	p.Start()
	clockMock.Add(25 * time.Second)

	// Stop ShortBreak
	p.Stop()
	assert.False(t, p.Running)
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 5*time.Second, p.Remaining)

	// Finish ShortBreak, start Work
	p.Start()
	clockMock.Add(5 * time.Second)

	// Stop Work
	p.Stop()
	assert.False(t, p.Running)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.Remaining)

	// 2nd Pomodoro

	// Finish Work, start ShortBreak
	p.Start()
	clockMock.Add(25 * time.Second)

	// Stop ShortBreak
	p.Stop()
	assert.False(t, p.Running)
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 5*time.Second, p.Remaining)

	// Finish ShortBreak, start Work
	p.Start()
	clockMock.Add(5 * time.Second)

	// Stop Work
	p.Stop()
	assert.False(t, p.Running)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.Remaining)
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
	assert.Equal(t, 24*time.Second, p.Remaining)

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
	assert.Equal(t, 4*time.Second, p.Remaining)
}

func TestPomodoro_Next(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&Params{
		WorkDuration:       25 * time.Second,
		ShortBreakDuration: 5 * time.Second,
		Clock:              clockMock,
	})

	// Initial state
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.Remaining)

	// Skip work
	p.Next()
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 5*time.Second, p.Remaining)

	// Skip shortBreak
	p.Next()
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.Remaining)

	// Skip work again
	p.Next()
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 5*time.Second, p.Remaining)

	// Skip shortBreak again
	p.Next()
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.Remaining)

	// While running

	// Start Work
	p.Start()
	clockMock.Add(1 * time.Second)

	// Skip work
	p.Next()
	assert.False(t, p.Running)
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 5*time.Second, p.Remaining)

	// Start Pause
	p.Start()
	clockMock.Add(1 * time.Second)

	// Skip pause
	p.Next()
	assert.False(t, p.Running)
	assert.Equal(t, Work, p.Kind)
	assert.Equal(t, 25*time.Second, p.Remaining)
}

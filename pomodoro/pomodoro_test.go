package pomodoro

import (
	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOk(t *testing.T) {
	clockMock := clock.NewMock()
	p := NewPomodoro(&PomodoroParams{
		WorkDuration:       4 * time.Second,
		ShortBreakDuration: 2 * time.Second,
		Clock:              clockMock,
	})

	p.Start()

	// Capture tick events
	tickedCount := 0
	p.OnTick = func() {
		tickedCount++
	}

	// Capture end event
	endCalled := false
	var endKind Kind
	p.OnEnd = func(kind Kind) {
		endCalled = true
		endKind = kind
	}

	// Ticks
	clockMock.Add(1 * time.Second)
	assert.Equal(t, 1, tickedCount)
	clockMock.Add(1 * time.Second)
	assert.Equal(t, 2, tickedCount)

	// Finish Work
	clockMock.Add(99 * time.Second)

	assert.False(t, p.Running)
	assert.Equal(t, ShortBreak, p.Kind)
	assert.Equal(t, 2*time.Second, p.Remaining)

	// Check End
	assert.True(t, endCalled)
	assert.Equal(t, Work, endKind)
}

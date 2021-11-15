package main

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	tables := []struct {
		d        time.Duration
		expected string
	}{
		{0, "00:00"},
		{1 * time.Second, "00:01"},
		{60 * time.Second, "01:00"},
		{61 * time.Second, "01:01"},
		{121 * time.Second, "02:01"},
		{25 * time.Minute, "25:00"},
		{61 * time.Minute, "61:00"},
	}

	for _, table := range tables {
		s := formatDuration(table.d)
		if s != table.expected {
			t.Errorf("Format of %v was incorrect, got: %v, want: %v.", table.d, s, table.expected)
		}
	}
}

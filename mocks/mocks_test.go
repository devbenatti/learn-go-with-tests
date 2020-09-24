package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"
)

const pause = "pause"
const writing = "writing"

type sleeperSpy struct {
	calls int
}

func (s *sleeperSpy) Sleep() {
	s.calls++
}

var _ sleeper = (*SpyCountOperations)(nil)
var _ sleeper = (*sleeperSpy)(nil)
var _ io.Writer = (*SpyCountOperations)(nil)

type SpyCountOperations struct {
	calls []string
}

func (s *SpyCountOperations) Sleep() {
	s.calls = append(s.calls, pause)
}

func (s *SpyCountOperations) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, writing)
	return
}

type TimeSpy struct {
	durationPause time.Duration
}

func (t *TimeSpy) pause(duration time.Duration) {
	t.durationPause = duration
}

func TestScore(t *testing.T) {
	t.Run("pause before home impression", func(t *testing.T) {
		spyPrinterSleep := &SpyCountOperations{}
		score(spyPrinterSleep, spyPrinterSleep)
		expect := []string{
			pause,
			writing,
			pause,
			writing,
			pause,
			writing,
			pause,
			writing,
		}

		if !reflect.DeepEqual(expect, spyPrinterSleep.calls) {
			t.Errorf("expect %v calls - result %v", expect, spyPrinterSleep.calls)
		}
	})

	t.Run("prints 3 until it Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &sleeperSpy{}

		score(buffer, sleeperSpy)

		result := buffer.String()
		expect := `3
2
1
Go!
`
		if result != expect {
			t.Errorf("result '%s' - expect '%s'", result, expect)
		}

		if sleeperSpy.calls != 4 {
			t.Errorf("there were not enough sleeper calls, expected 4, result %d", sleeperSpy.calls)
		}
	})
}

func TestSleeperConfig(t *testing.T) {
	timePause := 2 * time.Second

	timeSpy := &TimeSpy{}
	sleeper := sleeperConfig{timePause, timeSpy.pause}
	sleeper.Sleep()

	if timeSpy.durationPause != timePause {
		t.Errorf("should have paused for %v but paused for %v", timePause, timeSpy.durationPause)
	}
}

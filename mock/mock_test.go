package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// SpySleeper for count calls
type SpySleeper struct {
	calls int
}

// Call the sleeper
func (s *SpySleeper) Sleep(d time.Duration) {
	s.calls++
}

// OperationSpy impl
type OperationSpy struct {
	Calls []string
}

// DurationSpy duration
type DurationSpy struct {
	SleepTimes []time.Duration
}

//Sleep operation
func (s *DurationSpy) Sleep(d time.Duration) {
	s.SleepTimes = append(s.SleepTimes, d)
}

//Sleep operation
func (s *OperationSpy) Sleep(d time.Duration) {
	s.Calls = append(s.Calls, sleep)
}

//Write operation
func (s *OperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}
	sleepTime := 0 * time.Second

	Countdown(buffer, sleeper, sleepTime)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if sleeper.calls != 4 {
		t.Errorf("wrong amount of calls %d", sleeper.calls)
	}

	t.Run("check operation ordering", func(t *testing.T) {
		sleeper := &OperationSpy{}
		sleepTime := 0 * time.Second

		Countdown(sleeper, sleeper, sleepTime)

		got := sleeper.Calls
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("check sleep time", func(t *testing.T) {
		sleeper := &OperationSpy{}
		sleepTime := 1 * time.Second
		spy := &DurationSpy{}

		Countdown(sleeper, spy, sleepTime)

		got := spy.SleepTimes
		want := []time.Duration{
			sleepTime,
			sleepTime,
			sleepTime,
			sleepTime,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

const write = "write"
const sleep = "sleep"

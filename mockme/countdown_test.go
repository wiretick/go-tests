package mockme

import (
	"bytes"
	"slices"
	"testing"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

func TestCountdown(t *testing.T) {
	t.Run("prints the countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		Countdown(buffer, spySleeper)

		want := `3 2 1 GO!`
		got := buffer.String()

		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("sleep after print", func(t *testing.T) {
		spySleep := &SpyCountdownOperations{}
		Countdown(spySleep, spySleep)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, spySleep.Calls) {
			t.Errorf("want %v, got %v", want, spySleep.Calls)
		}
	})
}

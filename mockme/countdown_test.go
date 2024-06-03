package mockme

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	want := `3 2 1 GO!`
	got := buffer.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

	if spySleeper.Calls < 3 {
		t.Errorf("want %d, got %d", 3, spySleeper.Calls)
	}
}

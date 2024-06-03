package mockme

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	want := `3 2 1 GO!`
	got := buffer.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

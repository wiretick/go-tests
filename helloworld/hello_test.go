package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("say hello to person", func(t *testing.T) {
		got := Hello("Jon", "")
		want := "Hello Jon"

		assertTwoStrings(t, got, want)
	})

	t.Run("generic greeting", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertTwoStrings(t, got, want)
	})

	t.Run("greeting in norwegian", func(t *testing.T) {
		got := Hello("Per", "Norwegian")
		want := "Hei Per"

		assertTwoStrings(t, got, want)
	})

	t.Run("greeting in french", func(t *testing.T) {
		got := Hello("Per", "French")
		want := "Bonjour Per"

		assertTwoStrings(t, got, want)
	})
}

func TestGreet(t *testing.T) {
	// Testing dependency injection
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tom")

	got := buffer.String()
	want := "Hello, Tom"

	assertTwoStrings(t, got, want)
}

func assertTwoStrings(t testing.TB, got, want string) {
	t.Helper() // this gives the errorf context to say the actual line nr where the test fails

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

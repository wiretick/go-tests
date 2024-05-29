package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to person", func(t *testing.T) {
		got := hello("Jon", "")
		want := "Hello Jon"

		assertTwoStrings(t, got, want)
	})

	t.Run("generic greeting", func(t *testing.T) {
		got := hello("", "")
		want := "Hello World"

		assertTwoStrings(t, got, want)
	})

	t.Run("greeting in norwegian", func(t *testing.T) {
		got := hello("Per", "Norwegian")
		want := "Hei Per"

		assertTwoStrings(t, got, want)
	})

	t.Run("greeting in french", func(t *testing.T) {
		got := hello("Per", "French")
		want := "Bonjour Per"

		assertTwoStrings(t, got, want)
	})
}

func assertTwoStrings(t testing.TB, got, want string) {
	t.Helper() // this gives the errorf context to say the actual line nr where the test fails

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

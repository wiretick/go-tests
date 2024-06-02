package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("define term", func(t *testing.T) {
		dict := Dictionary{"hello": "Hello is a casual greeting"}
		got, _ := dict.Search("hello")
		want := "Hello is a casual greeting"

		assertString(t, got, want)
	})

	t.Run("search for unknown term", func(t *testing.T) {
		dict := Dictionary{"hello": "Hello is a casual greeting"}
		_, got := dict.Search("unknown")
		want := ErrUnknownTerm

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	dict.Add("new term", "new definition")

	want := "new definition"
	got, err := dict.Search("new term")

	if err != nil {
		t.Fatalf("Should not error: %q", err)
	}

	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}

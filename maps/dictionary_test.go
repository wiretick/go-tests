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
	t.Run("add new term", func(t *testing.T) {
		term := "new term"
		def := "new definition"
		dict := Dictionary{}

		what := dict.Add(term, def)
		if what != nil {
			t.Fatal("what", what)
		}
		got, err := dict.Search(term)

		if err != nil {
			t.Fatalf("Should not error: %q", err)
		}

		assertString(t, got, def)
	})

	t.Run("add existing term", func(t *testing.T) {
		term := "new term"
		def := "new definition"
		dict := Dictionary{term: def}

		got := dict.Add(term, def)
		want := ErrAlreadyExists

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}

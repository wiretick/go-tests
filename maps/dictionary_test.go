package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"hello": "Hello is a casual greeting"}
	got := dict.Search("hello")
	want := "Hello is a casual greeting"

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}

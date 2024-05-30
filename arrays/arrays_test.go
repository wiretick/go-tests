package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [...]int{4, 4, 4, 4}
	got := Sum(numbers)
	want := 16

	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	}
}

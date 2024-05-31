package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{8, 10}
	got := Sum(numbers)
	want := 18

	assert(t, want, got, numbers)
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{8, 10}
	numbers2 := []int{5, 10}

	got := SumAll(numbers1, numbers2)
	want := []int{18, 15}

	if !slices.Equal(got, want) {
		t.Errorf("want %d, got %d, given %v and %v", want, got, numbers1, numbers2)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum of slices", func(t *testing.T) {
		want := []int{8, 10}
		got := SumAllTails([]int{4, 8}, []int{0, 10})

		if !slices.Equal(want, got) {
			t.Errorf("want %d, got %d", want, got)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		want := []int{0, 3}
		got := SumAllTails([]int{}, []int{0, 1, 2})

		if !slices.Equal(want, got) {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}

func assert(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d, got %d, given %v", want, got, numbers)
	}
}

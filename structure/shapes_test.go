package structure

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 5., Height: 5.}
	want := 20.
	got := rect.Perimeter()

	if want != got {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if want != got {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	}

	t.Run("for rectangle", func(t *testing.T) {
		rect := Rectangle{Width: 5., Height: 5.}
		checkArea(t, rect, 25.)
	})

	t.Run("for rectangle", func(t *testing.T) {
		circle := Circle{Radius: 10.}
		checkArea(t, circle, 314.1592653589793)
	})
}

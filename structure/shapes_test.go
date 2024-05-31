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
	t.Run("for rectangle", func(t *testing.T) {
		rect := Rectangle{Width: 5., Height: 5.}
		want := 25.
		got := rect.Area()

		if want != got {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	})

	t.Run("for rectangle", func(t *testing.T) {
		circle := Circle{Radius: 10.}
		want := 314.1592653589793
		got := circle.Area()

		if want != got {
			t.Errorf("want %g, got %g", want, got)
		}
	})
}

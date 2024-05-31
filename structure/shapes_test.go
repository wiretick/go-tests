package structure

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 5., Height: 5.}
	want := 20.
	got := Perimeter(rect)

	if want != got {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	rect := Rectangle{Width: 5., Height: 5.}
	want := 25.
	got := Area(rect)

	if want != got {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

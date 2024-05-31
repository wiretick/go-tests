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
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 5, Height: 5}, want: 25},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) { // gives better context for failed tests by using name
			got := tt.shape.Area()
			if tt.want != got {
				t.Errorf("%#v, want %.2f, got %.2f", tt.shape, tt.want, got)
			}
		})
	}
}

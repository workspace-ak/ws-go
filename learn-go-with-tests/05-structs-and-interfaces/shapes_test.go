package shapes

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := rectangle.Perimeter()
	want := 36.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

// table driven tests
func TestArea(t *testing.T) {
	areaTests := map[string]struct {
		shape Shape
		want  float64
	}{
		"rectangle": {
			shape: Rectangle{12, 6},
			want:  72.0,
		},
		"circle": {
			shape: Circle{10},
			want:  314.1592653589793,
		},
	}
	for name, test := range areaTests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, want := test.shape.Area(), test.want; got != want {
				t.Errorf("got %g, want %g", got, test.want)

			}
		})
	}
}

func ExampleCalculateArea() {
	circle := Circle{10}
	res := CalculateArea(circle)
	fmt.Printf("%g", res)
	// Output: 314.1592653589793
}

func ExampleCalculateArea_second() {
	rectangle := Rectangle{12.0, 6.0}
	res := CalculateArea(rectangle)
	fmt.Println(res)
	// Output: 72

}

func ExampleCircle_Area() {
	circle := Circle{10}
	res := circle.Area()
	fmt.Printf("%g", res)
	// Output: 314.1592653589793
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{12.0, 6.0}
	res := rectangle.Area()
	fmt.Println(res)
	// Output: 72
}

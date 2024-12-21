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

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, s Shape, want float64) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
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

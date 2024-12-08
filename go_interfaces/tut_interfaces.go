package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measurable
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func describeGeometry(g Geometry) {
	fmt.Printf("[g] Area: %.2f; Perimeter: %.2f\n", g.Area(), g.Perimeter())
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// type Circle struct {
// 	radius float64
// }

//	func (c Circle) Area() float64 {
//		return math.Pow(c.radius, 2) * math.Pi
//	}
type CalulationError struct {
	msg string
}

func (ce CalulationError) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalulationError{msg: "Invalid input"}
	}
	return math.Sqrt(val), nil
}

type Tiles struct {
	genre       string
	length      float64
	width       float64
	thickness   float64
	tilesPerBox int
}

func (t Tiles) Area() float64 {
	return t.length * t.width * float64(t.tilesPerBox) * 10.764 / 10000
}

func main() {
	rect := Rectangle{width: 5, height: 4}
	// cir := Circle{radius: 3.5}

	fmt.Println("[r] Rectangle Area: ", calculateArea(rect))
	// fmt.Println("[c] Circle Area: ", calculateArea(cir))
	describeGeometry(rect)

	// v, err := performCalculation(-5)
	v, err := performCalculation(10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	ceramicTile := Tiles{
		genre:       "ceramic",
		length:      60,
		width:       120,
		thickness:   9,
		tilesPerBox: 2,
	}

	fmt.Printf("Tiles Coverage (sqft.): %.2f", ceramicTile.Area())

}

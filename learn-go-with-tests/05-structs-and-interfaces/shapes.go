// package shapes provides primitives for calculating areas and perimeter for different shapes
package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
}

// Calculate Area of a given shape
func CalculateArea(s Shape) float64 {
	return s.Area()
}

type Rectangle struct {
	width  float64
	length float64
}

// Calculate Area of a rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.length
}

// Calculate Perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.length)
}

type Circle struct {
	radius float64
}

// Calculate Area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

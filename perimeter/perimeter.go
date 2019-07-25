package perimeter

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle type
type Rectangle struct {
	Width  float64
	Height float64
}

// Triangle type
type Triangle struct {
	Height float64
	Width  float64
}

// Area for triangle
func (t Triangle) Area() float64 {
	return (t.Height * t.Width / 2)
}

// Area for rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle type
type Circle struct {
	radius float64
}

// Area for circle
func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

// Perimeter calculates the perimeter of a rectangle
func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = (rectangle.Width + rectangle.Height) * 2
	return
}

// Area calculates area of rectangle
func Area(rectangle Rectangle) (area float64) {
	area = rectangle.Width * rectangle.Height
	return
}

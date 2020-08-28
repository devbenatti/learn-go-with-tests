package structs

import "math"

// Form interface for rectangle and circle
type Form interface {
	Area() float64
}

// Rectangle represents a rectangle xD
type Rectangle struct {
	Width  float64
	Height float64
}

// Area receives width and height and return area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle xD
type Circle struct {
	Radius float64
}

// Area receives width and height and return area
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents a triangle xD
type Triangle struct {
	Base   float64
	Height float64
}

// Area receives width and height and return area
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

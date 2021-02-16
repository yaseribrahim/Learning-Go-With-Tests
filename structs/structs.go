package structs

import "math"

type Shape interface {
	Area() float64
}

//Rectangle holds width and height of rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Perimeter takes width and height of rectangle and returns perimeter length
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area takes width and height of rectangle and returns size of area
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

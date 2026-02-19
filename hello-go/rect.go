package main

type Rectangle struct {
	height float64
	width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) Scale(f float64) {
	r.height = r.height * f
	r.width = r.width * f
}

type Color struct {
	color string
}

type ColoredRectangle struct {
	Rectangle
	Color
}
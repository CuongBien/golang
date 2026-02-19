package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %0.2f\n", s.Area())
}
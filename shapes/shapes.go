package main

import (
	"math"
)

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (t Triangle) Perimeter() float64 {
	return 3 * t.base // Assuming an equilateral triangle for simplicity
}

func (r *Rectangle) Scale(f float64) {
	r.width *= f
	r.height *= f
}

func (c *Circle) Scale(f float64) {
	c.radius *= f
}

func (t *Triangle) Scale(f float64) {
	t.base *= f
	t.height *= f
}

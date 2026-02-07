package main

import "testing"

/* Rectangle Tests */

func TestRectangleArea(t *testing.T) {
	r := Rectangle{4, 5}

	if r.Area() != 20 {
		t.Errorf("Expected 20, got %v", r.Area())
	}
}

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{4, 5}

	if r.Perimeter() != 18 {
		t.Errorf("Expected 18, got %v", r.Perimeter())
	}
}

func TestRectangleScale(t *testing.T) {
	r := Rectangle{2, 3}

	r.Scale(2)

	if r.width != 4 || r.height != 6 {
		t.Errorf("Scale failed")
	}
}

/* Circle Tests */

func TestCircleArea(t *testing.T) {
	c := Circle{1}

	if c.Area() <= 3 {
		t.Errorf("Invalid circle area")
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{1}

	if c.Perimeter() <= 6 {
		t.Errorf("Invalid circle perimeter")
	}
}

func TestCircleScale(t *testing.T) {
	c := Circle{2}

	c.Scale(2)

	if c.radius != 4 {
		t.Errorf("Scale failed")
	}
}

/* Triangle Tests */

func TestTriangleArea(t *testing.T) {
	tr := Triangle{4, 6}

	if tr.Area() != 12 {
		t.Errorf("Expected 12, got %v", tr.Area())
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tr := Triangle{4, 6}

	if tr.Perimeter() != 12 {
		t.Errorf("Expected 12, got %v", tr.Perimeter())
	}
}

func TestTriangleScale(t *testing.T) {
	tr := Triangle{2, 3}

	tr.Scale(2)

	if tr.base != 4 || tr.height != 6 {
		t.Errorf("Scale failed")
	}
}

package main

import (
	"math"
	"testing"
)

const tolerance = 0.0001

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < tolerance
}

// Rectangle Tests
func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	if !almostEqual(r.Area(), 20) {
		t.Errorf("Expected 20, got %v", r.Area())
	}
}

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	if !almostEqual(r.Perimeter(), 18) {
		t.Errorf("Expected 18, got %v", r.Perimeter())
	}
}

func TestRectangleScale(t *testing.T) {
	r := Rectangle{Width: 2, Height: 3}
	r.Scale(2)
	if r.Width != 4 || r.Height != 6 {
		t.Errorf("Scale failed")
	}
}

// Circle Tests
func TestCircleArea(t *testing.T) {
	c := Circle{Radius: 3}
	expected := math.Pi * 9
	if !almostEqual(c.Area(), expected) {
		t.Errorf("Expected %v, got %v", expected, c.Area())
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{Radius: 3}
	expected := 2 * math.Pi * 3
	if !almostEqual(c.Perimeter(), expected) {
		t.Errorf("Expected %v, got %v", expected, c.Perimeter())
	}
}

func TestCircleScale(t *testing.T) {
	c := Circle{Radius: 3}
	c.Scale(3)
	if c.Radius != 9 {
		t.Errorf("Scale failed")
	}
}

// Triangle Tests
func TestTriangleArea(t *testing.T) {
	tg := Triangle{Base: 6, Height: 4}
	if !almostEqual(tg.Area(), 12) {
		t.Errorf("Expected 12, got %v", tg.Area())
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tg := Triangle{Base: 6, Height: 4}
	if !almostEqual(tg.Perimeter(), 18) {
		t.Errorf("Expected 18, got %v", tg.Perimeter())
	}
}

func TestTriangleScale(t *testing.T) {
	tg := Triangle{Base: 3, Height: 4}
	tg.Scale(2)
	if tg.Base != 6 || tg.Height != 8 {
		t.Errorf("Scale failed")
	}
}

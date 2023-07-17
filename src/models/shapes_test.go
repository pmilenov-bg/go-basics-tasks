package models

import (
	"math"
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	expectedArea := 15.0

	area := rect.Area()

	if area != expectedArea {
		t.Errorf("Rectangle.Area() returned %f, expected %f", area, expectedArea)
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	expectedPerimeter := 16.0

	perimeter := rect.Perimeter()

	if perimeter != expectedPerimeter {
		t.Errorf("Rectangle.Perimeter() returned %f, expected %f", perimeter, expectedPerimeter)
	}
}

func TestCircle_Area(t *testing.T) {
	circle := Circle{Radius: 2.5}
	expectedArea := math.Pi * 2.5 * 2.5

	area := circle.Area()

	if area != expectedArea {
		t.Errorf("Circle.Area() returned %f, expected %f", area, expectedArea)
	}
}

func TestCircle_Perimeter(t *testing.T) {
	circle := Circle{Radius: 2.5}
	expectedPerimeter := 2 * math.Pi * 2.5

	perimeter := circle.Perimeter()

	if perimeter != expectedPerimeter {
		t.Errorf("Circle.Perimeter() returned %f, expected %f", perimeter, expectedPerimeter)
	}
}

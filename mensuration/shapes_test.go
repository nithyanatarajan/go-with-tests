package mensuration_test

import (
	"testing"

	"github.com/nithyanatarajan/go-with-tests/mensuration"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   mensuration.Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: mensuration.Rectangle{Height: 12, Width: 10}, hasArea: 120.0},
		{name: "Circle", shape: mensuration.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: mensuration.Triangle{Side1: 10, Side2: 11, Side3: 3}, hasArea: 14.696938456699069},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.hasArea
			AssertEqual(t, got, want)
		})

	}
}

func TestPerimeter(t *testing.T) {
	areaTests := []struct {
		name         string
		shape        mensuration.Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: mensuration.Rectangle{Height: 12, Width: 10}, hasPerimeter: 44.0},
		{name: "Circle", shape: mensuration.Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: mensuration.Triangle{Side1: 10, Side2: 11, Side3: 3}, hasPerimeter: 24},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			want := tt.hasPerimeter
			AssertEqual(t, got, want)
		})

	}
}

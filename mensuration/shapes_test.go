package mensuration_test

import (
	"testing"

	"github.com/nithyanatarajan/go-with-tests/mensuration"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

func TestRectangle_Perimeter(t *testing.T) {
	rectangle := mensuration.Rectangle{
		Height: 10.0,
		Width:  10.0,
	}
	want := 40.0

	got := rectangle.Perimeter()

	AssertEqual(t, got, want)
}

func TestRectangle_Area(t *testing.T) {
	rectangle := mensuration.Rectangle{
		Height: 12.0,
		Width:  10.0,
	}
	want := 120.0

	got := rectangle.Area()

	AssertEqual(t, got, want)
}

func TestCircle_Perimeter(t *testing.T) {
	rectangle := mensuration.Circle{
		Radius: 10.0,
	}
	want := 62.83185307179586

	got := rectangle.Perimeter()

	AssertEqual(t, got, want)
}

func TestCircle_Area(t *testing.T) {
	rectangle := mensuration.Circle{
		Radius: 10.0,
	}
	want := 314.1592653589793

	got := rectangle.Area()

	AssertEqual(t, got, want)
}

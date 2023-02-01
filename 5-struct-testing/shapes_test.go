package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		got := rectangle.Perimeter()
		want := 60.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := math.Pi * 2 * 10.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	got := rectangle.Area()
	want := 200.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

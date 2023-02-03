package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerim := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		want := 60.0
		checkPerim(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := math.Pi * 2 * 10.0
		checkPerim(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangles Area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		got := rectangle.Area()
		want := 200.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles Area", func(t *testing.T) {

	})
}

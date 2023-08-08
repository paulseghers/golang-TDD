package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chungus")

	got := buffer.String()
	want := "Hello, Chungus"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{} //read incoming stuff

	Countdown(buffer)

	got := buffer.String() //returns the whole chungus from the byte slice that this buffer points to/"implements"(?)
	want := `3
	2
	1
	Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := bytes.Buffer{}

	CountDown(&buffer)

	got := buffer.String()
	want := `3
2
1
0
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

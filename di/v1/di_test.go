package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gogosha")

	got := buffer.String()
	want := "Hello, Gogosha"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

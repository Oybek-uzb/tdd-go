package main

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := X{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
}

type X struct {
	Name string
}

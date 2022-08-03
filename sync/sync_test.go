package main

import "testing"

func TestCounter(t *testing.T) {
	counter := Counter{}

	counter.Inc()
	counter.Inc()
	counter.Inc()

	if counter.Value() != 4 {
		t.Errorf("got %d, want %d", counter.Value(), 3)
	}
}

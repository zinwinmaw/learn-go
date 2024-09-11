package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(5, 10)
	want := -5
	if total != want {
		t.Errorf("Substract was incorrect, go %v, want: %v", total, want)
	}
}

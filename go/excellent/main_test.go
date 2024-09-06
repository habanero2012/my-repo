package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}

func TestSum(t *testing.T) {
	result := sum(10, 20)
	if result != 30 {
		t.Errorf("expected: even, actual: %d", result)
	}
}

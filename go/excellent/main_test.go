package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	return := EvenOrOdd(10)
	if return != "even" {
		t.Errorf("expected: even, actual: %s", return)
	}
}
package main

import "testing"

func TestSubtractOneFromOne(t *testing.T) {
	result := subtractOneFromOne()
	if result != 0 {
		t.Error("Expected", 0, "Got", result)
	}
}
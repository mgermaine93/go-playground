// // First test writing attempt in Go...

package main

import "testing"

func TestSumUpTheseNumbers(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	// This is an example of a table test
	// Basically, it's a bunch of tests in one file
	tests := []test{
		// These are values of type test, detailed above
		// data is first, answer is second
		{[]int{1, 6}, 7},
		{[]int{8, 25}, 33},
		{[]int{-9, 10}, 1},
		{[]int{435, 2}, 437},
	}

	// Range over the tests
	for _, value := range tests {
		result := sumUpTheseNumbers(value.data...) // unfurl the data
		if result != value.answer {
			t.Error("Expected", value.answer, "Got", result)
		}
	}

}

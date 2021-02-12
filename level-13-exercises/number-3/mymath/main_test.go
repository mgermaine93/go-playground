package mymath

import (
	"fmt"
	"testing"
)

// This is a Table Test
func TestCenteredAvg(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	// This is an example of a table test
	// Basically, it's a bunch of tests in one file
	tests := []test{
		// These are values of type test, detailed above
		// data is first, answer is second
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{2, 4, 6, 8, 10}, 6},
		{[]int{-10, 0, 10}, 0},
		{[]int{-100003, 0, 100003}, 0},
	}

	// Range over the tests
	for _, value := range tests {
		result := CenteredAvg(value.data) // unfurl the data
		if result != value.answer {
			t.Error("Expected", value.answer, "Got", result)
		}
	}

}

// This is an Example
func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{6, 7, 8, 9, 10}))
	// Output:
	// 8
}

// This is benchmarking
func BenchmarkCenteredAvg(benchmark *testing.B) {
	for i := 0; i < benchmark.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 200})
	}
}

// go test // tests
// go test -bench . // benchmarks
// go test -cover // shows coverage
// go test -coverprofile c.out // shows coverage in web browser
// go tool cover -html=c.out // shows documentation examples in web browser

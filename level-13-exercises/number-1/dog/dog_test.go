package dog

import (
	"fmt"
	"testing"
)

// This is a Test
func TestHumanYearsToDogYears(t *testing.T) {
	result := HumanYearsToDogYears(6)
	if result != 42 {
		t.Error("Expected", 42, "Got", result)
	}
}

// This is an Example
func ExampleHumanYearsToDogYears() {
	fmt.Println(HumanYearsToDogYears(7))
	// Output:
	// 49
}

// This is benchmarking
func BenchmarkHumanYearsToDogYears(benchmark *testing.B) {
	for i := 0; i < benchmark.N; i++ {
		HumanYearsToDogYears(3)
	}
}

// go test // tests
// go test -bench . // benchmarks
// go test -cover // shows coverage
// go test -coverprofile c.out // shows coverage in web browser
// go tool cover -html=c.out // shows documentation examples in web browser

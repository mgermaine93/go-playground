package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	result := Greet("Altoona")
	if result != "Hello there, Altoona" {
		t.Error("Expected", "Hello there, Altoona", "Got", result)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Leah"))
	// Output:
	// Hello there, Leah
}

// This is benchmarking
func BenchmarkGreet(benchmark *testing.B) {
	for i := 0; i < benchmark.N; i++ {
		Greet("Leah")
	}
}

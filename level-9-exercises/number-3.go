// Using GoRoutines, create an incrementer program.  Follow the comments below...

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("Go Routines", runtime.NumGoroutine())

	// Have a variable to hold the incrementer value
	count := 0

	// These are go routines
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		// Creates the Go Routine using an anonymous function
		go func() {
			// Read the incrementer value and store it in a new value
			value := count
			// Yield the processor
			runtime.Gosched()
			// Increment the new value
			value++
			// Write the value in the new variable back to the original incrementer
			count = value
			wg.Done()
		}()
		fmt.Println("Go Routines", runtime.NumGoroutine())
	}

	// "Don't exit my program until everything is done"
	wg.Wait()
	fmt.Println("Go Routines", runtime.NumGoroutine())
	fmt.Println("count: ", count)

}

// To prove that it is a race condition, run "go run -race number-3.go"

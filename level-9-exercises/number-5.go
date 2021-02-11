// Fix the race condition created in #3 by using package atomic.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("Go Routines", runtime.NumGoroutine())

	/// int64 is usually a giveaway for "package atomic"
	var count int64
	var wg sync.WaitGroup

	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		// Creates the Go Routine using an anonymous function
		go func() {

			atomic.AddInt64(&count, 1) // Change by 1
			fmt.Println("Counter\t", atomic.LoadInt64(&count))
			wg.Done()

		}()
	}

	// "Don't exit my program until everything is done"
	wg.Wait()
	fmt.Println("End Value: ", count)

}

// To prove that there isn't a race condition, run "go run -race number-5.go"

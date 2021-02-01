// This is an example of a Race Condition.

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

	// int64 is usually a giveaway for "package atomic"
	var counter int64

	// These are go routines
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// This creates a mutex
	// var mutex sync.Mutex

	for i := 0; i < gs; i++ {
		// Creates the Go Routine
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			// mutex.Lock()
			// v := counter
			// time.sleep(time.Second)
			// runtime.Gosched()
			// v++
			// "Write the value back to the shared variable"
			// counter = v
			// mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routines", runtime.NumGoroutine())
	}

	// "Don't exit my program until everything is done"
	wg.Wait()
	fmt.Println("Go Routines", runtime.NumGoroutine())
	fmt.Println("count: ", counter)

}

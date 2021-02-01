package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Loads in the WaitGroup
var waitGroup sync.WaitGroup

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// "Adds one thing to wait for"
	waitGroup.Add(1)
	// Launches out into its separate own Go routine
	// This is a concurrent design pattern
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// This waits until WaitGroup.Done() is hit
	waitGroup.Wait()
}

func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	// When "foo()" is done running, "remove the one thing to wait for"
	waitGroup.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

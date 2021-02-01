// In addition to the main go routine, launch two additional go routines.
// Each additional goroutine should print something out.
// Use WaitGroups to make sure that each go routine finishes before your program exits.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Loads in the WaitGroup
var waitGroup sync.WaitGroup

func main() {

	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// "Add two things to wait for"
	waitGroup.Add(2)

	// Launches out into its separate own Go routine
	// This is a concurrent design pattern
	go firstRoutine()
	go secondRoutine()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// This waits until WaitGroup.Done() is hit
	waitGroup.Wait()

	fmt.Println("The program has exited!")
}

func firstRoutine() {

	fmt.Println("I'm the first GoRoutine.")
	// When "firstRoutine()" is done running, "remove the one thing to wait for"
	waitGroup.Done()
}

func secondRoutine() {
	fmt.Println("I'm the second GoRoutine.")
	// When "firstRoutine()" is done running, "remove the second thing to wait for"
	waitGroup.Done()
}

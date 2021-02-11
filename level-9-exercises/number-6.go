// Create a program that prints out your OS and ARCH.  Use the following commands to run it:  go run, go build, go install.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This is the Operating System at runtime: ", runtime.GOOS)
	fmt.Println("This is the architecture target of the running program: ", runtime.GOARCH)
}

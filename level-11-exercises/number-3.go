// Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”.

package main

import "fmt"

type customErr struct {
	information string
}

// This has the error signature
func (anError customErr) Error() string {
	return fmt.Sprintf("An error: %v", anError.information)
}

func main() {
	customError1 := customErr{
		information: "I'm an error",
	}
	foo(customError1)
}

func foo(err error) {
	fmt.Println("Inside of foo()", err)
}

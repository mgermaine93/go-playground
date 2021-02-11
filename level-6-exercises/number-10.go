// "Closure" is when we have "enclosed" the scope of a variable in some code block.  Create a func that "encloses" the scope of a variable.

package main

import "fmt"

func main() {

	returnedValue := sayHelloToName("Fred")
	fmt.Println(returnedValue)
	// This will not print because "greeting" is out-of-scope
	// fmt.Println(greeting)

}

func sayHelloToName(name string) string {
	greeting := "Hello, "
	a := greeting + name
	return a
}

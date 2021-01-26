// Build and use an anonymous func

package main

import "fmt"

func main() {

	func(x int) {
		fmt.Printf("Brendan Shanahan's Red Wings number was: %v\n", x)
	}(14) // Arguments are passed in here.

}

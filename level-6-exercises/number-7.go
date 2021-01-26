// Assign a func to variable, then call the func.

package main

import "fmt"

func main() {

	number := func() {
		fmt.Printf("Brendan Shanahan's Red Wings number was %v\n", 14)
	}
	number()

}

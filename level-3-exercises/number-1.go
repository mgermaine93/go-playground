// Loops from 33 to 122 (inclusive)  and print out each number in ASCII and code points.

package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}

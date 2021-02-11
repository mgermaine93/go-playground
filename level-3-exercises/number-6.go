// Use a for {} loop to print out all of the years you've been alive.

package main

import "fmt"

func main() {
	// Current as of 2021, haha
	birthYear := 1993
	for {
		if birthYear > 2021 {
			break
		}
		fmt.Println(birthYear)
		birthYear++
	}
}

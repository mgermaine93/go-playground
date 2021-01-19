// Using iota, create 4 constants for the last four years.  Print the constant values.

package main

import "fmt"

func main() {

	const (
		year1 = 2021 + iota
		year2 = 2021 + iota
		year3 = 2021 + iota
		year4 = 2021 + iota
	)

	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)

}

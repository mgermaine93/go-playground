// Create a SLICE of a SLICE of TYPE string.  Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {

	bond := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(bond)

	moneypenny := []string{"Miss", "Moneypenny", "Helloooooooooo, James"}
	fmt.Println(moneypenny)

	sliceOfSlice := [][]string{bond, moneypenny}
	fmt.Println(sliceOfSlice)

	for index, record := range sliceOfSlice {
		fmt.Println("Record: ", index)
		for index2, dataPoint := range record {
			fmt.Printf("\t Index position: %v \t Value: \t %v \n ", index2, dataPoint)
		}
	}

}

// Create a map with a key of TYPE string which is a person's "last_first" name, and a value of type []string which stores their favorite things.  Store three recrods in your map.  Print the values, along with their index position in the slice.

// Create a SLICE of a SLICE of TYPE string.  Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {

	// First value is the key type, second value is the value type
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	for index, record := range m {
		fmt.Println("This is the record for: ", index)
		for index2, dataPoint := range record {
			fmt.Printf("\t Index position: %v \t Value: \t %v \n ", index2, dataPoint)
		}
	}
}

// Using the previous code, delete a record from the map.  Then, print the map out using the "range" loop.

package main

import "fmt"

func main() {

	// First value is the key type, second value is the value type
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	// Add to the map
	m["Pittsburgh"] = []string{"Penguins", "Steelers", "Pirates", "Riverhounds"}

	// Delete from the map
	delete(m, "no_dr")

	for index, record := range m {
		fmt.Println("This is the record for: ", index)
		for index2, dataPoint := range record {
			fmt.Printf("\t Index position: %v \t Value: \t %v \n ", index2, dataPoint)
		}
	}
}

// Package dog take in a quantity of "human" years
// and converts them to "dog" years.  One "human"
// year is equal to seven "dog" years.
package dog

// HumanYearsToDogYears takes in a single int representing
// a quantity of human years and returns the equivalent
// quantity of dog years.
func HumanYearsToDogYears(humanYears int) int {

	// Each dog year is equivalent to seven human years.
	dogYears := humanYears * 7
	return dogYears
}

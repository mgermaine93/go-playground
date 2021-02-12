// Package dmb inclues the "Sum" function
package dmb

// Sum adds together an unlimited number of values of type int
func Sum(sliceOfInt ...int) int {
	sum := 0
	for _, value := range sliceOfInt {
		sum += value
	}
	return sum
}

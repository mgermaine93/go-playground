// Starting with the following code, use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your solution:
// lat "50.2289 N"
// long "99.4656 W"

package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

// Takes in a variable se of type sqrtError and outputs a string, implements Error()
func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	// Catch the error here
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

// Outputs float64 and an error
func sqrt(f float64) (float64, error) {
	if f < 0 {
		anError := fmt.Errorf("I'm an error, my value is %v", f)
		// Need to return a float64 and a value of type "Error"
		return 0, sqrtError{"50.2289 N", "99.4656 W", anError}
	}
	return 42, nil
}

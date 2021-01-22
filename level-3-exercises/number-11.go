// Create a program that uses a switch statement WITH a switch expression specified.

package main

import "fmt"

func main() {

	favSport := "Running"

	switch favSport {
	case "Rowing":
		fmt.Println("You must like rowing!")
	case "Running":
		fmt.Println("You must like running!")
	case "Surfing":
		fmt.Println("You must like surfing!")
	case "Wrestling":
		fmt.Println("You must like wrestling!")
	case "Cycling":
		fmt.Println("You must like cycling!")
	case "Swimming":
		fmt.Println("You must like swimming!")
	case "Hockey":
		fmt.Println("You must like hockey!")
	case "Frisbee Golf":
		fmt.Println("You must like Frisbee golf!")
	case "Boxing":
		fmt.Println("You must like boxing!")
	case "Fencing":
		fmt.Println("You must like fencing!")
	default:
		fmt.Println("You must like a sport that involves a ball, or a sport I didn't think of!  Well done!")
	}
}

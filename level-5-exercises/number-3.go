// Follow the comments below

package main

import "fmt"

func main() {

	// Create a new type "vehicle" with underlying type "struct"
	// Fields should be "doors" and "color"
	type vehicle struct {
		doors int
		color string
	}

	// Create two new types, "truck" and "sedan"
	// The "vehicle" type should be embedded in each of these new types

	// "truck" should also have a field "fourWheel" set to "bool"
	type truck struct {
		vehicle
		fourWheel bool
	}

	// "sedan" should also have a field "luxury" set to "bool"
	type sedan struct {
		vehicle
		luxury bool
	}

	// Using a composite literal, create a value of type "truck" and assign values to the fields
	toyotaTacoma := truck{
		vehicle: vehicle{
			doors: 4,
			color: "scarlet",
		},
		fourWheel: true,
	}

	// Using a composite literal, create a value of type "sedan" and assign values to the fields
	hondaCivic := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "gold",
		},
		luxury: false,
	}

	// Print out each of the above values
	fmt.Println(toyotaTacoma)
	fmt.Println(hondaCivic)

	// Print out a single field from each of the above values
	fmt.Println(toyotaTacoma.fourWheel)
	fmt.Println(hondaCivic.color)

}

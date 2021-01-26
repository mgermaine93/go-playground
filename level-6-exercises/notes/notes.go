package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

// This is a method
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// Using an interface -- "any TYPE that has this METHOD 'speak' is also of TYPE 'human'"
// Uses the same keyword-identifier-type pattern, as always
// A VALUE can be of more than one TYPE
type human interface {
	speak()
}

// Assertion based on type using a switch statement
func callingHuman(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into callingHuman().", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into callingHuman().", h.(secretAgent).first)
	}
	fmt.Println("I was passed into callingHuman().", h)
}

// This is used in conversion below
type hamburger int

func main() {

	bar("Pittsburgh")
	// Function syntax is as follows:
	// func (r receiver) identifier(parameters) (returns) { code }

	s1 := woo("Allentown")
	fmt.Println(s1)

	x, y := mouse("Mac", "Miller")
	fmt.Println(x)
	fmt.Println(y)

	// Using the "defer" keyword
	// Right when func main() exits, that's whenever the "deferred" stuff runs
	defer foo()
	blah()

	secretAgent1 := secretAgent{
		person: person{
			"Alex",
			"Rider",
		},
		licenseToKill: false,
	}

	fmt.Println(secretAgent1)
	secretAgent1.speak()

	person1 := person{
		first: "Sabina",
		last:  "Pleasure",
	}

	fmt.Println(person1)

	// callingHuman() can take in both type "secretAgent" and type "person" because both "secretAgent" and "person" are ALSO of type "human", which callingHuman() accepts.
	// This is polymorphism, because callingHuman() the function can take in many different types.
	callingHuman(secretAgent1)
	callingHuman(person1)

	// Conversion --- can convert from one type to another.
	var a hamburger = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	var b int
	b = int(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// This is an anonymous function without arguments
	func() {
		fmt.Println("Anonymous func ran.")
	}() // The arguments would be passed in here.

	// This is an anonymous function with arguments
	func(x int) {
		fmt.Println("Brendan Shanahan's Red Wings number was:", x)
	}(14) // Arguments are passed in here.

	// This is a func expression that assigns the function to a variable
	result := func() {
		fmt.Println("My first func expression.")
	}
	result() // Calls the function above

	// Returning a function from a function
	result2 := returnAString()
	fmt.Println(result2)

}

// This function is called immediately above
func returnAString() string {
	s := "Hello World"
	return s
}

// Sample function
func bar(s string) {
	fmt.Println("Hello,", s)
}

// Sample function with return statement
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

// Sample function with multiple parameters
func mouse(firstName string, lastName string) (string, bool) {
	a := fmt.Sprint(firstName, " ", lastName, ` says "Hello."`)
	b := false
	return a, b
}

func foo() {
	fmt.Println("foo")
}

func blah() {
	fmt.Println("blah")
}

package main

import (
	"fmt"
	"sort"
)

type FrazzCharacter struct {
	FirstName  string
	Occupation string
}

type ByName []FrazzCharacter

func (byName ByName) Len() int           { return len(byName) } // gets the length of byName
func (byName ByName) Swap(i, j int)      { byName[i], byName[j] = byName[j], byName[i] }
func (byName ByName) Less(i, j int) bool { return byName[i].FirstName < byName[j].FirstName }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"Frazz", "Jane", "Mrs. Olsen", "Mrs. Burke", "Caulfield"}

	fmt.Println(xi)
	fmt.Println(xs)

	// How to sort these slices?
	// Alphabetical order
	sort.Ints(xi)
	fmt.Println(xi)
	// Increasing order
	sort.Strings(xs)
	fmt.Println(xs)

	frazz := FrazzCharacter{
		"Frazz",
		"Custodian",
	}

	jane := FrazzCharacter{
		"Jane Plainwell",
		"Teacher",
	}

	burke := FrazzCharacter{
		"Mr. Burke",
		"Teacher",
	}

	caulfield := FrazzCharacter{
		"Caulfield",
		"Student",
	}

	characters := []FrazzCharacter{frazz, jane, burke, caulfield}

	fmt.Println(characters)

	// Sort "characters" by name
	sort.Sort(ByName(characters))
	fmt.Println(characters)

}

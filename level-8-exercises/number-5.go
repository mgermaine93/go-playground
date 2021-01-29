// Sort the []user by age and last.
// Also sort each []string "Sayings" for each user
// Print everything in a way that's pleasant

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

//////////////////// TAKEN FROM https://pkg.go.dev/sort ////////////////////
// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("********************** SAYINGS ARE ALPHABETIZED *************************")
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		// Sort their sayings in alphabetical order
		sort.Strings(user.Sayings)
		for _, saying := range user.Sayings {
			fmt.Println("\t", saying)
		}
	}

	fmt.Println("********************** BY AGE *************************")

	// Sorts users by age
	sort.Sort(ByAge(users))
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Println("\t", saying)
		}
	}

	fmt.Println("*********************** BY LAST NAME ************************")

	sort.Sort(ByLast(users))
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Println("\t", saying)
		}
	}

	// // Sort "users" by age
	// sort.Sort(ByAge(users))
	// fmt.Println(users)

	// // Sort "users" by last name
	// sort.Sort(ByLast(users))
	// fmt.Println(users)

}

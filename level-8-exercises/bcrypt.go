// Seeing how bcrypt works

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `pssword123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	somePasswordHere := `pssword123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(somePasswordHere))
	if err != nil {
		fmt.Println("You can't login!", err)
		return
	}

	fmt.Println("You're logged in!")

}

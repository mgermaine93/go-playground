package main

import (
	"fmt"

	"github.com/mgermaine93/learn-to-go/level-13-exercises/example-tests/dmb"
)

func main() {
	fmt.Println(dmb.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(dmb.Sum(1, 2, 3, 4, 5))
}

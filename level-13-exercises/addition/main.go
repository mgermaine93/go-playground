package addition

import "fmt"

func addOneAndOne() int {
	sum := 1 + 1
	return sum
}

func main() {
	returnedValue := addOneAndOne()
	fmt.Println(returnedValue)
}

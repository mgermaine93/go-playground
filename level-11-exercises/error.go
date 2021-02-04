package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// prints out what happened
		fmt.Println("error happened: ", err)
		// gets data/time stamp
		log.Println("error happened: ", err)
		// exits with a given status code, terminating the program
		log.Fatal(err)
		// deferred functions run, but you can use "recover" to resume
		log.Panicln(err)

	}
}

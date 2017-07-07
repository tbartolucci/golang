package main

import (
	"fmt"
	"os"
)

func main() {
	var message string
	message = "Hello, I am Gopher"
	// type inference
	// args := os.Args
	var args []string
	args = os.Args

	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println(message)
	}
}


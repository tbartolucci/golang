package main

import "fmt"

func main() {
	gopher1 := gopher{name: "Phil", age: 30}
	gopher2 := gopher{name: "Noodles", age: 90}

	fmt.Println(gopher1.jump())
	fmt.Println(gopher2.jump())
}

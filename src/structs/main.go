package main

import "fmt"

func main() {
	gopher1 := &gopher{name: "Phil", age: 30}
	gopher2 := gopher{name: "Noodles", age: 90}

	fmt.Println(gopher1.jump())
	fmt.Println(gopher2.jump())

	validateAge(gopher1)
	fmt.Println(gopher1)

	language := "go"
	favoriteLanguage := language // copies the value to new value

	language1 := "go"
	favoriteLanguage1 := &language1 // a single memory address is used

	fmt.Println(favoriteLanguage + *favoriteLanguage1)
}

func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}
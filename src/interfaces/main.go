package main

import "fmt"

func main() {
	jumperList := getList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.jump())
	}
}

// Interfaces are implemented implicitly by implementing the methods
// convention in go, one method interface method name + "er"
type jumper interface {
	jump() string
}

// return a slice of pointers to gophers
func getGopherList() []*gopher {
	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 90}

	list := []*gopher{phil, noodles}
	return list
}

func getList() []jumper {
	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 90}
	barbaro := &horse{name: "Barbaro", weight: 2000}

	list := []jumper{phil, noodles, barbaro}
	return list
}
package main

type gopher struct {
	name string
	age int
	isAdult bool  // default is false
}

func (g gopher) jump() string {
	if g.age < 65 {
		return g.name + " can jump HIGH"
	}
	return g.name + " can still jump"
}
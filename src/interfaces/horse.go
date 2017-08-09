package main

type horse struct {
	name string
	weight int
}

func (h horse) jump() string {
	if h.weight > 2500 {
		return "I'm too heavy, can't jump..."
	}
	return "I will jump, neigh!!"
}
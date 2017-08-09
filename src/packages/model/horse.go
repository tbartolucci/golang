package model

type horse struct {
	name string
	weight int
}

// Capitalizing the method "exports" it from the package.
func (h horse) Jump() string {
	if h.weight > 2500 {
		return "I'm too heavy, can't jump..."
	}
	return "I will jump, neigh!!"
}
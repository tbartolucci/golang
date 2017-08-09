package model

// Interfaces are implemented implicitly by implementing the methods
// convention in go, one method interface method name + "er"
type jumper interface {
	Jump() string
}

func GetList() []jumper {
	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 90}
	barbaro := &horse{name: "Barbaro", weight: 2000}

	list := []jumper{phil, noodles, barbaro}
	return list
}
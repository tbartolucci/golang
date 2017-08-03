package main

import "fmt"

func main() {
	var langs [3]string

	langs[0] = "Go"
	langs[1] = "Ruby"
	langs[2] = "Javascript"
	// langs[3] = "BAD CODE" // Will cause compile error

	fmt.Println(langs)

	// "Slice"
	var slicelang []string

	slicelang = append(slicelang, "Go")
	slicelang = append(slicelang, "Ruby")
	slicelang = append(slicelang, "Javascript")
	slicelang = append(slicelang, "No Longer Bad code")

	fmt.Println(slicelang)

	newLangs := getLangs()

	for i := range newLangs {
		fmt.Println(langs[i])
	}

	// for i, element := range newLangs { // compile error, unused variable
	for _, element := range newLangs {
		fmt.Println(element)
	}
}

func getLangs() []string {
	langs := []string{"Go", "Ruby", "Javscript"}

	return langs
}

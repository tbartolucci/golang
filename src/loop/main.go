package main

import "fmt"

func main() {
	i := 0
	isLessThanFive := true
	for isLessThanFive {
		if i >= 5 {
			isLessThanFive = false
		}
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}
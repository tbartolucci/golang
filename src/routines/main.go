package main

//GOMAXPROCS=1 will force only a single CPU

import (
	"fmt"
	"sync"
	"math"
)

func main() {
	names := []string{"Phil", "Noodles", "Barbaro"}
	var wg sync.WaitGroup
	wg.Add(len(names)) // wait for a number of go routines to finish
	for _, name := range names {
		go printName(name, &wg)
	}
	wg.Wait()  // don't exit the program until the wg is complete
}

func printName(n string, wg *sync.WaitGroup){
	result := 0.0
	for i := 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(n)))
	}
	fmt.Println("Name: ", n)
	wg.Done()
}
package main

import (
	"fmt"
	"packages/model"
)

func main() {
	jumperList := model.GetList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}

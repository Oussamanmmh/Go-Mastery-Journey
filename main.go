package main

import (
	"fmt"
	"math"
)

func myFunction(n string) {
	fmt.Println("the name is", n)
	n = "aymen"
}
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

var score int = 300

func main() {
	// var name string = "oussama"
	// myFunction(name)
	// fmt.Println("the name now:", name)
	// fmt.Println("the area :", circleArea(4))
	Greeting()
	displayScore()
}

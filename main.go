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
	// Greeting()
	// displayScore()
	//maps
	menu := map[string]float64{
		"soup":   4.99,
		"pie":    7.99,
		"salad":  6.99,
		"coffee": 1.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
	//looping through maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//deleting from maps
	delete(menu, "pie")
	fmt.Println(menu)
	//replace value
	menu["coffee"] = 2.99

	student1.displayInfo()

}

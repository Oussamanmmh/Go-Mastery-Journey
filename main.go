package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
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

	// student1.displayInfo()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age:")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age) // Trim the newline characters
	p, err := strconv.ParseFloat(age, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please enter a valid number")
	}
	fmt.Print("the age ", p)
	//age = strings.TrimSpace(age)

	fmt.Print("the age ", reflect.TypeOf(p))

}

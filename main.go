// Description: This file contains the code for the first part of the tutorial which is about variables in Go.

package main // Package declaration which is the main package

import (
	"fmt" // Importing the fmt package which is used for printing to the console it stands for format
	"sort"
)

// Main function
func main() {
	//Strings
	// var nameOne string = "OUSSAMA"
	// var nameTwo = "YASSINE"
	// var nameThree string //Default value is ""

	// nameThree = "AMINE" //Assigning a value to nameThree

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "IMANE" //Short hand declaration

	// fmt.Println(nameFour)

	// //Integers
	// var ageOne int = 20 //Explicit declaration
	// var ageTwo = 30
	// ageThree := -40 //Short hand declaration
	// fmt.Println(ageOne, ageTwo, ageThree)

	// //Bits & Memory
	// var ageFour int8 = 127  // 8 bits for signed numbers range [-128, 127]
	// var ageFive uint8 = 255 // 8 bits for positive numbers only range [0, 255]
	// fmt.Println(ageFour, ageFive)

	// //Floats
	// var scoreOne float32 = 25.98 // 32 bits for floating point numbers
	// var scoreTwo float64 = 25.98 // 64 bits for floating point numbers
	// fmt.Println(scoreOne, scoreTwo)

	// //fmt package
	// fmt.Printf("ageFour: %v, ScoreTwo: %f\n", ageFour, scoreTwo)
	// fmt.Printf("nameOne:%s , nameTwo:%q\n", nameOne, nameTwo)
	// fmt.Printf(("The type of nameOne is %T\n"), nameOne)
	// //Sprintf (save formatted strings)
	// formattedString := fmt.Sprintf("nameOne:%s , nameTwo:%q\n", nameOne, nameTwo)
	// print(formattedString)

	// var ages [3]int = [3]int{20, 30, 40} //Array of integers fix length
	// var ages2 = [3]int{20, 30, 40}
	// names := [4]string{"Oussama", "Yassine", "Amine", "Imane"} //Array of strings fix length

	// //slices (dynamic arrays)
	// var scores = []int{100, 200, 300, 400, 500}

	// scores = append(scores, 500) //Adding an element to the slice
	// scores[2] = 25

	// fmt.Println(len(ages), ages2, names, len(scores))

	// //slices ranges
	// rangeOne := scores[:3] // [200, 300]
	// fmt.Print(rangeOne)
	// greeting := "Hello there friends!"
	// fmt.Println(strings.Contains(greeting, "Hello"))                       //returns true if the strings containes the substring else false
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "oussama nemamcha")) //Does not replace the original one
	// fmt.Println(strings.ToUpper(greeting))                                 //returns new string
	// fmt.Println(strings.Index(greeting, "0"))                              //return -1 if the substring does not exist
	// fmt.Println(strings.Split(greeting, " "))
	numbers := []int{23, 32, 56, 1, 2, 100}
	sort.Ints(numbers)
	fmt.Println(numbers)
	index := sort.SearchInts(numbers, 203)
	println(index)

}

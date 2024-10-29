// Description: This file contains the code for the first part of the tutorial which is about variables in Go.

package main // Package declaration which is the main package

import "fmt" // Importing the fmt package which is used for printing to the console it stands for format

// Main function
func main() {
	//Strings
	var nameOne string = "OUSSAMA"
	var nameTwo = "YASSINE"
	var nameThree string //Default value is ""

	nameThree = "AMINE" //Assigning a value to nameThree

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "IMANE" //Short hand declaration

	fmt.Println(nameFour)

	//Integers
	var ageOne int = 20 //Explicit declaration
	var ageTwo = 30
	ageThree := -40 //Short hand declaration
	fmt.Println(ageOne, ageTwo, ageThree)

	//Bits & Memory
	var ageFour int8 = 127  // 8 bits for signed numbers range [-128, 127]
	var ageFive uint8 = 255 // 8 bits for positive numbers only range [0, 255]
	fmt.Println(ageFour, ageFive)

	//Floats
	var scoreOne float32 = 25.98 // 32 bits for floating point numbers
	var scoreTwo float64 = 25.98 // 64 bits for floating point numbers
	fmt.Println(scoreOne, scoreTwo)
}

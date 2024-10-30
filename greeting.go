package main

import "fmt"

//object
type student struct {
	name    string
	age     int
	college string
}

var student1 = student{
	name:    "oussama",
	age:     22,
	college: "esi",
}

//reciever function
//this function is associated with the student object
func (s student) displayInfo() {
	fmt.Printf("the student name is %s and he is %d years old and he is studying at %s\n", s.name, s.age, s.college)
}

func Greeting() {
	fmt.Println("Hello, World!")
}

func displayScore() {
	fmt.Println("Score is ", score)
}

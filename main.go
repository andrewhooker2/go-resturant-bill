package main

import "fmt"

func main() {

	// String  declatration examples
	fmt.Println("Hello, friends!")
	var nameOne string = "Andy"
	var nameTwo = "Danny"

	// This will just store and empty string for now
	var nameThree string

	// You can declare varibles like this "in" functions
	nameFour := "Alaina"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// Ints
	var ageOne int = 27
	var ageTwo int = 25
	ageThree := 29

	fmt.Println(ageOne, ageTwo, ageThree)

	// Floats ( you must say what the size of the float you want to use is )
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 8881245621

	// unless you use this ( it sets it to 64 any ways)
	scoreThree := 123456.123456

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}

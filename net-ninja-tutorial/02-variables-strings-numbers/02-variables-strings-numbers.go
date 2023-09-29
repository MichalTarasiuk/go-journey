package main

import "fmt"

// ok
// var someName = "hello"
// error
// someName := "hello"

func main() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	// can't change the type
	// nameOne = 23

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)

	// ints
	var age int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(age, ageTwo, ageThree)

	// bits & memory
	var numberOne int8 = 25
	var numberTwo int8 = -128
	var numberThree int16 = 256

	fmt.Println(numberOne, numberTwo, numberThree)

	// floats
	var scoreOne float32 = 25.90
	var scoreTwo float64 = 39028903829083902.32

	fmt.Println(scoreOne, scoreTwo)

}

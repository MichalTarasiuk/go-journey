package main

import "fmt"

// Notes:
// zero value: ints -> 0, bools -> false, string -> ""
// GO creates a copy of variabe when when it' s passed around as variabe

/*
|---var---|-pointer-|
|  0x001  |  0x002  |
|---------|---------|
| "value" | p0x001  |
|---------|---------|
*/

func updateName(name *string) {
	*name = "updated name"
}

func main() {
	fmt.Println("Hello World")

	// var name type = expression
	var integer int = 1
	fmt.Println(integer)

	// zero value
	var integer2 int
	// output: 0
	fmt.Println(integer2)

	// multiple variables
	// 1:
	var integer3, integer4 int
	fmt.Println(integer3, integer4)

	// 2:
	var integer5, integer6 int = 5, 6
	fmt.Println(integer5, integer6)

	// 3:
	var integer7, string = 7, "string"
	fmt.Println(integer7, string)

	// short declaration
	// name := expression
	boolean := true
	fmt.Println(boolean)

	// pointers
	name := "name"
	updateName(&name)
	// updated name
	fmt.Println(name)

	// type declarations
	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	// error: cannot use f (variable of type fahrenheit) as celsius value in assignment
	// c = f

	c = celsius((f - 32) * 5 / 9)

	fmt.Println(f, c)
}

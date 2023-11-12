package main

import (
	"fmt"
	"golang-dojo/03-scope/integers"
)

// Summary notes:
// - Go has three main levels of scopes: blocks, packages, and the universe (the built-in functions and types).
// - Unlike some other languages, Go does not hoist variables; they must be declared before use within the same scope.
// - Variables defined at the package level can be shared between functions in that package. Variables declared outside functions are shared between packages as long as the variable is defined outside a function.
// - If a variable should be accessed from outside its package, its name must start with an uppercase letter.

var five = 5

func main() {
	// package scope
	fmt.Println(helloWorld)

	// function scope = block scope
	var integer = 1

	fmt.Println(integer)

	{
		// block scope
		var integer = 2

		fmt.Println(integer)
	}

	{
		// block scope
		integer2 := 3

		fmt.Println(integer2)
	}

	// universe scope
	fmt.Println(integers.UniverseThree)
}

func nonMain() {
	// package scope
	fmt.Println(five)
}

package main

import "fmt"

func main() {
	v := 42 // change me!
	// v = "Hello World", cannot use "" (untyped string constant) as int value in assignmentcompilerIncompatibleAssign
	fmt.Printf("v is of type %T\n", v)
}

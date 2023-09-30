package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	// error
	// updateName(name)
	// fmt.Println("memory address of name is: ", &name)

	m := &name
	fmt.Println("memory address: ", m)
	fmt.Println("value at memory address: ", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}

// |---name---|--pointer--|
// |   0xe01  |  0x002    |
// |   "tifa" |  pêx001   |

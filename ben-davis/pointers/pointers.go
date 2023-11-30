package main

import "fmt"

func main() {
	y := 23
	x := &y

	fmt.Println(y == **&x)

	updateYWithPtr(x)
	updateYNoPtr(y)

	fmt.Println(y)
}

func updateYWithPtr(y *int) {
	*y = 10
}

func updateYNoPtr(y int) {
	y = 20
}

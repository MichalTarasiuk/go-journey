package functions

import "fmt"

func add(x int, y int) int {
	return x + y
}

func Example() {
	fmt.Println(add(42, 13))
}
  
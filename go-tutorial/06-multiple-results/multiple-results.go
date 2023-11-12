package multipleResults

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func Example() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

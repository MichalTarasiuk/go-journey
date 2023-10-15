package main

import "fmt"

func main() {
	var x = 1 // int32 or int64 depending on machine
	var x32 int32 = 1
	var x64 int64 = 1

	fmt.Println(x, x32, x64)
	fmt.Printf("%T, %T, %T", x, x32, x64)
}

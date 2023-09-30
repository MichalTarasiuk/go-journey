package main

import "fmt"

func main() {
	mybill := makeBill("mario's bill")

	fmt.Println(mybill.format())
}

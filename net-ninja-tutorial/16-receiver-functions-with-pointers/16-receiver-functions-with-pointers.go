package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")

	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 8.95)
	mybill.addItem("toffe puding", 4.95)
	mybill.addItem("coffe", 3.25)

	mybill.updateTip(10)

	fmt.Println(mybill.format())
}

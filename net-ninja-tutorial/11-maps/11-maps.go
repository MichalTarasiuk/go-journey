package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffie pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		312899432: "mario",
		432432423: "luigi",
		492304348: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[312899432])

	phonebook[312899432] = "bowser"

	fmt.Println(phonebook)
}

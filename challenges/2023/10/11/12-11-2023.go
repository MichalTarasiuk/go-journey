package main

import "fmt"

func main() {
	list := []string{"apple", "banana", "cherry"}
	myMap := map[string]bool{}

	for _, item := range list {
		myMap[item] = true
	}

	fmt.Println(myMap)
}

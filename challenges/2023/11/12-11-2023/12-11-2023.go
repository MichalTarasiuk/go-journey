package main

import (
	"fmt"
	"slices"
)

func arrToMap(arr []string) map[string]bool {
	myMap := map[string]bool{}

	for _, item := range arr {
		myMap[item] = true
	}

	return myMap
}

func findUniqueElements(arr []int) []int {
	uniqueElements := []int{}

	for _, item := range arr {
		if !slices.Contains(uniqueElements, item) {
			uniqueElements = append(uniqueElements, item)
		}
	}

	return uniqueElements
}

func createDynamicMap[Key comparable, Value any](keys []Key, value Value) map[Key]Value {
	myMap := map[Key]Value{}

	for _, key := range keys {
		myMap[key] = value
	}

	return myMap
}

func groupBy[Item, Key comparable](items []Item, getKey func(Item) Key) map[Key][]Item {
	result := make(map[Key][]Item)

	for _, item := range items {
		key := getKey(item)
		result[key] = append(result[key], item)
	}

	return result
}

func main() {
	fmt.Println(arrToMap([]string{"apple", "banana", "cherry"}))
	fmt.Println(findUniqueElements([]int{1, 2, 3, 2, 4, 5, 6, 4}))
	fmt.Println(createDynamicMap[int, string]([]int{1, 2, 3, 4, 5}, "value"))

	type Person struct {
		Name string
		Age  int
	}

	grouped := groupBy([]Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
		{"David", 30},
		{"Eva", 35},
	}, func(p Person) int {
		return p.Age
	})

	fmt.Println(grouped)
}

package main

import "fmt"

func getMapKeys[Key comparable](map1 map[Key]any) []Key {
	keys := []Key{}

	for key := range map1 {
		keys = append(keys, key)
	}

	return keys
}

func commonKeys[Key comparable](map1 map[Key]any, map2 map[Key]any) []Key {
	keys1 := getMapKeys(map1)
	commonKeys := []Key{}

	for _, key1 := range keys1 {
		if _, exists := map2[key1]; exists {
			commonKeys = append(commonKeys, key1)
		}
	}

	return commonKeys
}

func truthCheckCollection[Item any](items []Item, assert func(Item) bool) bool {
	for _, item := range items {
		if !assert(item) {
			return false
		}
	}

	return true
}

func main() {
	map1 := map[string]any{"a": 1, "b": 2, "c": 3}
	map2 := map[string]any{"b": 2, "c": 3, "d": 4}

	commonKeys := commonKeys(map1, map2)

	fmt.Println("Common keys:", commonKeys)

	type User struct {
		Name string
		Age  int
	}

	users := []User{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 22},
	}

	allAdults := truthCheckCollection(users, func(u User) bool {
		return u.Age >= 18
	})

	fmt.Println("Are all users adults?", allAdults)
}

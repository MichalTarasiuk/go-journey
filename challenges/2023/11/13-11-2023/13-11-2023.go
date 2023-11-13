package main

import (
	"fmt"
)

func renameKeys(map1 map[string]string, map2 map[string]string) map[string]string {
	newMap := make(map[string]string)

	for key, value := range map1 {
		if newKey, ok := map2[key]; ok {
			newMap[newKey] = value
		} else {
			newMap[key] = value
		}
	}

	return newMap
}

func copyMap[Key comparable, Value any](originalMap map[Key]Value) map[Key]Value {
	copiedMap := make(map[Key]Value)

	for key, value := range originalMap {
		copiedMap[key] = value
	}

	return copiedMap
}

func unwind[Key comparable](map1 map[Key]interface{}, key Key) []map[Key]interface{} {
	newArr := []map[Key]interface{}{}

	if slice, ok := map1[key].([]interface{}); ok {
		for _, item := range slice {
			copiedMap := copyMap(map1)
			copiedMap[key] = item

			newArr = append(newArr, copiedMap)
		}
	}

	return newArr
}

func main() {
	map1 := map[string]string{
		"name": "John",
		"age":  "25",
		"city": "New York",
	}
	mapping := map[string]string{
		"name": "full_name",
		"age":  "years",
	}
	fmt.Println(renameKeys(map1, mapping))

	fmt.Println(unwind(map[string]interface{}{
		"a": 1, "b": []interface{}{2, 3},
	}, "b"))
}

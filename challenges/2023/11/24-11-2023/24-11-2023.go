package main

import (
	"challenges/lib"
	"fmt"
)

func main() {
	var sum int
	for _, line := range lib.AocInputLines(2022, 3) {
		lib.Assertf(len(line)%2 == 0, "Line %q has %d items", line, len(line))

		first := make(map[rune]int)
		for _, rune := range line[:len(line)/2] {
			first[rune]++
		}

		second := make(map[rune]int)
		for _, rune := range line[len(line)/2:] {
			second[rune]++
		}

		both := lib.Intersect(first, second)
		lib.Assertf(len(both) == 1, "Line %q does not have exactly one common element: %v", line, both)

		sum += priority(lib.MapSomeKey(both))
	}

	fmt.Println(sum)
}

func priority(item rune) int {
	switch {
	case item >= 'a' && item <= 'z':
		return int(item-'a') + 1
	case item >= 'A' && item <= 'Z':
		return int(item-'A') + 27
	default:
		lib.Panicf("Invalid item %q", item)
		return 0
	}
}

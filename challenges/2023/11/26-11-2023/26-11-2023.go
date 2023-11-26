package main

import (
	"challenges/lib"
	"fmt"
	"regexp"
)

func extractNumbers(input string) ([]string, error) {
	pattern := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	matches := pattern.FindStringSubmatch(input)

	if len(matches) == 0 {
		return nil, fmt.Errorf("no matches found")
	}

	return matches[1:], nil
}

func main() {
	var enclosed, overlapping int
	for _, line := range lib.AocInputLines(2022, 4) {
		numbers, error := extractNumbers(line)

		if error != nil {
			panic("no numbers found")
		}

		var (
			a1 = numbers[0]
			a2 = numbers[1]
			b1 = numbers[2]
			b2 = numbers[3]
		)

		if (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2) {
			enclosed++
		}
		if (a1 <= b1 && a2 >= b1) || (a1 <= b2 && a2 >= b2) || (a1 >= b1 && a2 <= b2) {
			overlapping++
		}
	}
	fmt.Println(enclosed)
	fmt.Println(overlapping)
}

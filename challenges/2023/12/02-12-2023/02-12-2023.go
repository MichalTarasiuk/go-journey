package main

import (
	"challenges/lib"
	"fmt"
	"strconv"
)

func main() {
	result1 := 0
	for _, line := range lib.AocInputLines(2023, 01) {
		ints := lib.ExtractIndividualInts(line)

		first := ints[0]
		last := ints[len(ints)-1]

		final := fmt.Sprintf("%d%d", first, last)
		finalInt, err := strconv.Atoi(final)

		if err != nil {
			continue
		}

		result1 += finalInt
	}

	fmt.Println(result1)
}

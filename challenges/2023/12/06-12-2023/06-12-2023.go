package main

import (
	"challenges/lib"
	"fmt"
	"strings"
)

func main() {
	lines := lib.AocInputLines(2023, 06)
	lib.Assert(len(lines) == 2)
	times := lib.ExtractInts(lines[0])
	records := lib.ExtractInts(lines[1])
	lib.Assert(len(times) == len(records))
	fmt.Println(compute(times, records))

	times2 := []int{lib.ExtractInt(strings.ReplaceAll(lines[0], " ", ""))}
	records2 := []int{lib.ExtractInt(strings.ReplaceAll(lines[1], " ", ""))}
	fmt.Println(compute(times2, records2))
}

func compute(times, records []int) int {
	result := 1
	for index, time := range times {
		var multiplier = 0
		for speed := 1; speed <= time; speed++ {
			distance := (time - speed) * speed
			if distance > records[index] {
				multiplier++
			}
		}
		lib.Assert(multiplier != 0)
		result *= multiplier
	}
	return result
}

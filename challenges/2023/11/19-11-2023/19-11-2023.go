package main

import (
	"challenges/lib"
	"fmt"
	"sort"
	"strings"
)

func solvePart01() (int, int) {
	var sums []int
	for _, pg := range lib.AocInputParagraphs(2022, 1) {
		sums = append(sums, lib.Sum(lib.ExtractInts(strings.Join(pg, " "))...))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return sums[0], sums[0] + sums[1] + sums[2]
}

func main() {
	elfCarryingTheMostCalories, totalCalories := solvePart01()
	fmt.Printf("total calories: %v, elf carrying the most calories: %v", totalCalories, elfCarryingTheMostCalories)
}

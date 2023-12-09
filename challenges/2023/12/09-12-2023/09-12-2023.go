package main

import (
	"challenges/lib"
	"fmt"
	"sort"
)

func main() {
	lines := lib.AocInputLines(2023, 9)
	sort.Strings(lines)

	var sum int
	for _, line := range lines {
		history := lib.ExtractInts(line)
		sum += predicateHistoryValue(history)
	}
	fmt.Println(sum)
}

func predicateHistoryValue(lastSequence []int) int {
	sequences := [][]int{lastSequence}
	currentSequence := lastSequence

	for !lib.All(currentSequence, func(i int) bool { return i == 0 }) {
		var newSequence []int
		for index, value := range currentSequence {
			nextIndex := index + 1
			if nextIndex <= len(currentSequence)-1 {
				newSequence = append(newSequence, currentSequence[nextIndex]-value)
			}
		}
		currentSequence = newSequence
		sequences = lib.Shift(sequences, currentSequence)
	}

	var historyValue int
	for _, sequence := range sequences {
		historyValue = -historyValue + sequence[0]
	}
	return historyValue
}

package main

import (
	"challenges/lib"
	"fmt"
	"sort"
)

func main() {
	lines := lib.AocInputLines(2023, 9)
	sort.Strings(lines)

	var sum1 int
	for _, line := range lines {
		history := lib.ExtractInts(line)
		sum1 += predicateHistoryValue(history)
	}
	fmt.Println(sum1)
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
		sequences = append(sequences, currentSequence)
	}

	var historyValue int
	for _, sequence := range sequences {
		historyValue += sequence[len(sequence)-1]
	}
	return historyValue
}

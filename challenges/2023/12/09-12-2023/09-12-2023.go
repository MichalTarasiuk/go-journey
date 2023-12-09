package main

import (
	"challenges/lib"
	"fmt"
	"sort"
)

func main() {
	lines := lib.AocInputLines(2023, 9)
	sort.Strings(lines)

	var (
		sum1, sum2 int
	)
	for _, line := range lines {
		history := lib.ExtractInts(line)
		sum1 += predicateHistoryValue(history, 1)
		sum2 += predicateHistoryValue(history, -1)
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func predicateHistoryValue(lastSequence []int, move int) int {
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
		if move == 1 {
			historyValue += sequence[len(sequence)-1]
		} else {
			historyValue = -historyValue + sequence[0]
		}
	}
	return historyValue
}

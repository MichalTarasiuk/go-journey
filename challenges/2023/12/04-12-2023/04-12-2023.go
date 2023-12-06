package main

import (
	"challenges/lib"
	"fmt"
)

func findCommonElements(scratchcard string) []int {
	var winningNumbersSubmatch string
	var myNumbersSubmatch string
	lib.Extract(scratchcard, `^Card\s+\d+:(.+)\|(.+)$`, &winningNumbersSubmatch, &myNumbersSubmatch)

	winningNumbers := lib.ExtractInts(winningNumbersSubmatch)
	myNumbers := lib.ExtractInts(myNumbersSubmatch)

	return lib.FindCommonElements(winningNumbers, myNumbers)
}

var scoreCache = make(map[string]int)

func getScratchcardScore(scratchcard string) int {
	if score, ok := scoreCache[scratchcard]; ok {
		return score
	}

	var score int
	for range findCommonElements(scratchcard) {
		if score == 0 {
			score = 1
		} else {
			score *= 2
		}
	}
	scoreCache[scratchcard] = score
	return score
}

var totalScratchcardsCache = make(map[string]int)

func getTotalScratchcards(input map[int]string, scratchcard string, scratchcardIndex int) int {
	if score, ok := totalScratchcardsCache[scratchcard]; ok {
		return score
	}

	var result int
	commonElements := findCommonElements(scratchcard)

	if len(commonElements) == 0 {
		return result + 1
	}

	for index := scratchcardIndex + 1; index <= scratchcardIndex+len(commonElements); index++ {
		if line, ok := input[index]; ok {
			result += getTotalScratchcards(input, line, index)
		}
	}

	result++
	totalScratchcardsCache[scratchcard] = result

	return result
}

func main() {
	input := lib.SliceToMap(lib.AocInputLines(2023, 04))

	var totalScore int
	var totalScratchcards int
	for index, scratchcard := range input {
		totalScore += getScratchcardScore(scratchcard)
		totalScratchcards += getTotalScratchcards(input, scratchcard, index)
	}
	fmt.Println(totalScore)
	fmt.Println(totalScratchcards)
}

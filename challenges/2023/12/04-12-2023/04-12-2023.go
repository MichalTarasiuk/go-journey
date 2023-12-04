package main

import (
	"challenges/lib"
	"fmt"
)

func main() {
	var result int
	for _, line := range lib.AocInputLines(2023, 04) {
		var winningNumbersSubmatch string
		var myNumbersSubmatch string
		lib.Extract(line, `^Card\s+\d+:(.+)\|(.+)$`, &winningNumbersSubmatch, &myNumbersSubmatch)

		winningNumbers := lib.ExtractInts(winningNumbersSubmatch)
		myNumbers := lib.ExtractInts(myNumbersSubmatch)

		var score int
		for range lib.FindCommonElements(winningNumbers, myNumbers) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
		result += score
	}
	fmt.Println(result)
}

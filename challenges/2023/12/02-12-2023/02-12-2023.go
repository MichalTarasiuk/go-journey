package main

import (
	"challenges/lib"
	"fmt"
	"regexp"
	"strconv"
)

func solvePart1(inputLines []string) int {
	resultSum := 0
	for _, line := range lib.AocInputLines(2023, 01) {
		ints := lib.ExtractIndividualInts(line)

		first := ints[0]
		last := ints[len(ints)-1]

		final := fmt.Sprintf("%d%d", first, last)

		if finalInt, err := strconv.Atoi(final); err == nil {
			resultSum += finalInt
		}
	}
	return resultSum
}

var numberAndTextRegex = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)

var textToNumberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func extractIntegersFromString(line string) []int {
	integers := []int{}
	for _, str := range numberAndTextRegex.FindAllString(line, -1) {
		if num, err := strconv.Atoi(str); err == nil {
			integers = append(integers, num)
		} else if mappedNum, ok := textToNumberMap[str]; ok {
			integers = append(integers, mappedNum)
		}
	}
	return integers
}

func solvePart2(inputLines []string) int {
	resultSum := 0
	for _, inputLine := range inputLines {
		integers := extractIntegersFromString(inputLine)

		firstInteger := integers[0]
		lastInteger := integers[len(integers)-1]

		concatenatedString := fmt.Sprintf("%d%d", firstInteger, lastInteger)
		concatenatedInteger, err := strconv.Atoi(concatenatedString)

		if err != nil {
			continue
		}

		resultSum += concatenatedInteger
	}
	return resultSum
}

func main() {
	inputLines := lib.AocInputLines(2023, 01)

	fmt.Println(solvePart1(inputLines))
	fmt.Println(solvePart2(inputLines))
}

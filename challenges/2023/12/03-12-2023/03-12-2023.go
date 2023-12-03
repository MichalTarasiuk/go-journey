package main

import (
	"challenges/lib"
	"fmt"
	"strings"
	"unicode"
)

func hasSymbol(s string) bool {
	return strings.ContainsFunc(s, func(r rune) bool {
		return !unicode.IsNumber(r) && !(string(r) == ".")
	})
}

func solvePart1(input []string) int {
	result := 0
	inputMap := lib.SliceToMap(input)
	for index, line := range input {
		for _, intWithIndex := range lib.ExtractPositiveIntsWithIndex(line) {
			startIndex := lib.Max(intWithIndex.Index-1, 0)
			endIndex := lib.Min(startIndex+len(fmt.Sprintf("%d", intWithIndex.Value))+2, len(line))

			if lineBefore, ok := inputMap[index-1]; ok && hasSymbol(lineBefore[startIndex:endIndex]) {
				result += intWithIndex.Value
				continue
			}

			if hasSymbol(line[startIndex:endIndex]) {
				result += intWithIndex.Value
				continue
			}

			if lineAfter, ok := inputMap[index+1]; ok && hasSymbol(lineAfter[startIndex:endIndex]) {
				result += intWithIndex.Value
				continue
			}
		}
	}
	return result
}

func solvePart2(input []string) int {
	result := 0

	return result
}

func main() {
	input := lib.AocInputLines(2023, 03)

	fmt.Println(solvePart1(input))
	fmt.Println(solvePart2(input))
}

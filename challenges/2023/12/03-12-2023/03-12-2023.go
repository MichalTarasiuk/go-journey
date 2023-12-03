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

func main() {
	result := 0
	input := lib.SliceToMap(lib.AocInputLines(2023, 03))
	for index, line := range input {
		for _, intWithIndex := range lib.ExtractIntsWithIndex(line) {
			startIndex := lib.Max(intWithIndex.Index-1, 0)
			endIndex := lib.Min(startIndex+len(fmt.Sprintf("%d", intWithIndex.Value))+2, len(line))

			if lineBefore, ok := input[index-1]; ok && hasSymbol(lineBefore[startIndex:endIndex]) {
				result += intWithIndex.Value
				continue
			}

			if hasSymbol(line[startIndex:endIndex]) {
				result += intWithIndex.Value
				continue
			}

			if lineAfter, ok := input[index+1]; ok && hasSymbol(lineAfter[startIndex:endIndex]) {
				result += intWithIndex.Value
				continue
			}
		}
	}
	fmt.Println(result)
}

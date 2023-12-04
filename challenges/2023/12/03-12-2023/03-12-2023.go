package main

import (
	"challenges/lib"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type IntRange struct {
	startIndex int
	endIndex   int
}

func getIntRange(intWithIndex lib.IntWithIndex, line string) IntRange {
	startIndex := lib.Max(intWithIndex.Index-1, 0)

	return IntRange{
		startIndex: startIndex,
		endIndex:   lib.Min(startIndex+len(fmt.Sprintf("%d", intWithIndex.Value))+2, len(line)),
	}
}

func hasSymbol(s string) bool {
	return strings.ContainsFunc(s, func(r rune) bool {
		return !unicode.IsNumber(r) && !(string(r) == ".")
	})
}

func solvePart1(input []string) int {
	result := 0
	inputMap := lib.SliceToMap(input)
	for index, line := range inputMap {
		for _, intWithIndex := range lib.ExtractPositiveIntsWithIndex(line) {
			intRange := getIntRange(intWithIndex, line)

			if lineBefore, ok := inputMap[index-1]; ok && hasSymbol(lineBefore[intRange.startIndex:intRange.endIndex]) {
				result += intWithIndex.Value
				continue
			}

			if hasSymbol(line[intRange.startIndex:intRange.endIndex]) {
				result += intWithIndex.Value
				continue
			}

			if lineAfter, ok := inputMap[index+1]; ok && hasSymbol(lineAfter[intRange.startIndex:intRange.endIndex]) {
				result += intWithIndex.Value
				continue
			}
		}
	}
	return result
}

type GearsMap map[string][]int

func (gearsMap *GearsMap) set(key string, value int) {
	if (*gearsMap)[key] == nil {
		(*gearsMap)[key] = []int{value}
	} else {
		(*gearsMap)[key] = append((*gearsMap)[key], value)
	}
}

type Ordinates struct {
	x int
	y int
}

func findGearOrdinates(inputMap map[int]string, index int, intRange IntRange) (Ordinates, error) {
	line, ok := inputMap[index]

	if !ok {
		return Ordinates{}, errors.New("gear not found")
	}

	gearIndex := strings.Index(line[intRange.startIndex:intRange.endIndex], "*")

	if gearIndex < 0 {
		return Ordinates{}, errors.New("gear not found")
	}

	return Ordinates{
		x: index,
		y: intRange.startIndex + gearIndex,
	}, nil
}

func solvePart2(input []string) int {
	gearsMap := GearsMap{}
	inputMap := lib.SliceToMap(lib.AocInputLines(2023, 03))
	for lineIndex, line := range inputMap {
		for _, intWithIndex := range lib.ExtractPositiveIntsWithIndex(line) {
			intRange := getIntRange(intWithIndex, line)

			for index := lineIndex - 1; index <= lineIndex+1; index++ {
				gearOrdinates, error := findGearOrdinates(inputMap, index, intRange)
				if error == nil {
					gearsMap.set(fmt.Sprintf("%d-%d", gearOrdinates.x, gearOrdinates.y), intWithIndex.Value)
					break
				}
			}
		}
	}

	var result int
	for _, gears := range gearsMap {
		if len(gears) == 2 {
			result += gears[0] * gears[1]
		}
	}
	return result
}

func main() {
	input := lib.AocInputLines(2023, 03)

	fmt.Println(solvePart1(input))
	fmt.Println(solvePart2(input))
}

package main

import (
	"challenges/lib"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func splitInput(input string) (string, string) {
	splitedInput := strings.Split(input, "\n\n")

	return splitedInput[0], splitedInput[1]
}

type ParsedLetter struct {
	position int
	letter   string
}

func parseStack(stack string) []ParsedLetter {
	splittedStack := strings.Split(stack, "")
	splittedStack = splittedStack[1 : len(splittedStack)-1]

	parsedLetters := []ParsedLetter{}
	for index, letter := range lib.TakeEvery(splittedStack, 4) {
		if letter != " " {
			parsedLetters = append(parsedLetters, ParsedLetter{
				position: index + 1,
				letter:   letter,
			})
		}
	}
	return parsedLetters
}

func parseStacks(stacks string) map[int][]string {
	splittedStacks := strings.Split(stacks, "\n")
	slices.Reverse(splittedStacks)

	parsedStacks := map[int][]string{}
	for _, stack := range lib.DropFirst(splittedStacks) {
		for _, parsedLetter := range parseStack(stack) {
			parsedStacks[parsedLetter.position] = append(parsedStacks[parsedLetter.position], parsedLetter.letter)
		}
	}
	return parsedStacks
}

type ParsedProcedure struct {
	count, source, destination int
}

func parseProcedures(procedures string) []ParsedProcedure {
	procedures = strings.Trim(procedures, "\n")

	parsedProcedures := []ParsedProcedure{}
	for _, procedure := range strings.Split(procedures, "\n") {
		pattern := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := pattern.FindStringSubmatch(procedure)
		numbers, error := lib.StringsToNumbers(matches[1:])

		if error != nil {
			break
		}
		parsedProcedures = append(parsedProcedures, ParsedProcedure{
			count:       numbers[0],
			source:      numbers[1],
			destination: numbers[2],
		})
	}
	return parsedProcedures
}

func move(stacks map[int][]string, procedures []ParsedProcedure, reverse bool) map[int][]string {
	stacks = lib.CopyMap(stacks)
	for _, procedure := range procedures {
		source, exists := stacks[procedure.source]
		if !exists {
			continue
		}

		if len(source) >= procedure.count {
			items := source[len(source)-procedure.count:]
			stacks[procedure.destination] = append(stacks[procedure.destination], lib.If(reverse, lib.CopyAndReverseSlice(items), items)...)
			stacks[procedure.source] = source[:len(source)-procedure.count]
		}
	}
	return stacks
}

func extractLettersFromStacks(parsedStacks map[int][]string) string {
	result := make([]string, len(parsedStacks))
	for position, parsedStack := range parsedStacks {
		letter, _, error := lib.At(parsedStack, -1)

		if error != nil {
			continue
		}
		result[position-1] = letter
	}
	return strings.Join(result, "")
}

func main() {
	input := lib.AocInput(2022, 5)

	stacks, procedures := splitInput(input)
	parsedStacks := parseStacks(stacks)
	parsedProcedures := parseProcedures(procedures)

	parsedStacks1 := move(parsedStacks, parsedProcedures, true)
	parsedStacks2 := move(parsedStacks, parsedProcedures, false)

	fmt.Println(extractLettersFromStacks(parsedStacks1))
	fmt.Println(extractLettersFromStacks(parsedStacks2))
}

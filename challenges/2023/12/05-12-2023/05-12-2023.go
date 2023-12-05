package main

import (
	"challenges/lib"
	"fmt"
	"regexp"
	"strings"
)

func parseSeeds(seeds string) []int {
	match := regexp.MustCompile(`^seeds: (\d+(?: \d+)*)`).FindStringSubmatch(seeds)

	ints, error := lib.StringsToInts(strings.Split(match[1], " "))
	lib.AssertNil(error)

	return ints
}

func main() {
	doubleNewlines := lib.AocInputDoubleNewline(2023, 05)
	seeds := parseSeeds(doubleNewlines[0])

	fmt.Println(seeds)
}

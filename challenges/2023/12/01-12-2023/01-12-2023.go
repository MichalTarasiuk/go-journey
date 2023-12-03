package main

import (
	"challenges/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var sum, sum2 int
	for _, line := range lib.AocInputLines(2023, 01) {
		firstIndex := strings.IndexFunc(line, unicode.IsDigit)
		lastIndex := strings.LastIndexFunc(line, unicode.IsDigit)

		finalInt, err := strconv.Atoi(string(line[firstIndex]) + string(line[lastIndex]))
		lib.AssertNil(err)

		sum += finalInt

		first2 := part2Vals[part2Regexp.FindString(line)]
		last2 := part2Vals[part2Regexp2.FindStringSubmatch(line)[1]]
		finalInt2 := 10*first2 + last2

		sum2 += finalInt2

	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

var part2Regexp = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
var part2Regexp2 = regexp.MustCompile(".*(" + part2Regexp.String() + ")")

var part2Vals = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
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

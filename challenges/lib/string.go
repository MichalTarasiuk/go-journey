package lib

import (
	"regexp"
	"strings"
)

func ExtractLetters(input string) []string {
	return regexp.MustCompile("[A-Za-z]+").FindAllString(input, -1)
}

func HasSameValue(input string) bool {
	return len(input) == 0 || strings.Count(input, string(input[0])) == len(input)
}

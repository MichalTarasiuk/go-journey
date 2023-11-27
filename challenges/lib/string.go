package lib

import "regexp"

func ExtractLetters(input string) []string {
	return regexp.MustCompile("[A-Za-z]+").FindAllString(input, -1)
}

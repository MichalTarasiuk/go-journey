package main

import (
	"challenges/lib"
	"fmt"
	"strings"
)

type shape int

const (
	rock     shape = 1
	paper    shape = 2
	scissors shape = 3
)

var beats = map[shape]shape{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}
var loses = lib.InvertMap(beats)

func main() {
	meChars := map[string]shape{
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	themChars := map[string]shape{
		"A": rock,
		"B": paper,
		"C": scissors,
	}

	var total int

	for _, line := range lib.AocInputLines(2022, 2) {
		players := strings.Split(line, " ")
		me := players[1]
		them := players[0]

		total += play(meChars[me], themChars[them])
	}

	fmt.Println(total)
}

func play(meChar, themChar shape) int {
	if beats[meChar] == themChar {
		return 6 + int(meChar)
	} else if loses[meChar] == themChar {
		return 0 + int(meChar)
	} else if meChar == themChar {
		return 3 + int(meChar)
	}

	panic("not reached")
}

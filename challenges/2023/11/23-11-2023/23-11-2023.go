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

type outcome int

const (
	lose outcome = iota
	draw
	win
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
	outChars := map[string]outcome{
		"X": lose,
		"Y": draw,
		"Z": win,
	}

	var total, total2 int

	for _, line := range lib.AocInputLines(2022, 2) {
		lineArr := strings.Split(line, " ")

		me := lineArr[1]

		them := lineArr[0]
		themChar := themChars[them]

		total += play(meChars[me], themChar)
		total2 += play(choose(themChar, outChars[me]), themChar)
	}

	fmt.Println(total)
	fmt.Println(total2)
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

func choose(them shape, out outcome) shape {
	if out == lose {
		return beats[them]
	} else if out == draw {
		return them
	} else if out == win {
		return loses[them]
	}

	panic("not reached")
}

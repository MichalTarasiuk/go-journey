package main

import (
	"challenges/lib"
	"fmt"
	"strings"
)

func main() {
	total := map[string]int{"red": 12, "green": 13, "blue": 14}

	var result int
	for _, line := range lib.AocInputLines(2023, 2) {
		var id int
		var sets string
		lib.Extract(line, `^Game (\d+): (.+)`, &id, &sets)

		possible := true
		for _, set := range strings.Split(sets, "; ") {
			for _, cubes := range strings.Split(set, ", ") {
				var cnt int
				var color string
				lib.Extract(cubes, `^(\d+) (red|green|blue)$`, &cnt, &color)

				if cnt > total[color] {
					possible = false
				}
			}
		}
		if possible {
			result += id
		}
	}
	fmt.Println(result)
}

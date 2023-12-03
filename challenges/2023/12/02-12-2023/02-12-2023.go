package main

import (
	"challenges/lib"
	"fmt"
	"strings"
)

func main() {
	total := map[string]int{"red": 12, "green": 13, "blue": 14}

	var result1 int
	var result2 int
	for _, line := range lib.AocInputLines(2023, 2) {
		possible := true

		var id int
		var sets string
		lib.Extract(line, `^Game (\d+): (.+)`, &id, &sets)

		counts := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, set := range strings.Split(sets, "; ") {
			for _, cubes := range strings.Split(set, ", ") {
				var cnt int
				var color string
				lib.Extract(cubes, `^(\d+) (red|green|blue)$`, &cnt, &color)

				if cnt > total[color] {
					possible = false
				}
				counts[color] = lib.Max(counts[color], cnt)
			}
		}

		if possible {
			result1 += id
		}
		power := counts["red"] * counts["green"] * counts["blue"]
		result2 += power
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

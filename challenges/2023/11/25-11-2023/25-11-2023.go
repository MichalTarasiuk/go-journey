package main

import (
	"challenges/lib"
	"fmt"
	"strings"
)

func main() {
	r := 0
	for _, chunk := range lib.ChunkEvery(lib.AocInputLines(2022, 3), 3) {
		maps := lib.MapOverSlice(chunk, func(item string) map[string]string {
			return lib.SliceToMap(strings.Split(item, ""))
		})
		badge := lib.MapSomeKey(lib.MapIntersect(maps...))

		r += priority([]rune(badge)[0])
	}
	fmt.Println(r)
}

func priority(item rune) int {
	switch {
	case item >= 'a' && item <= 'z':
		return int(item-'a') + 1
	case item >= 'A' && item <= 'Z':
		return int(item-'A') + 27
	default:
		lib.Panicf("Invalid item %q", item)
		return 0
	}
}

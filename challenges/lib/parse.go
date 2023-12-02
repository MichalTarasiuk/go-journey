package lib

import (
	"regexp"
	"strconv"
)

var intRegexp = regexp.MustCompile(`-?\d+`)

func ExtractInts(s string) []int {
	v64s := ExtractInt64s(s)
	vals := make([]int, len(v64s))
	for index, v64 := range v64s {
		vals[index] = int(v64)
	}
	return vals
}

func ExtractInt64s(s string) []int64 {
	var vals []int64
	for _, s := range intRegexp.FindAllString(s, -1) {
		v, err := strconv.ParseInt(s, 10, 64)
		Assertf(err == nil, "Failed parsing %q as int64: %v", s, err)
		vals = append(vals, v)
	}
	Assertf(len(vals) > 0, "No ints found")
	return vals
}

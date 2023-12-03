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

func StringsToNumbers(s []string) ([]int, error) {
	var r []int
	for _, str := range s {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		r = append(r, n)
	}
	return r, nil
}

type IntWithIndex struct {
	Value int
	Index int
}

func ExtractIntsWithIndex(s string) []IntWithIndex {
	int64sWithIndex := ExtractInt64sWithIndex(s)
	vals := make([]IntWithIndex, len(int64sWithIndex))

	for i, v64 := range int64sWithIndex {
		vals[i] = IntWithIndex{Value: int(v64.Value), Index: v64.Index}
	}

	return vals
}

type Int64WithIndex struct {
	Value int64
	Index int
}

func ExtractInt64sWithIndex(s string) []Int64WithIndex {
	var vals []Int64WithIndex
	matches := intRegexp.FindAllStringIndex(s, -1)

	for _, match := range matches {
		strNum := s[match[0]:match[1]]
		value, err := strconv.ParseInt(strNum, 10, 64)
		Assertf(err == nil, "Failed parsing %q as int64: %v", strNum, err)

		vals = append(vals, Int64WithIndex{Value: value, Index: match[0]})
	}

	return vals
}

var regexpCache = make(map[string]*regexp.Regexp)

func getRegexp(re string) *regexp.Regexp {
	if !MapHasKey(regexpCache, re) {
		regexpCache[re] = regexp.MustCompile(re)
	}
	comp, ok := regexpCache[re]
	Assert(ok)
	return comp
}

func Extract(s, re string, dsts ...interface{}) {
	ms := getRegexp(re).FindStringSubmatch(s)

	for i, dst := range dsts {
		m := ms[i+1]
		if m == "" { // skip optional groups (should maybe set dst to zero?)
			continue
		}

		var err error
		switch t := dst.(type) {
		case *int:
			*t, err = strconv.Atoi(m)
		case *string:
			*t = m
		default:
			Assertf(false, "Unknown dest type %T for group %v of %q", t, i, re)
		}
		Assertf(err == nil, "Failed to parse group %q matched by %q: %v", m, re, err)
	}
}

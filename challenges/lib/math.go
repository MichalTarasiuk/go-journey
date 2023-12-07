package lib

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](vals ...T) T {
	Assertf(len(vals) > 0, "No values given")
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func Max[T constraints.Ordered](vals ...T) T {
	Assertf(len(vals) > 0, "No values given")
	max := vals[0]
	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func Sum[Num constraints.Integer | constraints.Float](numbers ...Num) Num {
	var sum Num
	for _, num := range numbers {
		sum += num
	}
	return sum
}

type NumberRange struct {
	Start int
	End   int
}

func (r1 NumberRange) FindCommonRange(r2 NumberRange) NumberRange {
	start := max(r1.Start, r2.Start)
	end := min(r1.End, r2.End)

	if start <= end {
		return NumberRange{Start: start, End: end}
	}
	return NumberRange{Start: -1, End: -1}
}

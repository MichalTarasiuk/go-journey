package lib

import "golang.org/x/exp/constraints"

func Sum[Num constraints.Integer | constraints.Float](numbers ...Num) Num {
	var sum Num
	for _, num := range numbers {
		sum += num
	}
	return sum
}

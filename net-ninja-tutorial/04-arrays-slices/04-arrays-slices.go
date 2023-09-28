package main

import "fmt"

func main() {

	// arrays
	var ages [3]int = [3]int{20, 18, 39}

	var numbers = [4]int{1, 2, 3, 4}
	numbers[1] = 5


	fmt.Println(ages, len(ages))
	fmt.Println(numbers, len(numbers))

	// slices
	var scores = []int{100, 50, 20}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := numbers[1:3]
	rangeTwo := numbers[2:]
	rangeThree := numbers[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

} 
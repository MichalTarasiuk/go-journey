package main

import "fmt"

func squareSlice(input []int) []int {
	output := []int{}

	for _, value := range input {
		output = append(output, value*value)
	}

	return output
}

func findMax(input []int) int {
	max := input[0]

	for _, value := range input {
		if value > max {
			max = value
		}
	}

	return max
}

func main() {
	// squareSlice
	input := []int{1, 2, 3, 4, 5}

	// output: [1 4 9 16 25]
	fmt.Println(squareSlice(input))

	// findMax
	input2 := []int{1, 2, 3, 4, 5}

	// output: 5
	fmt.Println(findMax(input2))
}

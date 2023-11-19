package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solvePart01() (int, int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	var (
		totalCalories              int
		elfCarryingTheMostCalories int
		currentElfCalories         int
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		calorieValue, err := strconv.Atoi(line)

		if err != nil {
			if currentElfCalories > elfCarryingTheMostCalories {
				elfCarryingTheMostCalories = currentElfCalories
			}

			totalCalories += currentElfCalories
			currentElfCalories = 0
		} else {
			currentElfCalories += calorieValue
		}
	}

	return totalCalories, elfCarryingTheMostCalories, err
}

func main() {
	totalCalories, elfCarryingTheMostCalories, err := solvePart01()
	if err != nil {
		return
	}

	fmt.Printf("total calories: %v, elf carrying the most calories: %v", totalCalories, elfCarryingTheMostCalories)
}

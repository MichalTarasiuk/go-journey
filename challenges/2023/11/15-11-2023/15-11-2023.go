package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func mapToQueryString(map1 map[string]string) string {
	queryStringSlice := []string{}

	for key, value := range map1 {
		queryStringSlice = append(queryStringSlice, fmt.Sprintf("%s=%s", key, value))
	}

	return fmt.Sprintf("?%s", strings.Join(queryStringSlice, "&"))
}

func mapToQueryStringExample() {
	fmt.Println(mapToQueryString(map[string]string{
		"a": "1",
		"b": "2",
	}))
}

func flatMap[Item any](items []Item, fn func(Item) []Item) []Item {
	result := []Item{}

	for _, item := range items {
		result = append(result, fn(item)...)
	}

	return result
}

func flatMapExample() {
	// Example usage:
	numbers := []int{1, 2, 3, 4, 5}

	// Define a function that doubles each element and returns a slice.
	doubleFunc := func(x int) []int {
		return []int{x * 2, x * 2}
	}

	// Apply the flatMap function.
	result := flatMap(numbers, doubleFunc)

	// Print the result.
	fmt.Println(result)

	// Another example with strings.
	words := []string{"hello", "world"}

	// Define a function that appends "!" to each string and returns a slice.
	appendExclamation := func(s string) []string {
		return []string{s + "!"}
	}

	// Apply the flatMap function.
	resultStrings := flatMap(words, appendExclamation)

	// Print the result for strings.
	fmt.Println(resultStrings)
}

type WordFrequency struct {
	Word  string
	Count int
}

func readFileAndCountWords(filePath string) ([]WordFrequency, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		for _, word := range words {
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var frequencies []WordFrequency
	for word, count := range wordCount {
		frequencies = append(frequencies, WordFrequency{Word: word, Count: count})
	}

	return frequencies, nil
}

func readFileAndCountWordsExample() {
	filePath := "input.txt"

	wordFrequencies, err := readFileAndCountWords(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Sort word frequencies by count in descending order.
	sort.Slice(wordFrequencies, func(i, j int) bool {
		return wordFrequencies[i].Count > wordFrequencies[j].Count
	})

	// Print word frequencies.
	fmt.Println("Word Frequencies:")
	for _, wf := range wordFrequencies {
		fmt.Printf("%s: %d\n", wf.Word, wf.Count)
	}
}

func main() {
	mapToQueryStringExample()

	flatMapExample()

	readFileAndCountWordsExample()
}

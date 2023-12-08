package main

import (
	"challenges/lib"
	"fmt"
	"strings"
)

type nodes map[string][]string

var instructionToIndex = map[string]int{
	"L": 0,
	"R": 1,
}

func countSteps(nodeName string, nodes nodes, instructions []string) int {
	if len(instructions) == 0 || strings.Compare("ZZZ", nodeName) == 0 {
		return 0
	}

	node := nodes[nodeName]
	index := instructionToIndex[instructions[0]]

	return 1 + countSteps(node[index], nodes, append(instructions[1:], instructions[0]))
}

func main() {
	input := strings.TrimSpace(lib.AocInput(2023, 8))
	inputSlice := strings.Split(input, "\n\n")

	var nodes = make(nodes)
	for _, node := range strings.Split(inputSlice[1], "\n") {
		var name string
		var left string
		var right string
		lib.Extract(node, `(\w{3}) = \((\w{3}), (\w{3})\)`, &name, &left, &right)

		nodes[name] = []string{left, right}
	}

	fmt.Println(countSteps("AAA", nodes, strings.Split(inputSlice[0], "")))
}

package main

import (
	"challenges/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseSeeds(seeds string) []int {
	match := regexp.MustCompile(`^seeds: (\d+(?: \d+)*)`).FindStringSubmatch(seeds)

	ints, error := lib.StringsToInts(strings.Split(match[1], " "))
	lib.AssertNil(error)

	return ints
}

func parseMapperName(mapName string) (string, string) {
	var source string
	var destination string
	lib.Extract(mapName, `(\w+)-to-(\w+) map:`, &source, &destination)

	return source, destination
}

func parseMapperValues(mapValues []string) [][]int {
	var parsedMapValues [][]int
	for _, mapValue := range mapValues {
		var parsedMapValue []int
		for _, value := range strings.Split(mapValue, " ") {
			if parsedInt64, error := strconv.ParseInt(value, 10, 64); error == nil {
				parsedMapValue = append(parsedMapValue, int(parsedInt64))
			}
		}
		parsedMapValues = append(parsedMapValues, parsedMapValue)
	}
	return parsedMapValues
}

type ParsedMapper struct {
	destination string
	values      [][]int
}

func parseMappers(mappers []string) map[string]ParsedMapper {
	parsedMappers := map[string]ParsedMapper{}
	for _, mapper := range mappers {
		splittedMapper := strings.Split(mapper, "\n")
		source, destionation := parseMapperName(splittedMapper[0])

		parsedMappers[source] = ParsedMapper{
			destination: destionation,
			values:      parseMapperValues(splittedMapper[1:]),
		}
	}
	return parsedMappers
}

func findLocation(mappers map[string]ParsedMapper, source string, value int) int {
	if parsedMapper, ok := mappers[source]; ok {
		for _, recipe := range parsedMapper.values {
			if !(value >= recipe[1] && value <= recipe[1]+(recipe[2]-1)) {
				continue
			}
			return findLocation(mappers, parsedMapper.destination, value-recipe[1]+recipe[0])
		}
		return findLocation(mappers, parsedMapper.destination, value)
	}
	return value
}

func main() {
	doubleNewlines := lib.AocInputDoubleNewline(2023, 05)

	if len(doubleNewlines) == 0 {
		panic("No double new lines found")
	}

	seeds := parseSeeds(doubleNewlines[0])
	mappers := parseMappers(doubleNewlines[1:])

	var locations []int
	for _, seed := range seeds {
		locations = append(locations, findLocation(mappers, "seed", seed))
	}

	result := locations[0]
	for _, location := range locations[1:] {
		if result > location {
			result = location
		}
	}
	fmt.Println(result)
}

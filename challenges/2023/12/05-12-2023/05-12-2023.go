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

type MapperName struct {
	source      string
	destination string
}

func parseMapperName(mapName string) MapperName {
	var source string
	var destination string
	lib.Extract(mapName, `(\w+)-to-(\w+) map:`, &source, &destination)

	return MapperName{
		source:      source,
		destination: destination,
	}
}

func parseMapperValues(mapValues []string) [][]int {
	var parsedMapValues [][]int
	for _, mapValue := range mapValues {
		var parsedMapValue []int
		for _, rune := range mapValue {
			if parsedInt64, error := strconv.ParseInt(string(rune), 10, 64); error == nil {
				parsedMapValue = append(parsedMapValue, int(parsedInt64))
			}
		}
		parsedMapValues = append(parsedMapValues, parsedMapValue)
	}
	return parsedMapValues
}

type ParsedMapper struct {
	name   MapperName
	values [][]int
}

func parseMapper(mapper string) ParsedMapper {
	splittedMapper := strings.Split(mapper, "\n")

	return ParsedMapper{
		name:   parseMapperName(splittedMapper[0]),
		values: parseMapperValues(splittedMapper[1:]),
	}
}

func parseMappers(mappers []string) []ParsedMapper {
	var parsedMappers []ParsedMapper
	for _, mapper := range mappers {
		parsedMappers = append(parsedMappers, parseMapper(mapper))
	}
	return parsedMappers
}

func main() {
	doubleNewlines := lib.AocInputDoubleNewline(2023, 05)
	seeds := parseSeeds(doubleNewlines[0])
	mappers := parseMappers(doubleNewlines[1:])

	fmt.Println(seeds)
	fmt.Println((mappers))
}

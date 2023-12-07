package main

import (
	"challenges/lib"
	"fmt"
	"io/ioutil"
	"log"
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

func parseMapName(mapName string) (string, string) {
	var source string
	var destination string
	lib.Extract(mapName, `(\w+)-to-(\w+) map:`, &source, &destination)

	return source, destination
}

func parseMapValues(mapValues []string) [][]int {
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

type ParsedMap struct {
	destination string
	values      [][]int
}

func parseMaps(maps []string) map[string]ParsedMap {
	parsedMaps := map[string]ParsedMap{}
	for _, aMap := range maps {
		splittedMap := strings.Split(aMap, "\n")
		source, destination := parseMapName(splittedMap[0])

		parsedMaps[source] = ParsedMap{
			destination: destination,
			values:      parseMapValues(splittedMap[1:]),
		}
	}
	return parsedMaps
}

func findLocation(maps map[string]ParsedMap, source string, value int) int {
	if parsedMap, ok := maps[source]; ok {
		for _, recipe := range parsedMap.values {
			if !(value >= recipe[1] && value <= recipe[1]+(recipe[2]-1)) {
				continue
			}
			return findLocation(maps, parsedMap.destination, value-recipe[1]+recipe[0])
		}
		return findLocation(maps, parsedMap.destination, value)
	}
	return value
}

func main() {
	body, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	input := strings.TrimSpace(string(body))
	doubleNewlines := strings.Split(input, "\n\n")

	if len(doubleNewlines) == 0 {
		panic("No double new lines found")
	}

	seeds := parseSeeds(doubleNewlines[0])
	maps := parseMaps(doubleNewlines[1:])

	var locations1 []int
	for _, seed := range seeds {
		locations1 = append(locations1, findLocation(maps, "seed", seed))
	}
	fmt.Println(lib.Min(locations1...))

	var locations2 []int
	for _, seedChunk := range lib.ChunkEvery(seeds, 2) {
		lib.Assert(len(seedChunk) == 2)
		seedRange := lib.NumberRange{
			Start: seedChunk[0],
			End:   seedChunk[0] + seedChunk[1],
		}

		seedToSoil := maps["seed"]
		var soils []int
		for _, soil := range seedToSoil.values {
			soilRange := lib.NumberRange{
				Start: soil[1],
				End:   soil[1] + soil[2] - 1,
			}
			commonRange := seedRange.FindCommonRange(soilRange)

			if commonRange.Start == -1 && commonRange.End == -1 {
				continue
			}
			for value := commonRange.Start; value <= commonRange.End; value++ {
				soils = append(soils, findLocation(maps, seedToSoil.destination, value))
			}
		}

		for _, soil := range soils {
			locations2 = append(locations2, findLocation(maps, "fertilizer", soil))
		}
	}
	fmt.Println(lib.Min(locations2...))
}

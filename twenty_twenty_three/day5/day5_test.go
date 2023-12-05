package day3_test

import (
	"math"
	"strconv"
	"strings"
	"sync"
	"testing"
	"unicode"

	"github.com/mckornfield/aoc/twenty_twenty_three/resources"
	"gotest.tools/assert"
)

type Set map[string]bool

type MappingFn func(int) int

type MappingWithRange struct {
	sourceToTargetFunctions []MappingFn
	source, target          string
}

func dangerousAtoi(str string) int {
	if str == "" {
		return 0
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func parseSeedsPartOne(line string) []SeedRange {
	seeds := []SeedRange{}
	seedList := strings.Split(line, ":")[1]
	for _, seed := range strings.Split(seedList, " ") {
		if seed != "" {
			seeds = append(seeds, SeedRange{start: dangerousAtoi(seed), length: 1})
		}
	}
	return seeds
}

func parseSeedsPartTwo(line string) []SeedRange {
	seeds := []SeedRange{}
	seedList := strings.Split(line, ":")[1]
	seedRangeStart := 0
	for _, seed := range strings.Split(seedList, " ") {
		// fmt.Println(seed)
		if seed != "" && seedRangeStart == 0 {
			seedRangeStart = dangerousAtoi(seed)
		} else if seed != "" {
			seedRangeLength := dangerousAtoi(seed)
			seeds = append(seeds, SeedRange{start: seedRangeStart, length: seedRangeLength})
			seedRangeStart = 0
		}
		// fmt.Println(seeds)
	}
	return seeds
}

type SeedParser func(line string) []SeedRange

type SeedRange struct {
	start, length int
}

func getLowestSeedMapping(lines []string, seedParser SeedParser) (int, error) {
	seeds := []SeedRange{}
	mappings := map[string]MappingWithRange{}
	currentMapping := MappingWithRange{sourceToTargetFunctions: []MappingFn{}}
	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			seeds = seedParser(line)
		}
		if strings.HasSuffix(line, "map:") {
			if currentMapping.source != "" {
				mappings[currentMapping.source] = currentMapping
				currentMapping = MappingWithRange{sourceToTargetFunctions: []MappingFn{}}
			}
			splitLine := strings.Split(line, " ")
			splitLine = strings.Split(splitLine[0], "-to-")
			currentMapping.source, currentMapping.target = splitLine[0], splitLine[1]
		}
		if line != "" && unicode.IsDigit(rune(line[0])) {
			splitLine := strings.Split(line, " ")
			destStart, sourceStart, length := dangerousAtoi(splitLine[0]), dangerousAtoi(splitLine[1]), dangerousAtoi(splitLine[2])

			mappingFunc := func(i int) int {
				if i >= sourceStart && i < sourceStart+length {
					return i - sourceStart + destStart
				}
				return i
			}
			currentMapping.sourceToTargetFunctions = append(currentMapping.sourceToTargetFunctions, mappingFunc)
		}
	}
	if currentMapping.source != "" {
		mappings[currentMapping.source] = currentMapping
	}
	// fmt.Println(mappings)
	minValue := math.MaxInt
	// fmt.Println(len(seeds))
	var wg sync.WaitGroup
	minValueFound := make(chan int)

	for _, seedRange := range seeds {
		// fmt.Println(seedRange)
		wg.Add(1)
		go func(seedRange SeedRange) {
			for i := seedRange.start; i < seedRange.start+seedRange.length; i++ {
				currentValue := mapValue(i, "seed", mappings)
				if currentValue < minValue {
					minValue = currentValue
				}
			}
			minValueFound <- minValue
			wg.Done()
		}(seedRange)
	}
	// for _, seedValue := range seeds {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		currentValue = mapValue(i, "seed", mappings)
	// 		wg.Done()
	// 	}(seedValue)
	// }
	go func() {
		wg.Wait()
		close(minValueFound)
	}()
	for elem := range minValueFound {
		if elem < minValue {
			minValue = elem
		}
	}
	return minValue, nil
}

func mapValue(currentValue int, currentSource string, mappings map[string]MappingWithRange) int {
	for currentSource != "" {
		// fmt.Print(currentSource, " ", currentValue, " ")
		currentMapping := mappings[currentSource]
		for _, mappingFunc := range currentMapping.sourceToTargetFunctions {
			newValue := mappingFunc(currentValue)
			if newValue != currentValue {
				currentValue = newValue
				break
			}
		}
		currentSource = currentMapping.target
	}
	return currentValue
}

func TestRunSamplePt1(t *testing.T) {
	lines := resources.ReadLines("pt1_example.txt")
	result, err := getLowestSeedMapping(lines, parseSeedsPartOne)
	assert.NilError(t, err)
	assert.Equal(t, result, 35)
}

func TestRunPt1(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result, err := getLowestSeedMapping(lines, parseSeedsPartOne)
	assert.NilError(t, err)
	assert.Equal(t, result, 265018614)
}

func TestRunPt2Example(t *testing.T) {
	lines := resources.ReadLines("pt2_example.txt")
	result, err := getLowestSeedMapping(lines, parseSeedsPartTwo)
	assert.NilError(t, err)
	assert.Equal(t, result, 46)
}

func TestRunPt2(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result, err := getLowestSeedMapping(lines, parseSeedsPartTwo)
	assert.NilError(t, err)
	assert.Equal(t, result, 63179500)
}

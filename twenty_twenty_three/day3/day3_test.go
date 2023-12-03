package day3_test

import (
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/mckornfield/aoc/twenty_twenty_three/resources"
	"gotest.tools/assert"
)

type Pair struct {
	x, y int
}

type NumberLocation struct {
	locations []Pair
	value     int
}

type SymbolAdjacentLocations map[Pair]bool

func processLines(lines []string) (SymbolAdjacentLocations, []NumberLocation, error) {
	symbolAdjacentLocations := make(map[Pair]bool)
	numberLocations := []NumberLocation{}
	// Rows -> Y, from top
	// Columns -> X from left
	for y, line := range lines {
		var b strings.Builder
		locations := []Pair{}
		for x, c := range line {
			isNumber := unicode.IsDigit(c)
			if isNumber {
				b.WriteRune(c)
				locations = append(locations, Pair{x: x, y: y})
			}
			if b.Len() > 0 && (!isNumber || x == len(line)-1) {
				value, err := strconv.Atoi(b.String())
				if err != nil {
					return nil, nil, err
				}
				location := NumberLocation{
					locations: locations,
					value:     value,
				}
				numberLocations = append(numberLocations, location)
				locations = []Pair{}
				b.Reset()
			}

			if c != '.' && !isNumber {
				for dy := -1; dy < 2; dy++ {
					for dx := -1; dx < 2; dx++ {
						symbolAdjacentLocations[Pair{x: x + dx, y: y + dy}] = true
					}
				}
			}

		}
	}
	// fmt.Println(symbolAdjacentLocations)
	// fmt.Println(numberLocations)
	return symbolAdjacentLocations, numberLocations, nil
}

func hasAdjacentSymbol(numberLocation NumberLocation, symbolAdjacentLocations SymbolAdjacentLocations) bool {
	// checkedLocations := make(map[Pair]bool)
	for _, location := range numberLocation.locations {
		if symbolAdjacentLocations[location] {
			return true
		}
	}

	return false
}

func sumOfPartsWithAdjacentSymbols(lines []string) (int, error) {
	symbolLocations, numberLocations, err := processLines(lines)
	if err != nil {
		return -1, err
	}
	// Check if there's any adjacent values for each number
	sum := 0
	for _, numberLocation := range numberLocations {
		// fmt.Println(numberLocation.value)
		if hasAdjacentSymbol(numberLocation, symbolLocations) {
			sum += numberLocation.value
		}
	}
	return sum, nil
}

type UniqueNumber struct {
	value         int
	startingIndex Pair
}

type SymbolLocations map[Pair]bool

type NumberLocationsUnique map[Pair]UniqueNumber

func processLinesPt2(lines []string) (SymbolLocations, NumberLocationsUnique, error) {
	symbolLocations := SymbolLocations{}
	numberLocations := NumberLocationsUnique{}
	// Rows -> Y, from top
	// Columns -> X from left
	for y, line := range lines {
		var b strings.Builder
		locations := []Pair{}
		for x, c := range line {
			isNumber := unicode.IsDigit(c)
			if isNumber {
				b.WriteRune(c)
				locations = append(locations, Pair{x: x, y: y})
			}
			if b.Len() > 0 && (!isNumber || x == len(line)-1) {
				value, err := strconv.Atoi(b.String())
				if err != nil {
					return nil, nil, err
				}
				uniqueNumber := UniqueNumber{
					value:         value,
					startingIndex: locations[0],
				}
				for _, location := range locations {
					numberLocations[location] = uniqueNumber
				}
				locations = []Pair{}
				b.Reset()
			}

			if c != '.' && !isNumber {
				symbolLocations[Pair{x: x, y: y}] = true
			}

		}
	}
	// fmt.Println(symbolAdjacentLocations)
	// fmt.Println(numberLocations)
	return symbolLocations, numberLocations, nil
}

func getAdjacencyBox(pair Pair) []Pair {
	pairs := []Pair{}
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			pairs = append(pairs, Pair{x: pair.x + dx, y: pair.y + dy})
		}
	}
	return pairs
}

func sumOfPartsWithGears(lines []string) (int, error) {
	symbolLocations, numberLocationsUnique, err := processLinesPt2(lines)
	if err != nil {
		return -1, err
	}
	// Check if there's any adjacent values for each number
	sum := 0
	for symbolLocation := range symbolLocations {
		matchingNumbers := make(map[UniqueNumber]bool)
		for _, pair := range getAdjacencyBox(symbolLocation) {
			if uniqueNumber, ok := numberLocationsUnique[pair]; ok {
				matchingNumbers[uniqueNumber] = true
			}
		}
		// if len(matchingNumbers) == 1 {
		// 	fmt.Println("one match", matchingNumbers)
		// 	for uniqueNumber := range matchingNumbers {
		// 		sum += uniqueNumber.value
		// 	}
		if len(matchingNumbers) == 2 {
			// fmt.Println("two match", matchingNumbers)
			gearSum := 1
			for uniqueNumber := range matchingNumbers {
				gearSum *= uniqueNumber.value
			}
			sum += gearSum
		}

	}
	return sum, nil
}

func TestRunSamplePt1(t *testing.T) {
	lines := resources.ReadLines("pt1_example.txt")
	result, err := sumOfPartsWithAdjacentSymbols(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 4361)
}

func TestRunPt1(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result, err := sumOfPartsWithAdjacentSymbols(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 557705)
}

func TestRunPt2Example(t *testing.T) {
	lines := resources.ReadLines("pt2_example.txt")
	result, err := sumOfPartsWithGears(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 467835)
}

func TestRunPt2(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result, err := sumOfPartsWithGears(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 84266818)
}

package aoc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mckornfield/aoc/twenty_twenty_three/resources"
	"gotest.tools/assert"
)

type Set map[string]bool

func parseLine(line string) (int, error) {
	if line == "" {
		return 0, nil
	}
	gameAndNumbers := strings.Split(line, ":")
	winningAndActualNumbers := strings.Split(gameAndNumbers[1], "|")
	winningNumbersStr, actualNumbersStr := winningAndActualNumbers[0], winningAndActualNumbers[1]
	winningNumbers := Set{}
	for _, numberStr := range strings.Split(winningNumbersStr, " ") {
		if numberStr == "" {
			continue
		}
		winningNumbers[numberStr] = true
	}
	sum := 0
	for _, numberStr := range strings.Split(actualNumbersStr, " ") {
		if numberStr == "" {
			continue
		}
		if winningNumbers[numberStr] {
			if sum == 0 {
				sum++
			} else {
				sum *= 2
			}
		}
	}
	return sum, nil
}

func scratchCardValues(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		value, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		sum += value
	}
	return sum, nil
}

func TestRunSamplePt1(t *testing.T) {
	lines := resources.ReadLines("pt1_example.txt")
	result, err := scratchCardValues(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 13)
}

func TestRunPt1(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result, err := scratchCardValues(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 23441)
}

func processDuplicatingLine(line string) (int, error) {
	if line == "" {
		return 0, nil
	}
	gameAndNumbers := strings.Split(line, ":")
	winningAndActualNumbers := strings.Split(gameAndNumbers[1], "|")
	winningNumbersStr, actualNumbersStr := winningAndActualNumbers[0], winningAndActualNumbers[1]
	winningNumbers := Set{}
	for _, numberStr := range strings.Split(winningNumbersStr, " ") {
		if numberStr == "" {
			continue
		}
		winningNumbers[numberStr] = true
	}
	sum := 0
	for _, numberStr := range strings.Split(actualNumbersStr, " ") {
		if numberStr == "" {
			continue
		}
		if winningNumbers[numberStr] {
			sum++
		}
	}
	return sum, nil
}

type TimesToProcessCard map[int]int

func processDuplicatingCards(lines []string) (int, error) {
	timesToProcess := TimesToProcessCard{}
	for i, line := range lines {
		if line == "" {
			continue
		}
		timesToProcess[i+1] = 1
	}
	for gameNum, line := range lines {
		value, err := processDuplicatingLine(line)
		if err != nil {
			return -1, err
		}
		for i := 1; i < value+1; i++ {
			timesToProcess[gameNum+i+1] += 1 * timesToProcess[gameNum+1]
		}
	}
	fmt.Println(timesToProcess)
	// Count the cards at the end
	sum := 0
	for _, count := range timesToProcess {
		sum += count
	}
	return sum, nil
}

func TestRunPt2Example(t *testing.T) {
	lines := resources.ReadLines("pt2_example.txt")
	result, err := processDuplicatingCards(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 30)
}

func TestRunPt2(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result, err := processDuplicatingCards(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 5923918)
}

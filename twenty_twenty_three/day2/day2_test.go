package day1_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mckornfield/aoc/twenty_twenty_three/resources"
	"gotest.tools/assert"
)

type Game map[string]int

var maxGame = Game{"red": 12, "blue": 14, "green": 13}

func NewGame(gameString string) (Game, error) {
	game := Game{}
	for _, countAndColor := range strings.Split(gameString, ", ") {
		countAndColorSplit := strings.Split(countAndColor, " ")
		count, color := countAndColorSplit[0], countAndColorSplit[1]
		countInt, err := strconv.Atoi(count)
		if err != nil {
			return game, err
		}
		game[color] = countInt
	}
	return game, nil
}

func isGamePossible(lineSegment string) (bool, error) {
	game, err := NewGame(lineSegment)
	if err != nil {
		return false, err
	}
	for color, count := range game {
		if maxGame[color] < count {
			return false, nil
		}
	}
	return true, nil
}

func getInvalidGameSum(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, ": ")
		gameSection, gamesString := splitLine[0], splitLine[1]
		gameSectionSplit := strings.Split(gameSection, " ")
		gameNumber, err := strconv.Atoi(gameSectionSplit[1])
		if err != nil {
			return -1, err
		}

		for _, gameString := range strings.Split(gamesString, "; ") {
			isPossible, err := isGamePossible(gameString)
			if err != nil {
				return -1, err
			}
			if !isPossible {
				gameNumber = 0
			}
		}
		sum += gameNumber
	}
	return sum, nil
}

func applyMaxes(lineSegment string, gameMaxes Game) error {
	game, err := NewGame(lineSegment)
	if err != nil {
		return err
	}
	for color, count := range game {
		if gameMaxes[color] < count {
			gameMaxes[color] = count
		}
	}
	return nil
}

func getNecessaryMaxSum(lines []string) (int, error) {
	sumOfGamePower := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, ": ")
		gamesString := splitLine[1]
		gameMaxes := Game{"red": 0, "green": 0, "blue": 0}
		for _, gameString := range strings.Split(gamesString, "; ") {
			err := applyMaxes(gameString, gameMaxes)
			if err != nil {
				return -1, err
			}
		}
		sumOfGamePower += gameMaxes["red"] * gameMaxes["green"] * gameMaxes["blue"]
	}
	return sumOfGamePower, nil
}

func TestRunSamplePt1(t *testing.T) {
	lines := resources.ReadLines("pt1_example.txt")
	result, _ := getInvalidGameSum(lines)
	assert.Equal(t, result, 8)
}

func TestRunPt1(t *testing.T) {
	lines := resources.ReadLines("day2.txt")
	result, _ := getInvalidGameSum(lines)
	assert.Equal(t, result, 2061)
}

func TestRunPt2Example(t *testing.T) {
	lines := strings.Split(resources.ReadFile("pt2_example.txt"), "\n")
	sum, err := getNecessaryMaxSum(lines)
	assert.NilError(t, err)
	assert.Equal(t, sum, 2286)
}

func TestRunPt2(t *testing.T) {
	lines := strings.Split(resources.ReadFile("day2.txt"), "\n")
	sum, err := getNecessaryMaxSum(lines)
	assert.NilError(t, err)
	assert.Equal(t, sum, 72596)
}

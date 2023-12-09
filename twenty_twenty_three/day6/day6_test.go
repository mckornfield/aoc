package aoc_test

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/mckornfield/aoc/twenty_twenty_three/resources"
	"gotest.tools/assert"
)

func dangerousAtoi(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}

type Race struct {
	distance, time int
}

func parseRaces(lines []string) []Race {
	re := regexp.MustCompile(`\s+`)
	raceTimes := []int{}
	raceDistances := []int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, ":")
		lineType, numbers := lineSplit[0], lineSplit[1]
		if lineType == "Time" {
			for _, number := range re.Split(numbers, -1) {
				if number == "" {
					continue
				}
				raceTimes = append(raceTimes, dangerousAtoi(number))
			}
		} else if lineType == "Distance" {
			for _, number := range re.Split(numbers, -1) {
				if number == "" {
					continue
				}
				raceDistances = append(raceDistances, dangerousAtoi(number))
			}
		}
	}
	races := []Race{}
	for i := 0; i < len(raceTimes); i++ {
		races = append(races, Race{
			distance: raceDistances[i],
			time:     raceTimes[i],
		})
	}
	return races
}

func (r Race) calculateWaysToWin() int {
	waysToWin := 0
	for holdDuration := 0; holdDuration < r.time; holdDuration++ {
		velocity := holdDuration // 1 millimeter per second held
		distanceTraveled := velocity * (r.time - holdDuration)
		if distanceTraveled > r.distance {
			// fmt.Println("Race", r, "Won with hold duration", holdDuration, "distance", distanceTraveled)
			waysToWin++
		}
	}
	return waysToWin
}

func getMultipliedWaysToWinRace(lines []string) (int, error) {
	races := parseRaces(lines)

	waysToWinList := []int{}
	for _, race := range races {
		result := race.calculateWaysToWin()
		waysToWinList = append(waysToWinList, result)
	}
	fmt.Println("ways to win", waysToWinList)
	product := 1
	for _, wayToWin := range waysToWinList {
		product *= wayToWin
	}
	return product, nil
}

func TestRunSamplePt1(t *testing.T) {
	lines := resources.ReadLines("pt1_example.txt")
	result, err := getMultipliedWaysToWinRace(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 288)
}

func TestRunPt1(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result, err := getMultipliedWaysToWinRace(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 625968)
}

func TestRunPt2Example(t *testing.T) {
	lines := resources.ReadLines("pt2_example.txt")
	result, err := getMultipliedWaysToWinRace(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 71503)
}

func TestRunPt2(t *testing.T) {
	lines := resources.ReadLines("input_2.txt")
	result, err := getMultipliedWaysToWinRace(lines)
	assert.NilError(t, err)
	assert.Equal(t, result, 43663323)
}

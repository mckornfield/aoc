package day1_test

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/mckornfield/aoc/twenty_twenty_three/resources"
	"gotest.tools/assert"
)

var alpha = regexp.MustCompile(`[^a-zA-Z ]+`)

type getNumbers func(line string) (int, error)

func getNumbersFromLine(line string) (int, error) {
	firstNumber := ' '
	secondNumber := ' '
	for _, c := range line {
		if unicode.IsDigit(c) {
			if firstNumber == ' ' {
				firstNumber = c
			}
			secondNumber = c
		}
	}
	res, err := strconv.Atoi(string(firstNumber) + string(secondNumber))
	return res, err
}

var numberMapping = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func getNumberFromLineWithWordReplacement(line string) (int, error) {
	firstNumber := ' '
	secondNumber := ' '
	for i, c := range line {
		for word, number := range numberMapping {
			if strings.HasPrefix(line[i:], word) {
				c = rune(number)
			}
		}
		if unicode.IsDigit(c) {
			if firstNumber == ' ' {
				firstNumber = c
			}
			secondNumber = c
		}
	}
	res, err := strconv.Atoi(string(firstNumber) + string(secondNumber))
	return res, err
}

func getSumsFromLines(lines []string, getNumberFunc getNumbers) (int, error) {
	sum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		result, err := getNumberFunc(line)
		if err != nil {
			return 0, err
		}
		fmt.Println(line + " " + fmt.Sprint(result))
		sum += result
	}
	return sum, nil
}

func TestRunSamplePt1(t *testing.T) {
	result := resources.ReadFile("day1_pt1_example.txt")
	assert.Equal(t, result, `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`)
	lines := strings.Split(result, "\n")
	cleanedLine, err := getNumbersFromLine(lines[0])
	assert.NilError(t, err)
	assert.Equal(t, cleanedLine, 12)

	sum, err := getSumsFromLines(lines, getNumbersFromLine)
	assert.NilError(t, err)
	assert.Equal(t, sum, 142)
}

func TestRunPt1(t *testing.T) {
	lines := strings.Split(resources.ReadFile("day1.txt"), "\n")
	sum, err := getSumsFromLines(lines, getNumbersFromLine)
	assert.NilError(t, err)
	assert.Equal(t, sum, 55002)
}

func TestRunPt2Example(t *testing.T) {
	lines := strings.Split(resources.ReadFile("day1_pt2_example.txt"), "\n")
	sum, err := getSumsFromLines(lines, getNumberFromLineWithWordReplacement)
	assert.NilError(t, err)
	assert.Equal(t, sum, 281)
}

func TestRunPt2(t *testing.T) {
	lines := strings.Split(resources.ReadFile("day1.txt"), "\n")
	sum, err := getSumsFromLines(lines, getNumberFromLineWithWordReplacement)
	assert.NilError(t, err)
	assert.Equal(t, sum, 55093)
}

package aoc_test

import (
	"fmt"
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

func calculateDifferences(sequence []int) []int {
	differences := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		differences[i] = sequence[i+1] - sequence[i]
	}
	return differences
}

func allZeroes(sequence []int) bool {
	for _, num := range sequence {
		if num != 0 {
			return false
		}
	}
	return true
}

func nextNumberInSequence(sequence []int) int {
	differences := calculateDifferences(sequence)
	// Whole sequence is the same
	if allZeroes(differences) {
		return sequence[len(sequence)-1]
	}
	return sequence[len(sequence)-1] + nextNumberInSequence(differences)
}

func sequenceStrToSequenceReversed(sequenceStr string) []int {
	sequenceSplit := strings.Split(sequenceStr, " ")
	sequence := []int{}
	for i := len(sequenceSplit) - 1; i >= 0; i-- {
		sequence = append(sequence, dangerousAtoi(sequenceSplit[i]))
	}
	return sequence
}

func sequenceStrToSequence(sequenceStr string) []int {
	sequenceSplit := strings.Split(sequenceStr, " ")
	sequence := []int{}
	for _, numStr := range sequenceSplit {
		sequence = append(sequence, dangerousAtoi(numStr))
	}
	return sequence
}

func getSumOfLastNumbersInSequences(lines []string) int {
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		result += nextNumberInSequence(sequenceStrToSequence(line))
	}
	return result
}

func getSumOfPreviousNumbersInSequences(lines []string) int {
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		reversedSeq := sequenceStrToSequenceReversed(line)
		res := nextNumberInSequence(reversedSeq)
		fmt.Println(reversedSeq, res)
		result += res
	}
	return result
}

func TestGetNextNumberInSequence(t *testing.T) {
	for _, test := range []struct {
		sequence string
		expected int
	}{
		{sequence: "0 3 6 9 12 15", expected: 18},
		{sequence: "1 3 6 10 15 21", expected: 28},
		{sequence: "10 13 16 21 30 45", expected: 68},
	} {
		result := nextNumberInSequence(sequenceStrToSequence(test.sequence))
		assert.Equal(t, result, test.expected)
	}
}

func TestGetNextNumberInSequenceReversed(t *testing.T) {
	for _, test := range []struct {
		sequence string
		expected int
	}{
		{sequence: "0 3 6 9 12 15", expected: -3},
		{sequence: "1 3 6 10 15 21", expected: 0},
		{sequence: "10 13 16 21 30 45", expected: 5},
	} {
		result := nextNumberInSequence(sequenceStrToSequenceReversed(test.sequence))
		assert.Equal(t, result, test.expected)
	}
}

func TestRunSamplePt1(t *testing.T) {
	lines := resources.ReadLines("pt1_example.txt")
	result := getSumOfLastNumbersInSequences(lines)
	assert.Equal(t, result, 114)
}

func TestRunPt1(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result := getSumOfLastNumbersInSequences(lines)
	assert.Equal(t, result, 1955513104)
}

func TestRunSamplePt2(t *testing.T) {
	lines := resources.ReadLines("pt1_example.txt")
	result := getSumOfPreviousNumbersInSequences(lines)
	assert.Equal(t, result, 2)
}

func TestRunPt2(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result := getSumOfPreviousNumbersInSequences(lines)
	assert.Equal(t, result, 1131)
}

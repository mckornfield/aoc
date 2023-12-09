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

type LeftRightMap map[string]LeftRightSegment

type LeftRightSegment struct {
	left, right string
}

func parseLine(line string) (string, LeftRightSegment) {
	lineSplit := strings.Split(line, " = ")
	key := lineSplit[0]
	leftRightOptions := strings.Trim(lineSplit[1], "()")
	leftRightSplit := strings.Split(leftRightOptions, ", ")
	segment := LeftRightSegment{
		left:  leftRightSplit[0],
		right: leftRightSplit[1],
	}
	return key, segment
}

func getStepsToZZZ(lines []string) int {
	leftRightMap := LeftRightMap{}
	instructions := ""
	for _, line := range lines {
		if strings.Contains(line, " = ") {
			key, segment := parseLine(line)
			leftRightMap[key] = segment
		} else if line != "" {
			instructions = line
		}
	}
	currentNode := "AAA"
	steps := 0
	for currentNode != "ZZZ" {
		segment := leftRightMap[currentNode]
		direction := instructions[steps%len(instructions)]
		if direction == 'L' {
			currentNode = segment.left
		} else {
			currentNode = segment.right
		}
		steps++
	}
	return steps
}

func allNodesAreAtTheEnd(currentNodes []string) bool {
	for _, node := range currentNodes {
		if !strings.HasSuffix(node, "Z") {
			return false
		}
	}
	return true
}

func doesMapContainNode(leftRightMap map[string]int, currentNode string) bool {
	_, ok := leftRightMap[currentNode]
	return ok
}

func findLengthOfCycle(currentNode, instructions string, leftRightMap LeftRightMap) (uint64, uint64) {
	steps := 0
	visitedNodes := map[string]int{}
	visitedNodesList := []string{}
	visitedNodes[currentNode] = steps
	visitedNodesList = append(visitedNodesList, currentNode)
	for {
		segment := leftRightMap[currentNode]
		direction := instructions[steps%len(instructions)]
		if direction == 'L' {
			currentNode = segment.left
		} else {
			currentNode = segment.right
		}
		steps++
		if found, ok := visitedNodes[currentNode]; ok {
			for i := found; i < steps; i++ { // The Z node has to be in the cycle
				if strings.HasSuffix(visitedNodesList[i], "Z") {
					return uint64(found), uint64(+steps - visitedNodes[currentNode]) // cycle length
				}
			}
		}
		visitedNodes[currentNode] = steps
		visitedNodesList = append(visitedNodesList, currentNode)
	}
}

func getSimultaneousStepsToZZZ(lines []string) uint64 {
	leftRightMap := LeftRightMap{}
	instructions := ""
	currentNodes := []string{}
	for _, line := range lines {
		if strings.Contains(line, " = ") {
			key, segment := parseLine(line)
			leftRightMap[key] = segment
			if strings.HasSuffix(key, "A") {
				currentNodes = append(currentNodes, key)
			}
		} else if line != "" {
			instructions = line
		}
	}
	var lcf uint64
	lcf = 1
	starts := []uint64{}
	lengths := []uint64{}
	for _, currentNode := range currentNodes {
		start, length := findLengthOfCycle(currentNode, instructions, leftRightMap)
		fmt.Println("Node", currentNode, "cycle length", length, "cycle start", start)
		starts = append(starts, start)
		lengths = append(lengths, length)
	}

	maxStart := uint64(0)
	for _, start := range starts {
		if start > maxStart {
			maxStart = start
		}
	}
	// Calculate cycle offsets
	offsets := []uint64{}
	for i := 0; i < len(starts); i++ {
		cycleOffset := (maxStart - starts[i]) % lengths[i]
		if cycleOffset == 0 {
			offsets = append(offsets, lengths[i])
		} else {
			offsets = append(offsets, cycleOffset)
		}
	}
	fmt.Println(lengths)
	fmt.Println(maxStart)
	fmt.Println(offsets)
	for i := 0; i < len(lengths); i++ {
		lcf *= offsets[i]
	}

	return lcf
}

// func TestRunSamplePt1(t *testing.T) {
// 	lines := resources.ReadLines("pt1_example.txt")
// 	result := getStepsToZZZ(lines)
// 	assert.Equal(t, result, 2)
// }

// func TestRunSample2Pt1(t *testing.T) {
// 	lines := resources.ReadLines("pt1_example_2.txt")
// 	result := getStepsToZZZ(lines)
// 	assert.Equal(t, result, 6)
// }

// func TestRunPt1(t *testing.T) {
// 	lines := resources.ReadLines("input.txt")
// 	result := getStepsToZZZ(lines)
// 	assert.Equal(t, result, 23147)
// }

// func TestRunSamplePt2(t *testing.T) {
// 	lines := resources.ReadLines("pt2_example.txt")
// 	result := getSimultaneousStepsToZZZ(lines)
// 	assert.Equal(t, result, 6)
// }

func TestRunPt2(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result := getSimultaneousStepsToZZZ(lines)
	assert.Equal(t, result, 6)
}

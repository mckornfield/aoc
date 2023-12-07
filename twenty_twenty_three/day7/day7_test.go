package day3_test

import (
	"sort"
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

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type HandAndBid struct {
	hand     string
	handType HandType
	bid      int
}

var cardValuesNoJoker = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	// 'J': 1,
}

var cardValuesJoker = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	// 'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func isHandOneBigger(hand1, hand2 HandAndBid, cardValues map[rune]int) bool {
	if hand1.handType != hand2.handType {
		return hand1.handType > hand2.handType
	}
	hand2Runes := []rune(hand2.hand)
	for i, hand1C := range hand1.hand {
		hand1Value := cardValues[hand1C]
		hand2Value := cardValues[hand2Runes[i]]
		if hand1Value != hand2Value {
			return hand1Value > hand2Value
		}
	}
	return true // BaseCase?
}

type GetHandTypeInterface func(hand string) HandType

func getHandType(hand string) HandType {
	matches := map[rune]int{}
	for _, card := range hand {
		matches[card]++
	}
	switch len(matches) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		maxCount := 0
		for _, count := range matches {
			if count > maxCount {
				maxCount = count
			}
		}
		if maxCount == 3 {
			return ThreeOfAKind
		} else if maxCount == 2 {
			return TwoPair
		}
	case 2:
		maxCount := 0
		for _, count := range matches {
			if count > maxCount {
				maxCount = count
			}
		}

		if maxCount == 4 {
			return FourOfAKind
		} else if maxCount == 3 {
			return FullHouse
		}
	case 1:
		return FiveOfAKind
	}

	return HighCard
}

func getHandTypeJoker(hand string) HandType {
	matches := map[rune]int{}
	for _, card := range hand {
		matches[card]++
	}
	handType := getHandType(hand)
	if matches['J'] == 1 {
		if handType == HighCard {
			return OnePair
		} else if handType == OnePair {
			return ThreeOfAKind
		} else if handType == TwoPair {
			return FullHouse
		} else if handType == ThreeOfAKind {
			return FourOfAKind
		} else if handType == FourOfAKind {
			return FiveOfAKind
		}
	} else if matches['J'] == 2 {
		if handType == OnePair {
			return ThreeOfAKind // Add one in
		} else if handType == TwoPair {
			return FourOfAKind
		} else if handType == FullHouse {
			return FiveOfAKind
		}
	} else if matches['J'] == 3 {
		if handType == FullHouse {
			return FiveOfAKind
		} else if handType == ThreeOfAKind {
			return FourOfAKind
		}
	} else if matches['J'] == 4 {
		return FiveOfAKind
	}

	return handType
}

func getWinnings(lines []string, getHandTypeI GetHandTypeInterface, cardValues map[rune]int) (int, error) {
	handsAndBids := []HandAndBid{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, " ")
		hand, bidStr := splitLine[0], splitLine[1]
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			return -1, err
		}
		handsAndBids = append(handsAndBids, HandAndBid{
			hand:     hand,
			handType: getHandTypeI(hand),
			bid:      bid,
		})
	}
	sort.Slice(handsAndBids, func(i, j int) bool {
		return isHandOneBigger(handsAndBids[i], handsAndBids[j], cardValues)
	})
	result := 0
	for i := len(handsAndBids) - 1; i >= 0; i-- {
		rank := len(handsAndBids) - i

		result += handsAndBids[i].bid * rank
	}
	return result, nil
}

func TestGetHandType(t *testing.T) {
	assert.Equal(t, getHandType("72777"), FourOfAKind)
	assert.Equal(t, getHandType("77777"), FiveOfAKind)
	assert.Equal(t, getHandType("6QQ66"), FullHouse)
	assert.Equal(t, getHandType("37377"), FullHouse)
	assert.Equal(t, getHandType("A8624"), HighCard)
	assert.Equal(t, getHandType("A8666"), ThreeOfAKind)
	assert.Equal(t, getHandType("A8666"), ThreeOfAKind)
	assert.Equal(t, getHandType("27276"), TwoPair)
	assert.Equal(t, getHandType("22954"), OnePair)
	assert.Equal(t, getHandType("TK4JT"), OnePair)
}

func TestGetHandTypeJoker(t *testing.T) {
	tests := []struct {
		input    string
		expected HandType
	}{
		{"72777", FourOfAKind},
		{"JJ44A", FourOfAKind},
		{"77777", FiveOfAKind},
		{"6QQ66", FullHouse},
		{"37377", FullHouse},
		{"A8624", HighCard},
		{"A8666", ThreeOfAKind},
		{"A8666", ThreeOfAKind},
		{"27276", TwoPair},
		{"22954", OnePair},
		{"TK4JT", ThreeOfAKind},
		{"TKJJT", FourOfAKind},
		{"TKJJT", FourOfAKind},
		{"JJJJT", FiveOfAKind},
		{"TTAAJ", FullHouse},
		{"TTAAJ", FullHouse},
		{"T55J5", FourOfAKind},
		{"KTJJT", FourOfAKind},
		{"QQQJA", FourOfAKind},
		{"QQQQJ", FiveOfAKind},
		{"JJA45", ThreeOfAKind},
		{"JJA45", ThreeOfAKind},
		{"JJJ55", FiveOfAKind},
		{"JJ555", FiveOfAKind},
		{"J5555", FiveOfAKind},
		{"J5544", FullHouse},
		{"J5444", FourOfAKind},
		{"J2345", OnePair},
		{"J4455", FullHouse},
		{"JJ443", FourOfAKind},
		{"JJ444", FiveOfAKind},
		{"JJ234", ThreeOfAKind},
		{"JJJ34", FourOfAKind},
		{"JJJ44", FiveOfAKind},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, getHandTypeJoker(test.input))
	}
}

func TestRunSamplePt1(t *testing.T) {
	lines := resources.ReadLines("pt1_example.txt")
	result, err := getWinnings(lines, getHandType, cardValuesNoJoker)
	assert.NilError(t, err)
	assert.Equal(t, result, 6440)
}

func TestRunPt1(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result, err := getWinnings(lines, getHandType, cardValuesNoJoker)
	assert.NilError(t, err)
	assert.Equal(t, result, 250058342)
}

func TestRunSamplePt2(t *testing.T) {
	lines := resources.ReadLines("pt1_example.txt")
	result, err := getWinnings(lines, getHandTypeJoker, cardValuesJoker)
	assert.NilError(t, err)
	assert.Equal(t, result, 5905)
}

func TestRunPt2(t *testing.T) {
	lines := resources.ReadLines("input.txt")
	result, err := getWinnings(lines, getHandTypeJoker, cardValuesJoker)
	assert.NilError(t, err)
	assert.Equal(t, result, 251502303)
}

package part2

import (
	"log"
	"silverark/aoc-2023/pkg/shared"
	"strconv"
	"strings"
	"unicode"
)

const (
	maxInt = int(^uint(0) >> 1)
)

var numberWords = map[string]string{
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

func process(items []string) int {

	total := 0

	// Loop through the rows and get the first and last numberWords
	for _, item := range items {
		first := FirstNumber(item)
		last := LastNumber(item)

		coOrdinates, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalf("error converting number: %v", err)
		}
		total += coOrdinates
	}

	return total
}

func FirstNumber(s string) string {
	return findSets(s, false)
}

func LastNumber(s string) string {
	return findSets(shared.Reverse(s), true)
}

func findSets(s string, flip bool) string {
	// Set index as max uint so that we can just see which index is the lowest when comparing it with words
	numericalIndex := maxInt
	for i, r := range s {
		if unicode.IsDigit(r) {
			numericalIndex = i
			break
		}
	}
	// Loop through every word and find the integer of it. If it's lower than the current integer then save
	wordIndex := maxInt
	wordValue := ""
	for i, word := range numberWords {
		if flip {
			word = shared.Reverse(word)
		}
		foundIndex := strings.Index(s, word)
		if foundIndex != -1 && foundIndex < wordIndex {
			wordIndex = foundIndex
			wordValue = i
		}
	}

	if numericalIndex == -1 && wordIndex == -1 {
		return ""
	}

	// is numericalIndex or wordIndex lower?
	if numericalIndex < wordIndex {
		return s[numericalIndex : numericalIndex+1]
	} else {
		return wordValue
	}
}

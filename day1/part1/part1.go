package part1

import (
	"log"
	"silverark/aoc-2023/pkg/shared"
	"strconv"
	"unicode"
)

func process(items []string) int {

	total := 0

	// Loop through the rows and get the first and last numbers
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
	for i, r := range s {
		if unicode.IsDigit(r) {
			return s[i : i+1]
		}
	}
	return ""
}

func LastNumber(s string) string {
	return FirstNumber(shared.Reverse(s))
}

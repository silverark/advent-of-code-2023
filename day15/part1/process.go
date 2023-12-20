package part1

import (
	"math"
	"strings"
)

func process(input []string) int {
	steps := strings.Split(input[0], ",")
	total := 0
	for _, step := range steps {
		total += hashValue(step)
	}
	return total
}

func hashValue(input string) int {
	currVal := 0
	for _, char := range input {
		currVal += int(char)
		currVal *= 17
		remainder := math.Mod(float64(currVal), 256)
		currVal = int(remainder)
	}
	return currVal
}

package part1

import (
	"slices"
	"strings"
)

func process(input []string) int {
	total := 0
	for _, line := range input {
		parts := strings.Split(line, "|")
		winningParts := strings.Split(parts[0], ":")
		ticketNumbers := strings.Fields(parts[1])
		winNumbers := strings.Fields(winningParts[1])

		// Now calculate the score
		gameTotal := 0
		for _, winNumber := range winNumbers {
			if slices.Contains(ticketNumbers, winNumber) {
				// If it's zero set it to one
				if gameTotal == 0 {
					gameTotal = 1
					continue
				}
				// otherwise double the score
				gameTotal *= 2
			}
		}
		total += gameTotal
	}
	return total
}

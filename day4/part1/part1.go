package part1

import (
	"strconv"
	"strings"
)

func process(input []string) int {
	total := 0
	for _, line := range input {
		winNumbers := make(map[int]bool)
		ticketNumbers := make(map[int]bool)

		parts := strings.Split(line, "|")
		winningParts := strings.Split(parts[0], ":")
		ticketNumberStr := strings.Split(parts[1], " ")
		winNumbersStr := strings.Split(winningParts[1], " ")

		for _, ticketNumber := range ticketNumberStr {
			if ticketNumber == "" {
				continue
			}
			ticketNum, _ := strconv.Atoi(ticketNumber)
			winNumbers[ticketNum] = true
		}
		for _, winNumber := range winNumbersStr {
			if winNumber == "" {
				continue
			}
			winNum, _ := strconv.Atoi(winNumber)
			ticketNumbers[winNum] = true
		}

		// Now calculate the score
		gameTotal := 0
		for winNumber, _ := range winNumbers {
			if ticketNumbers[winNumber] {
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

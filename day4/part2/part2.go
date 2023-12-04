package part2

import (
	"slices"
	"strings"
)

type Game struct {
	winners int
	count   int
}

func process(input []string) int {

	var cards []Game

	for _, line := range input {
		parts := strings.Split(line, "|")
		winningParts := strings.Split(parts[0], ":")
		ticketNumbers := strings.Fields(parts[1])
		winNumbers := strings.Fields(winningParts[1])

		winners := 0
		for _, winNumber := range winNumbers {
			if slices.Contains(ticketNumbers, winNumber) {
				winners++
			}
		}

		cards = append(cards, Game{
			count:   1,
			winners: winners,
		})
	}

	// Now calculate the winners and add cards to the list
	extraCards := 0
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j <= (i + cards[i].winners); j++ {
			cards[j].count += cards[i].count // Adding the card count as that's how many times it's won
			extraCards += cards[i].count
		}
	}
	return len(cards) + extraCards
}

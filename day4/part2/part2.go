package part2

import (
	"strconv"
	"strings"
)

func process(input []string) int {

	var cards []Game

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

		cards = append(cards, Game{
			winNumbers:    winNumbers,
			ticketNumbers: ticketNumbers,
			count:         1,
		})
	}

	// Now calculate the winners and add cards to the list
	extraCards := 0
	for i := 0; i < len(cards); i++ {
		winners := cards[i].Winners()
		// We need to do this cards[j].count times
		for cardCount := 0; cardCount < cards[i].count; cardCount++ {
			for j := i + 1; j <= (i + winners); j++ {
				if j > len(cards) {
					continue
				}
				cards[j].count++
				extraCards++
			}
		}
	}
	return len(cards) + extraCards
}

type Game struct {
	winNumbers    map[int]bool
	ticketNumbers map[int]bool
	count         int
}

func (g Game) Winners() int {
	winners := 0
	for winNumber, _ := range g.winNumbers {
		if g.ticketNumbers[winNumber] {
			winners++
		}
	}
	return winners
}

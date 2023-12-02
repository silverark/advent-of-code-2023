package part1

import (
	"strconv"
	"strings"
)

func Process(items []string) int {
	total := 0
	for _, item := range items {
		gameColourCount := map[string]int{"red": 0, "green": 0, "blue": 0}
		parts := strings.Split(item, ":")
		id, _ := strconv.Atoi(strings.ReplaceAll(parts[0], "Game ", ""))
		games := strings.Split(parts[1], ";")
		for _, game := range games {
			gameColours := strings.Split(game, ",")
			for _, gameColour := range gameColours {
				gameColour = strings.TrimSpace(gameColour)
				numColour := strings.Split(gameColour, " ")
				itemCount, _ := strconv.Atoi(strings.TrimSpace(numColour[0]))
				if itemCount > gameColourCount[numColour[1]] {
					gameColourCount[numColour[1]] = itemCount
				}
			}
		}
		if gameColourCount["red"] <= 12 && gameColourCount["green"] <= 13 && gameColourCount["blue"] <= 14 {
			total += id
		}
	}
	return total
}

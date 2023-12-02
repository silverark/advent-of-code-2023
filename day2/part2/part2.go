package part2

import (
	"strconv"
	"strings"
)

type ColourCount struct {
	red   int
	green int
	blue  int
}

func Process(items []string) int {
	total := 0
	for _, item := range items {
		_, colours := ParseLine(item)
		total += colours.red * colours.green * colours.blue
	}
	return total
}

// ParseLine splits the string "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
func ParseLine(s string) (id int, colours ColourCount) {
	parts := strings.Split(s, ":")
	id, _ = strconv.Atoi(strings.ReplaceAll(parts[0], "Game ", ""))
	games := strings.Split(parts[1], ";")
	for _, game := range games {
		gameColours := strings.Split(game, ",")
		for _, gameColour := range gameColours {
			colour := strings.TrimSpace(gameColour)
			if strings.Contains(colour, "red") {
				itemCount, _ := strconv.Atoi(strings.ReplaceAll(colour, " red", ""))
				if itemCount > colours.red {
					colours.red = itemCount
				}
			} else if strings.Contains(colour, "green") {
				itemCount, _ := strconv.Atoi(strings.ReplaceAll(colour, " green", ""))
				if itemCount > colours.green {
					colours.green = itemCount
				}
			} else if strings.Contains(colour, "blue") {
				itemCount, _ := strconv.Atoi(strings.ReplaceAll(colour, " blue", ""))
				if itemCount > colours.blue {
					colours.blue = itemCount
				}
			}
		}
	}

	return
}

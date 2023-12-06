package part2

import (
	"strconv"
	"strings"
)

func process(input []string) int {
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(input[0], ":")[1], " ", ""))
	recordDistance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(input[1], ":")[1], " ", ""))
	winners := 0
	for speed := 1; speed < time; speed++ {
		distance := (time - speed) * speed
		if distance > recordDistance {
			winners++
		}
	}
	return winners
}

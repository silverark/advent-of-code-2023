package part1

import (
	"strconv"
	"strings"
)

func process(input []string) int {

	times := strings.Fields(input[0])
	distances := strings.Fields(input[1])

	winnersMultipled := 1
	for i := 1; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		recordDistance, _ := strconv.Atoi(distances[i])

		winners := 0
		for speed := 1; speed < time; speed++ {
			distance := (time - speed) * speed
			if distance > recordDistance {
				winners++
			}
		}
		winnersMultipled *= winners
	}

	return winnersMultipled
}

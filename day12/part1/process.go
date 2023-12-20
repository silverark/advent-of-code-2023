package part1

import (
	"strconv"
	"strings"
)

type SpringRow struct {
	row    string
	broken []int
}

func process(input []string) int {

	var allSprings []SpringRow

	for _, row := range input {
		parts := strings.Split(row, " ")
		spring := SpringRow{row: parts[0]}
		broken := strings.Split(parts[1], ",")
		for _, numStr := range broken {
			num, _ := strconv.Atoi(numStr)
			spring.broken = append(spring.broken, num)
		}
		allSprings = append(allSprings, spring)
	}

	return 0
}

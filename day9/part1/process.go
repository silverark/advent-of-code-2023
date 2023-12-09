package part1

import (
	"strconv"
	"strings"
)

func process(input []string) int {
	numbers := make([][]int, len(input))
	total := 0
	for i, line := range input {
		vals := strings.Fields(line)
		for _, number := range vals {
			num, _ := strconv.Atoi(number)
			numbers[i] = append(numbers[i], num)
		}
		total += calcNext(numbers[i])
	}
	return total
}

func calcNext(row []int) int {
	allSame := false
	var levels [][]int
	var nextLevel []int
	nextLevel = row
	for allSame == false {
		nextLevel, allSame = nextRow(nextLevel)
		if allSame {
			nextNum := nextLevel[0]
			for i := len(levels) - 1; i >= 0; i-- {
				nextNum = levels[i][len(levels[i])-1] + nextNum
				levels[i] = append(levels[i], nextNum)
			}
			return row[len(row)-1] + nextNum
		}
		levels = append(levels, nextLevel)
	}
	return 0
}

func nextRow(nums []int) (nextRow []int, allSame bool) {
	allSame = true
	for i := 1; i < len(nums); i++ {
		nextValue := nums[i] - nums[i-1]
		nextRow = append(nextRow, nextValue)
		if nextRow[0] != nextValue {
			allSame = false
		}
	}
	return
}

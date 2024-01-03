package shared

import (
	"log"
	"strconv"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Atoi(s string) int {
	intValue, err := strconv.Atoi(s)
	if err != nil {
		log.Panicln("Unable to parse an int: ", err)
	}
	return intValue
}

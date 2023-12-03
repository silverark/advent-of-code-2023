package part1

import (
	"strconv"
	"unicode"
)

func process(input []string) int {
	// go through char by char, we if find a number get the length
	// pass a string with characters around the number to work out if it's touching a char.
	// if it is add it to the map of solutions

	partNumbers := make(map[int]bool)
	total := 0

	for rowIndex, row := range input {
		for charIndex := 0; charIndex < len(row); charIndex++ {
			if unicode.IsDigit(rune(input[rowIndex][charIndex])) {
				startIndex := charIndex
				lastIndex := charIndex

				// Fnd the last number

				for lastIndex < len(row) && unicode.IsDigit(rune(input[rowIndex][lastIndex])) {
					lastIndex++
				}

				// Set loop to end
				charIndex = lastIndex
				// Now see if it's touching a special character by getthing the line above, current line and line below
				// and scanning for special characters
				searchRangeStart := startIndex
				if startIndex > 0 {
					searchRangeStart--
				}
				searchRangeEnd := lastIndex
				if lastIndex < len(row) {
					searchRangeEnd++
				}

				lineAbove := ""
				if rowIndex > 0 {
					// Get a substring of the line above with 1 char either side of the number
					lineAbove = input[rowIndex-1][searchRangeStart:searchRangeEnd]
				}

				currentLine := input[rowIndex][searchRangeStart:searchRangeEnd]

				lineBelow := ""
				if rowIndex < len(input)-1 {
					lineBelow = input[rowIndex+1][searchRangeStart:searchRangeEnd]
				}

				if hasSpecialChar(lineAbove) || hasSpecialChar(currentLine) || hasSpecialChar(lineBelow) {
					partNumber, _ := strconv.Atoi(input[rowIndex][startIndex:lastIndex])
					// We're keeping a map in case the part number is in there twice.
					//if partNumbers[partNumber] == false {
					partNumbers[partNumber] = true
					total += partNumber
					//}
				}

			}
		}
	}

	return total
}

func hasSpecialChar(s string) bool {
	for _, char := range s {
		// Make sure it's not a full stop "." or in the range 0-9
		if char == 46 || (char >= 48 && char <= 57) {
			continue
		}
		return true
	}
	return false
}

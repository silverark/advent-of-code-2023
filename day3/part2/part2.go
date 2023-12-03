package part2

import (
	"fmt"
	"strconv"
	"unicode"
)

var input []string
var starMap map[string][]int

func process(inputData []string) int {
	// go through char by char, we if find a number get the length
	// pass a string with characters around the number to work out if it's touching a char.
	// if it is add it to the map of solutions

	input = inputData
	starMap = make(map[string][]int)
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
				partNumber, _ := strconv.Atoi(input[rowIndex][startIndex:lastIndex])

				addStarMap(rowIndex, startIndex, lastIndex, partNumber)

			}
		}
	}

	// Now we have a map of stars, we need to go through each one and see if it's in more than one part
	for _, parts := range starMap {
		if len(parts) > 1 {
			total += parts[0] * parts[1]
		}
	}

	return total
}

func addStarMap(rowIndex, start, end, partNo int) {

	//get the start and end plus one either side
	searchColStart := start
	if start > 0 {
		searchColStart--
	}
	searchColEnd := end
	if end < len(input[rowIndex]) {
		searchColEnd++
	}

	searchRowStart := rowIndex
	if rowIndex > 0 {
		searchRowStart--
	}
	searchRowEnd := rowIndex
	if rowIndex < len(input)-1 {
		searchRowEnd++
	}

	for row := searchRowStart; row <= searchRowEnd; row++ {
		//log.Printf("Searching row %v\n", row)
		for charIndex, char := range input[row][searchColStart:searchColEnd] {

			// If it's a star (*) then add it to starMap
			if char == 42 {
				//log.Println("Found a star at", row, searchColStart+charIndex)
				starLoc := fmt.Sprintf("%v-%v", row, searchColStart+charIndex)
				if starMap[starLoc] == nil {
					starMap[starLoc] = []int{partNo}
				} else {
					starMap[starLoc] = append(starMap[starLoc], partNo)
				}
			}
		}
	}

}

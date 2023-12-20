package part1

import (
	"log"
)

func process(input []string) int {
	var puzzles [][]string
	// Break into patterns
	puzzles = make([][]string, 1)
	for _, row := range input {
		if row == "" {
			puzzles = append(puzzles, []string{})
			continue
		}
		puzzles[len(puzzles)-1] = append(puzzles[len(puzzles)-1], row)
	}

	total := 0
	for _, puzzle := range puzzles {
		total += findMirror(puzzle)
	}

	// Go through each horizontal and then vertical seeing if it's a mirror
	return total
}

// findMirror will find the mirror in the puzzle and return the `summary` of the pattern
func findMirror(puzzle []string) int {
	// work through on a diagonal line checking both horizontal and vertical
	position := 0
	flippedPuzzle := flipPuzzle(puzzle)
	for {
		if position+1 < len(puzzle) { // bigger than rows
			// Check the horizontal
			for line := position; line > -1; line-- {
				if checkMirror(puzzle, line) {
					return (position + 1) * 100
				}
			}
		}
		if position+1 < len(puzzle[0]) { // bigger than cols?
			// Check vertical (by flipping the puzzle)
			for line := position; line > -1; line-- {
				if checkMirror(flippedPuzzle, line) {
					return position + 1
				}
			}
		}
		if position > len(puzzle[0]) && position > len(puzzle) {
			break
		}
		position++
	}

	log.Panicf("No mirror found for %v", puzzle)
	return 0
}
func checkMirror(puzzle []string, position int) bool {
	// Check horizontally
	for line := position; line > -1; line-- {
		nextLine := position + (1 + (position - line))
		if len(puzzle) == nextLine {
			// We've checked all the rows and they have been the same so far so it's a match!
			return true
		}
		if puzzle[line] != puzzle[nextLine] {
			return false
		}
	}
	return true
}
func flipPuzzle(puzzle []string) []string {
	// Swap rows and columns
	flippedPuzzle := make([]string, len(puzzle[0]))
	for i := 0; i < len(puzzle[0]); i++ {
		for j := 0; j < len(puzzle); j++ {
			flippedPuzzle[i] += string(puzzle[j][i])
		}
	}
	return flippedPuzzle
}

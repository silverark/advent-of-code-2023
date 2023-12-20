package part1

import (
	"fmt"
)

type move struct {
	direction string // r,l,u,d
	x         int
	y         int
}

var turnMoves []move
var energized map[string]map[string]bool
var mirrorMap []string

func process(input []string) int {

	mirrorMap = input
	largestNumber := 0
	cols := len(input[0])
	rows := len(mirrorMap)

	//log.Println("Rows", rows, "Cols", cols)
	// All left
	for row := 0; row < len(mirrorMap); row++ {
		//log.Print("LEFT Testing X Y ", 0, row)
		energized = make(map[string]map[string]bool)
		turnMoves = []move{}
		// Plot first move
		turnMoves = append(turnMoves, move{direction: "r", x: 0, y: row})
		calculate()
		numberEnergised := len(energized)
		//log.Print(" received", numberEnergised, "\n")
		if numberEnergised > largestNumber {
			largestNumber = numberEnergised
		}
	}

	// all right
	for row := 0; row < len(mirrorMap); row++ {
		//log.Println("RIGHT Testing X Y ", cols-1, row)
		energized = make(map[string]map[string]bool)
		turnMoves = []move{}
		// Plot first move
		turnMoves = append(turnMoves, move{direction: "l", x: cols - 1, y: row})
		calculate()
		numberEnergised := len(energized)
		//log.Print(" received", numberEnergised, "\n")
		if numberEnergised > largestNumber {
			largestNumber = numberEnergised
		}
	}

	// All top
	for col := 0; col < cols; col++ {
		//log.Println("TOP Testing X Y ", col, 0)
		energized = make(map[string]map[string]bool)
		turnMoves = []move{}
		// Plot first move
		turnMoves = append(turnMoves, move{direction: "d", x: col, y: 0})
		calculate()
		numberEnergised := len(energized)
		//log.Print(" received", numberEnergised, "\n")
		if numberEnergised > largestNumber {
			largestNumber = numberEnergised
		}
	}

	//all bottom
	for col := 0; col < cols; col++ {
		//log.Println("BOTTOM Testing X Y ", col, rows-1)
		energized = make(map[string]map[string]bool)
		turnMoves = []move{}
		// Plot first move
		turnMoves = append(turnMoves, move{direction: "u", x: col, y: rows - 1})
		calculate()
		numberEnergised := len(energized)
		//log.Print(" received", numberEnergised, "\n")
		if numberEnergised > largestNumber {
			largestNumber = numberEnergised
		}
	}

	return largestNumber

}

func calculate() {
	counter := 0
	for len(turnMoves) > 0 {
		// log.Println("Turn", counter, len(turnMoves))
		// Copy the turn moves
		currentRound := make([]move, len(turnMoves))
		copy(currentRound, turnMoves)
		turnMoves = []move{}
		for _, turnMove := range currentRound {
			// First add it to the energized map
			if energized[fmt.Sprintf("%v,%v", turnMove.x, turnMove.y)] == nil {
				energized[fmt.Sprintf("%v,%v", turnMove.x, turnMove.y)] = make(map[string]bool)
			}
			if energized[fmt.Sprintf("%v,%v", turnMove.x, turnMove.y)][turnMove.direction] {
				// Already energized with direction. Kill the loop
				continue
			}
			energized[fmt.Sprintf("%v,%v", turnMove.x, turnMove.y)][turnMove.direction] = true

			// Now process the move
			tile := mirrorMap[turnMove.y][turnMove.x]
			switch tile {
			case '.': // Add to next tile based on direction
				switch turnMove.direction {
				case "r":
					AddMove(turnMove, "r")
				case "l":
					AddMove(turnMove, "l")
				case "u":
					AddMove(turnMove, "u")
				case "d":
					AddMove(turnMove, "d")
				}
			case '|': // Add to next tile based on direction
				switch turnMove.direction {
				case "r", "l": // Add one new up and one new down
					AddMove(turnMove, "u")
					AddMove(turnMove, "d")
				case "u": // Add one new left and one new right
					AddMove(turnMove, "u")
				case "d":
					AddMove(turnMove, "d")
				}
			case '-': // Add to next tile based on direction
				switch turnMove.direction {
				case "u", "d": // Add one new left and one new right
					AddMove(turnMove, "l")
					AddMove(turnMove, "r")
				case "r":
					AddMove(turnMove, "r")
				case "l":
					AddMove(turnMove, "l")
				}
			case '/':
				switch turnMove.direction {
				case "r":
					AddMove(turnMove, "u")
				case "l":
					AddMove(turnMove, "d")
				case "u":
					AddMove(turnMove, "r")
				case "d":
					AddMove(turnMove, "l")
				}
			case '\\':
				switch turnMove.direction {
				case "r":
					AddMove(turnMove, "d")

				case "l":
					AddMove(turnMove, "u")
				case "u":
					AddMove(turnMove, "l")
				case "d":
					AddMove(turnMove, "r")
				}
			}
		}
		counter++
	}
}

func AddMove(currentMove move, direction string) {
	switch direction {
	case "r":
		if len(mirrorMap[0]) <= currentMove.x+1 {
			return
		}
		turnMoves = append(turnMoves, move{direction: "r", x: currentMove.x + 1, y: currentMove.y})
	case "l":
		if currentMove.x-1 < 0 {
			return
		}
		turnMoves = append(turnMoves, move{direction: "l", x: currentMove.x - 1, y: currentMove.y})
	case "u":
		if currentMove.y-1 < 0 {
			return
		}
		turnMoves = append(turnMoves, move{direction: "u", x: currentMove.x, y: currentMove.y - 1})
	case "d":
		if len(mirrorMap) <= currentMove.y+1 {
			return
		}
		turnMoves = append(turnMoves, move{direction: "d", x: currentMove.x, y: currentMove.y + 1})
	}
}

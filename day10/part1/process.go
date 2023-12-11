package part1

import (
	"fmt"
	"log"
	"strings"
)

const (
	validNorth = "|7F"
	validSouth = "|LJ"
	validEast  = "-J7"
	validWest  = "-LF"
)

type pos struct {
	x int
	y int
}

type grid struct {
	rows     []string
	path     []string
	startPos *pos
	lastMove string // N,S,E,W
}

func (g *grid) followPath() {

	// find start XY
	for i := 0; i < len(g.rows); i++ {
		for j := 0; j < len(g.rows[i]); j++ {
			if g.rows[i][j] == 'S' {
				g.startPos = &pos{x: i, y: j}
				break
			}
		}
		if g.startPos != nil {
			break
		}
	}
	// Find a next move
	x, y := g.nextMove(g.startPos.x, g.startPos.y)

	log.Printf("Starting at %v,%v", x, y)
	// Now keep on going until we land on S again.
	for {
		x, y = g.nextMove(x, y)
		if x == 0 && y == 0 {
			log.Fatalf("INVALID MOVE")
		}
		log.Println(x, y, string(g.rows[x][y]), g.lastMove)
		if g.rows[x][y] == 'S' {
			break
		}
	}

}

func (g *grid) nextMove(x, y int) (int, int) {

	current := g.rows[x][y]

	if current == 'S' {
		// look north if we're not at the top line
		if x != 0 {
			n := g.rows[x-1][y]
			if strings.Contains(validNorth, string(n)) {
				g.path = append(g.path, fmt.Sprint("%,%", x-1, y))
				g.lastMove = "N"
				return x - 1, y
			}
		}
		//South
		if x != len(g.rows)-1 {
			n := g.rows[x+1][y]
			if strings.Contains(validSouth, string(n)) {
				g.path = append(g.path, fmt.Sprint("%,%", x+1, y))
				g.lastMove = "S"
				return x + 1, y
			}
		}
		//East
		if y != len(g.rows[x])-1 {
			n := g.rows[x][y+1]
			if strings.Contains(validEast, string(n)) {
				g.path = append(g.path, fmt.Sprint("%,%", x, y+1))
				g.lastMove = "E"
				return x, y + 1
			}
		}
		//west
		if y != 0 {
			n := g.rows[x][y-1]
			if strings.Contains(validWest, string(n)) {
				g.path = append(g.path, fmt.Sprint("%,%", x, y-1))
				g.lastMove = "W"
				return x, y - 1
			}
		}
	}

	if current == '|' {
		if g.lastMove == "N" {
			g.path = append(g.path, fmt.Sprint("%,%", x-1, y))
			g.lastMove = "N"
			return x - 1, y
		}
		if g.lastMove == "S" {
			g.path = append(g.path, fmt.Sprint("%,%", x+1, y))
			g.lastMove = "S"
			return x + 1, y
		}
	}
	if current == '-' {
		if g.lastMove == "E" {
			g.path = append(g.path, fmt.Sprint("%,%", x, y+1))
			g.lastMove = "E"
			return x, y + 1
		}
		if g.lastMove == "W" {
			g.path = append(g.path, fmt.Sprint("%,%", x, y-1))
			g.lastMove = "W"
			return x, y - 1
		}
	}

	if current == '7' {
		if g.lastMove == "E" {
			g.path = append(g.path, fmt.Sprint("%,%", x+1, y))
			g.lastMove = "S"
			return x + 1, y
		}
		if g.lastMove == "N" {
			g.path = append(g.path, fmt.Sprint("%,%", x, y-1))
			g.lastMove = "W"
			return x, y - 1
		}
	}

	if current == 'F' {
		if g.lastMove == "W" {
			g.path = append(g.path, fmt.Sprint("%,%", x+1, y))
			g.lastMove = "S"
			return x + 1, y
		}
		if g.lastMove == "N" {
			g.path = append(g.path, fmt.Sprint("%,%", x, y+1))
			g.lastMove = "E"
			return x, y + 1
		}
	}

	if current == 'L' {
		if g.lastMove == "S" {
			g.path = append(g.path, fmt.Sprint("%,%", x, y+1))
			g.lastMove = "E"
			return x, y + 1
		}
		if g.lastMove == "W" {
			g.path = append(g.path, fmt.Sprint("%,%", x-1, y))
			g.lastMove = "N"
			return x - 1, y
		}
	}

	if current == 'J' {
		if g.lastMove == "E" {
			g.path = append(g.path, fmt.Sprint("%,%", x-1, y))
			g.lastMove = "N"
			return x - 1, y
		}
		if g.lastMove == "S" {
			g.path = append(g.path, fmt.Sprint("%,%", x, y-1))
			g.lastMove = "W"
			return x, y - 1
		}
	}

	return 0, 0
}

func process(input []string) int {

	g := grid{rows: input}

	// Find the start
	g.followPath()

	// follow both paths

	return len(g.path) / 2
}

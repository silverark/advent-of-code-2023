package part2

import (
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	dir   string
	count int
}
type coord struct {
	X int
	Y int
}

func process(input []string) int {

	var instructions []Instruction
	for _, line := range input {
		parts := strings.Split(line, "#")
		colour := parts[1][0:6]
		count, err := strconv.ParseInt(colour[:5], 16, 64)
		if err != nil {
			panic(err)
		}
		//log.Println("Count", count)
		direction := "U"
		if colour[5] == '0' {
			direction = "R"
		}
		if colour[5] == '1' {
			direction = "D"
		}
		if colour[5] == '2' {
			direction = "L"
		}
		instructions = append(instructions, Instruction{
			dir:   direction,
			count: int(count),
		})
	}

	x := 0
	y := 0
	perimeter := 0
	var grid []coord
	for _, instruction := range instructions {
		grid = append(grid, coord{x, y})

		switch instruction.dir {
		case "R":
			x = x + instruction.count
		case "L":
			x = x - instruction.count
		case "U":
			y = y - instruction.count
		case "D":
			y = y + instruction.count
		}
		perimeter += instruction.count
		grid = append(grid, coord{x, y})

	}

	area := getArea(grid)
	log.Println("Area is", area)
	log.Println("Perimeter is", perimeter)

	result := area + perimeter/2 + 1
	return result
}

func getArea(grid []coord) int {
	//Let 'vertices' be an array of N pairs (X,Y), indexed from 0
	//Let 'area' = 0.0
	//for i = 0 to N-1, do
	//  Let j = (i+1) mod N
	//  Let area = area + vertices[i].X * vertices[j].Y
	//  Let area = area - vertices[i].Y * vertices[j].X
	//end for
	//Return 'area'
	//
	//https://web.archive.org/web/20100405070507/http://valis.cs.uiuc.edu/~sariel/research/CG/compgeom/msg00831.html

	area := 0.0
	for i := 0; i < len(grid); i++ {
		j := (i + 1) % len(grid)
		area = area + float64(grid[i].X*grid[j].Y)
		area = area - float64(grid[i].Y*grid[j].X)
	}

	return int(area / 2)

}

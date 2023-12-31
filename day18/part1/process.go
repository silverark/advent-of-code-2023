package part1

import (
	"log"
	"silverark/aoc-2023/pkg/shared"
	"strings"
)

type Instruction struct {
	dir    string
	count  int
	colour string
}
type coord struct {
	X int
	Y int
}

func process(input []string) int {

	var instructions []Instruction
	for _, line := range input {
		parts := strings.Fields(line)
		instructions = append(instructions, Instruction{
			dir:    parts[0],
			count:  shared.Atoi(parts[1]),
			colour: strings.ReplaceAll(strings.ReplaceAll(parts[2], "(", ""), ")", ""),
		})
	}

	x := 0
	y := 0
	perimeter := 0
	var grid []coord
	for _, instruction := range instructions {

		//log.Println(instruction)
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

	area := 0.0
	for i := 0; i < len(grid); i++ {
		j := (i + 1) % len(grid)
		area = area + float64(grid[i].X*grid[j].Y)
		area = area - float64(grid[i].Y*grid[j].X)
	}

	return int(area / 2)

}

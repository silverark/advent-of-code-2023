package part1

import (
	"fmt"
	"silverark/aoc-2023/pkg/shared"
	"slices"
	"strings"
)

type Brick struct {
	x, y, z    int
	x2, y2, z2 int
}

func (b *Brick) String() string {
	return fmt.Sprintf("%d,%d,%d~%d,%d,%d", b.x, b.y, b.z, b.x2, b.y2, b.z2)
}

func (b *Brick) WouldCollide(otherBrick *Brick) bool {
	if b.z <= otherBrick.z2 && b.z2 >= otherBrick.z {
		if b.y <= otherBrick.y2 && b.y2 >= otherBrick.y {
			if b.x <= otherBrick.x2 && b.x2 >= otherBrick.x {
				return true
			}
		}
	}
	return false
}

func stack(bricks []Brick) int {
	moves := make(map[int]int)
	for i := 0; i < len(bricks); i++ {
		// While the brick isn't on the floor keep checking the levels below
		for bricks[i].z > 1 {
			newBrick := bricks[i]
			newBrick.z--
			newBrick.z2--
			wouldCollide := false
			// See if moving the brick down a Z level would hit another brick
			for j := i - 1; j > -1; j-- {
				if bricks[j].WouldCollide(&newBrick) {
					wouldCollide = true
					break
				}
			}
			if wouldCollide {
				break
			}
			moves[i]++
			bricks[i] = newBrick
		}
	}
	return len(moves)
}

func process(input []string) int {

	// Build the bricks
	var bricks []Brick
	for _, row := range input {
		items := strings.FieldsFunc(row, func(r rune) bool {
			return r == ',' || r == '~'
		})
		brick := Brick{
			x:  shared.Atoi(items[0]),
			y:  shared.Atoi(items[1]),
			z:  shared.Atoi(items[2]),
			x2: shared.Atoi(items[3]),
			y2: shared.Atoi(items[4]),
			z2: shared.Atoi(items[5]),
		}
		bricks = append(bricks, brick)
	}
	// Input isn't in Z order so sort
	slices.SortFunc(bricks, func(a, b Brick) int {
		return min(a.z, a.z2) - min(b.z, b.z2)
	})

	// Stack the bricks
	stack(bricks)

	// Now take each brick out, and see it if it was supporting anything
	noneSupportingBricks := 0
	for i := 0; i < len(bricks); i++ {
		stackCopy := make([]Brick, len(bricks))
		copy(stackCopy, bricks)
		stackCopy = append(stackCopy[:i], stackCopy[i+1:]...)
		changes := stack(stackCopy)
		if changes == 0 {
			noneSupportingBricks++
		}
	}
	return noneSupportingBricks
}

package part2

type star struct {
	id int
	X  int
	Y  int
}

func (s *star) distance(otherStar *star) int {
	return abs(s.X-otherStar.X) + abs(s.Y-otherStar.Y)
}

func process(input []string, expansion int) int {
	var stars []*star
	// Find all the stars
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '#' {
				stars = append(stars, &star{X: j, Y: i, id: len(stars) + 1})
			}
		}
	}
	rowsAdded := 0
	for i := 0; i < len(input[0]); i++ { // Expand the universe (rows)
		blankRow := true
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '#' {
				blankRow = false
				break
			}
		}
		if blankRow {
			for _, s := range stars {
				if s.Y-rowsAdded > i {
					s.Y = s.Y + (expansion - 1)
				}
			}
			rowsAdded = rowsAdded + (expansion - 1)
		}
	}
	colsAdded := 0
	for i := 0; i < len(input[0]); i++ { // Expand the universe (cols)
		blankCol := true
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				blankCol = false
				break
			}
		}
		if blankCol {
			for _, s := range stars {
				if s.X > i+colsAdded {
					s.X = s.X + (expansion - 1)
				}
			}
			colsAdded = colsAdded + (expansion - 1)
		}
	}
	total := 0
	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			total += stars[i].distance(stars[j])
		}
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

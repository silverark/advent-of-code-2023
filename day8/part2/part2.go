package part2

type element struct {
	Id      string
	L       string
	R       string
	isStart bool
	isEnd   bool
}

func process(input []string) int {
	elements := make(map[string]element)
	directions := input[0]
	for i := 2; i < len(input); i++ {
		loc := input[i][0:3]
		l := input[i][7:10]
		r := input[i][12:15]
		elements[loc] = element{
			Id: loc, L: l, R: r,
			isStart: loc[2] == 'A',
			isEnd:   loc[2] == 'Z',
		}
	}
	counter := 0
	// Create a slice of starting elements
	parallelElements := make([]element, 0)
	for _, v := range elements {
		if v.isStart {
			parallelElements = append(parallelElements, v)
		}
	}
	var pathCount []int
	for j := 0; j < len(parallelElements); j++ {
		counter = 0
		// Loop through each element and increment
		for i := 0; i < len(directions); i++ {
			if parallelElements[j].isEnd {
				continue
			}
			counter++
			if directions[i] == 'L' {
				parallelElements[j] = elements[parallelElements[j].L]
			} else {
				parallelElements[j] = elements[parallelElements[j].R]
			}
			if parallelElements[j].Id == "ZZZ" {
				break
			}
			if i == len(directions)-1 {
				i = -1
			}
		}
		pathCount = append(pathCount, counter)
	}
	totalPath := LCM(pathCount[0], pathCount[1], pathCount[2:]...)
	return totalPath
}

// Thank you https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// GCD (greatest common divisor) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM (Least Common Multiple) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

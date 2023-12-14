package part1

import "log"

type element struct {
	Id string
	L  string
	R  string
}

func process(input []string) int {

	elements := make(map[string]element)
	directions := input[0]
	for i := 2; i < len(input); i++ {
		loc := input[i][0:3]
		l := input[i][7:10]
		r := input[i][12:15]
		elements[loc] = element{Id: loc, L: l, R: r}

		log.Println(elements[loc])
	}
	currentElement := elements["AAA"]
	counter := 0
	for i := 0; i < len(directions); i++ {
		counter++
		if directions[i] == 'L' {
			currentElement = elements[currentElement.L]
		} else {
			currentElement = elements[currentElement.R]
		}
		if currentElement.Id == "ZZZ" {
			break
		}
		if i == len(directions)-1 {
			i = -1
		}
	}
	return counter
}

package part1

import (
	"strings"
)

func process(input []string) int {

	// For each row, move it up if hte space above it is a "."

	for i := 1; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'O' {
				//Look at all rows above until it hits something.
				for k := i - 1; k >= 0; k-- {
					// See if above is clear
					if input[k][j] == '.' {
						//log.Println("The line above is clear so we're moving it up.")
						//log.Println("Lines Before:")
						//for l := 0; l <= k+1; l++ {
						//	fmt.Println(input[l])
						//}
						input[k] = input[k][:j] + "O" + input[k][j+1:]
						input[k+1] = input[k+1][:j] + "." + input[k+1][j+1:]

						//log.Println("Lines After:")
						//for l := 0; l <= k+1; l++ {
						//	fmt.Println(input[l])
						//}
					} else {
						break
					}
				}

			}
		}
	}

	//sum values up
	totalWeight := 0
	rows := len(input) - 1
	for i := 0; i <= len(input)-1; i++ {
		// count the `O` in the string
		rocks := strings.Count(input[i], "O")
		totalWeight += rocks * (rows + 1)
		rows--
	}

	return totalWeight
}

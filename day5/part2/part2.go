package part2

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func process(input []string) int {

	var mapFilters []*MapFilter

	// build the filters
	for i := 2; i < len(input)-1; i++ {
		if len(input[i]) == 0 || input[i] == "" {
			continue
		}
		// See if the first character of input[0] is a number
		if unicode.IsDigit(rune(input[i][0])) {
			// Add the filter
			values := strings.Fields(input[i])
			inputStart, _ := strconv.Atoi(values[1])
			filterRange, _ := strconv.Atoi(values[2])
			outputStart, _ := strconv.Atoi(values[0])

			// Add the filter
			currentFilter := mapFilters[len(mapFilters)-1]
			currentFilter.filters = append(currentFilter.filters, filter{
				inputStart: inputStart,
				inputEnd:   inputStart + (filterRange - 1),
				outputDiff: outputStart - inputStart,
			})
		} else {
			newMapFilter := &MapFilter{}
			if len(mapFilters) > 0 {
				mapFilters[len(mapFilters)-1].chain = newMapFilter
			}
			mapFilters = append(mapFilters, newMapFilter)
		}
	}
	// Get maxInt
	lowestNum := int(^uint(0) >> 1)
	startMapFilter := mapFilters[0]
	seedsParts := strings.Fields(strings.Split(input[0], ":")[1])
	for i := 0; i < len(seedsParts)-1; i++ {
		seedVal, _ := strconv.Atoi(seedsParts[i])
		rangeVal, _ := strconv.Atoi(seedsParts[i+1])
		log.Println("Processing range ", seedVal, seedVal+rangeVal-1)
		for j := 0; j <= rangeVal-1; j++ {
			result := startMapFilter.Calculate(seedVal + j)
			if result < lowestNum {
				lowestNum = result
			}
		}
		i++
	}

	return lowestNum
}

type MapFilter struct {
	filters []filter
	chain   *MapFilter
}

func (m *MapFilter) Calculate(input int) int {
	output := input
	// Convert the number by looking at local filters
	for _, localFilter := range m.filters {
		if input >= localFilter.inputStart && input <= localFilter.inputEnd {
			output = input + localFilter.outputDiff
			break
		}
	}
	// now pass to the chain
	if m.chain != nil {
		output = m.chain.Calculate(output)
	}
	return output
}

type filter struct {
	inputStart int
	inputEnd   int
	outputDiff int
}

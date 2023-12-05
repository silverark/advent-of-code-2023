package part1

import (
	"strconv"
	"strings"
	"unicode"
)

func process(input []string) int {
	var mapFilters []*MapFilter

	// Build the filters
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
	lowestNum := int(^uint(0) >> 1)
	startFilter := mapFilters[0]

	// Get the seeds
	seedsPart := strings.Fields(strings.Split(input[0], ":")[1])
	for _, seed := range seedsPart {
		seedVal, _ := strconv.Atoi(seed)
		result := startFilter.Calculate(seedVal)
		if result < lowestNum {
			lowestNum = result
		}
	}
	return lowestNum
}

type MapFilter struct {
	filters []filter
	chain   *MapFilter
}

func (m *MapFilter) Calculate(input int) int {
	output := input
	for _, localFilter := range m.filters {
		if input >= localFilter.inputStart && input <= localFilter.inputEnd {
			output = input + localFilter.outputDiff
			break
		}
	}
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

package part2

import (
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

	seedRanges := []*ItemRange{}

	// Build the seed ranges
	for i := 0; i < len(seedsParts)-1; i++ {
		seedVal, _ := strconv.Atoi(seedsParts[i])
		rangeVal, _ := strconv.Atoi(seedsParts[i+1])

		seedRanges = append(seedRanges, &ItemRange{
			start: seedVal,
			end:   seedVal + rangeVal - 1,
		})
		i++
	}

	result := startMapFilter.Calculate(seedRanges)

	for _, seedRange := range result {
		if seedRange.start < lowestNum {
			lowestNum = seedRange.start
		}
	}

	return lowestNum
}

type ItemRange struct {
	start     int
	end       int
	processed bool
}

type MapFilter struct {
	filters []filter
	chain   *MapFilter
}

func (m *MapFilter) Calculate(input []*ItemRange) (outputRanges []*ItemRange) {

PROCESS_INPUT:
	for _, inputRange := range input {
		if inputRange.processed {
			continue
		}

		// See if this range is covered by any of the filters
		for _, localFilter := range m.filters {
			// Check for exact Fit.
			if inputRange.start >= localFilter.inputStart && inputRange.end <= localFilter.inputEnd {
				outputRanges = append(outputRanges, &ItemRange{
					start: inputRange.start + localFilter.outputDiff,
					end:   inputRange.end + localFilter.outputDiff,
				})
				inputRange.processed = true
				break
			}

			// Check for overlapping
			if inputRange.start >= localFilter.inputStart && inputRange.start <= localFilter.inputEnd {
				outputRange := ItemRange{
					start: inputRange.start + localFilter.outputDiff,
					end:   localFilter.inputEnd + localFilter.outputDiff,
				}
				inputRange.processed = true
				// Need to split the output ranges and add them to the output
				outputRanges = append(outputRanges, &outputRange)
				// Now need to pass remaining part of the range to the end of input.
				remainingRange := ItemRange{
					start: localFilter.inputEnd + 1,
					end:   inputRange.end,
				}
				//log.Println("Remaining range: ", remainingRange)
				input = append(input, &remainingRange)
				break
			}

		}

		if !inputRange.processed {
			inputRange.processed = true
			// If we're here it hasn't matched any filters to add it to the output filter directly and mark as complete
			outputRanges = append(outputRanges, &ItemRange{
				start:     inputRange.start,
				end:       inputRange.end,
				processed: false,
			})
		}
	}

	// If there are any inputs not processed loop back to the top until they have. This will keep splitting on filters where needed
	for _, inputRange := range input {
		if !inputRange.processed {
			goto PROCESS_INPUT
		}
	}

	// now pass to the chain
	if m.chain != nil {
		outputRanges = m.chain.Calculate(outputRanges)
	}
	return outputRanges
}

type filter struct {
	inputStart int
	inputEnd   int
	outputDiff int
}

package main

import (

)

// part 1

// problem summary:
// - each character in a line is a single digit (1-9) representing a batteries joltage rating
// - for each bank, turn on two batteries
// - cannot reorder digits and have to respect their order in the string
// - for each bank, want the max possible two-digit value
// - answr is the sum over all banks of their max two-digit values

func sumMaxJoltageTwo(lines []string) int {
	// counter to keep track of total joltage
	total := 0

	for _, line := range lines {
		if len(line) < 2 {
			// a bank with < 2 batteries cant form a two-digit number
			continue
		}

		// store the best (largest) two-digit number for this line
		best := -1

		// try all pairs (i, j) with i < j
		for i := 0; i < len(line); i++ {
			// convert rune '0'...'9' to its int value
			tens := int(line[i] - '0')

			for j := i + 1; j < len(line); j++ {
				ones := int(line[j] - '0')

				// form the two digit number
				value := tens * 10 + ones

				if value > best {
					best = value
				}
			}
		}

		// add this banks best possible two-digit joltage to the global total
		total += best
	}

	return total
}
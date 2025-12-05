package main



// part 2

// problem summary:
// - each character in a line is a single digit (1-9) representing a batteries joltage rating
// - for each bank, turn on twelve batteries
// - cannot reorder digits and have to respect their order in the string
// - for each bank, want the max possible twelve-digit value
// - answr is the sum over all banks of their max twelve-digit values

// sumMaxJoltageTwelve computes the total maximum joltage for all banks,
// when need to chose exactly 12 digits per bank
func sumMaxJoltageTwelve(lines []string) int {
	// must pick exactly 12 digits
	const choose = 12

	// int64 extra safe
	var total int64

	for _, line := range lines {
		// if a line is shorter than 12, cannot form a 12-digit number
		if len(line) < choose {
			continue
		}

		n := len(line)
		// how many digits we are allowed to remove
		remove := n - choose

		// stack will store the chosen digits as bytes ('0'..'9')
		stack := make([]byte, 0, n)

		// process each digit from left to right
		for i := 0; i < n; i++ {
			// current digit as a byte
			d := line[i]

			// while can still remove digits, the stack is not empty, and the last
			// digit in the stack is smaller than the current digit, then pop the
			// smaller digit to make room for the larger one
			for remove > 0 && len(stack) > 0 && stack[len(stack) - 1] < d {
				stack = stack[:len(stack) - 1]
				remove--
			}

			// keep current digit
			stack = append(stack, d)
		}

		// if there are digits still left to remove, remove them from the end
		if remove > 0 {
			stack = stack[:len(stack) - remove]
		}

		// now stack contains largest sequence (should be 'choose' digits long)
		var value int64
		for i := 0; i < choose; i++ {
			value = value * 10 + int64(stack[i] - '0')
		}

		total += value
	}

	// convert to int for printing
	return int(total)
}
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// part 1

// problem summary:
// - the input is a single line containing ranges of product id's, seperrated by commas
// - each range looks like 'start-end' and includes all integers from start to end (inclusive)

// ex:
//  "11-22" means the id's 11, 12, 13, ..., 22

// - any invalid id is any number whose decimal representation is made by taking some non-empty
//   sequence of digits X and repeating it exactly twice

// ex:
//  55 -> "5" + "5" (invalid)
//  123123 -> "123" + "123" (invalid)

// - the decimal representation must have even length and no leading zeros

// sumInvalidIDs computes the sum of all invalid id's in a given input
func sumInvalidIDs(input string) int64 {
	// clean up input
	line := strings.TrimSpace(input)
	if line == "" {
		// no ranges at all
		return 0
	}

	// split the line into individual range strings using the comma seperator
	rawRanges := strings.Split(line, ",")

	// counter for the sum of all invalid id's
	var total int64 = 0

	// loop over each 'start-end' range string
	for _, r := range rawRanges {
		r = strings.TrimSpace(r)
		if r == "" {
			// skip empty pieces
			continue
		}

		// parse the "start-end" into two integers
		start, end, err := parseRange(r)
		if err != nil {
			fmt.Printf("warning: skipping bad range %q: %v\n", r, err)
			continue
		}

		// for each id in the range, check fi its an invalid id,
		// and if so, add it to the total
		for id := start; id <= end; id++ {
			if isDoubleRepeatID(id) {
				total += id
			}
		}
	}

	return total
}

// parseRange takes a string like "11-22" and returns (11, 22)
// if the string doesnt match the expected format, it returns an error
func parseRange(str string) (int64, int64, error) {
	// split on the dash
	parts := strings.Split(str, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("range %q does not contain exactly one '-'", str)
	}

	// convert both sides to integers
	startStr := strings.TrimSpace(parts[0])
	endStr := strings.TrimSpace(parts[1])

	start, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start %q: %w", startStr, err)
	} 

	end, err := strconv.ParseInt(endStr, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid end %q: %w", startStr, err)
	}

	if end < start {
		return 0, 0, fmt.Errorf("end (%d) < start (%d)", end, start)
	}

	return start, end, nil
}

// isDoubleRepeatID returns true if the decimal representation of num
// is made by repeating some non-empty digit sequence exactly twice
func isDoubleRepeatID(num int64) bool {
	// convert the integer to its base-10 string representation
	// ex: 123456 -> "123456"
	str := strconv.FormatInt(num, 10)

	// repeatable numbers (invalid id's) must be even numbers
	if len(str) % 2 != 0 {
		return false
	}

	// half is the length of each repeated block
	half := len(str) / 2

	left := str[:half]
	right := str[half:]

	// its an invalid id if and only if the two halves are exactly equal
	return left == right
}
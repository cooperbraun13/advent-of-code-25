package main

import (
	"strings"
	"fmt"
	"strconv"
)

// problem summary:
// - use the same id ranges as part 1, expect now an invalid id is when its digits are
//   made by repeating some digit sequence at least twice (e.g., 123123123, 11111111)
// - sum all such invalid id's across all ranges

// - going to reuse parseRange from part 1

// sumInvalidIDsPart2 computes the sum of all invalid id's using the new rules stated in
// the problem summary
func sumInvalidIDsPart2(input string) int64 {
	// clean up the input
	line := strings.TrimSpace(input)
	if line == "" { 
		return 0 
	}

	// split into "start-end" chunks on commas
	rawRanges := strings.Split(line, ",")

	// counter
	var total int64 = 0

	// loop through every "start-end" range
	for _, r := range rawRanges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}

		// parse "start-end" into numeric start/end
		start, end, err := parseRange(r)
		if err != nil {
			fmt.Printf("warning: skipping bad range %q: %v\n", r, err)
		}

		// loop from start to end
		for id := start; id <= end; id++ {
			if isRepeatedPatternAtLeastTwice(id) {
				total += id
			}
		}
	}

	return total
}


// isRepeatedPatternAtLeastTwice returns true if the decimal representation of num
// is some substring repeated at a minimum of two times
func isRepeatedPatternAtLeastTwice(num int64) bool {
	// convert the integer to its base-10 string representation
	// ex: 123456 -> "123456"
	str := strconv.FormatInt(num, 10)
	n := len(str)

	// need at least two repeats of some block
	// there exists a block length d such that:
	// - d divides n (n % d == 0)
	// - n / d >= 2 (at least 2 copies)
	// - str == (str[0:d] repeated n / d times)

	// we can brute-force over all possible block lengths d from 1 up to n/2
	for blockLen := 1; blockLen <= n / 2; blockLen++ {
		// if blockLen doesnt evenly divide n, we cant tile the string with copies of that block
		if n % blockLen != 0 {
			continue
		}

		// number of repeats if we use this block length
		repeats := n / blockLen
		if repeats < 2 {
			// redundant given blockLen <= n/2
			continue
		}

		// the candidate block is the first blockLen characters
		block := str[:blockLen]

		// assume its good until we find a mismatch
		ok := true

		// check each subsequent block of length blockLen to see if it matches 'block'
		for pos := blockLen; pos < n; pos += blockLen {
			if str[pos:pos+blockLen] != block {
				ok = false
				break
			}
		}

		if ok {
			// we found at least one decomposition of str as X repeated k >= 2 times
			return true
		}
	}

	// no suitable repeating pattern found
	return false
}
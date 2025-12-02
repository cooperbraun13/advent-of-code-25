package main

import (
	"fmt"
	"os"
)

// advent of code 2025 day 2

func main() {
	// read the whole file into memory
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error reading input:", err)
		return
	}

	// part 1
	part1 := sumInvalidIDs(string(data))
	fmt.Println("part 1:", part1)

	// part 2
	part2 := sumInvalidIDsPart2(string(data))
	fmt.Println("part 2:", part2)
}
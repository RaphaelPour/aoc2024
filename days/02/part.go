package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

func isMonotonic(report []int, minDiff, maxDiff int) bool {
	for i := 0; i < len(report)-1; i += 1 {
		diff := report[i+1] - report[i]
		fmt.Println(report[i+1], report[i], diff)
		if diff < minDiff || diff > maxDiff {
			fmt.Println("nope")
			return false
		}
	}
	return true
}

func part1(data [][]int) int {
	safe := 0
	for _, line := range data {
		if isMonotonic(line, 1, 3) || isMonotonic(line, -3, -1) {
			safe += 1
		}
	}
	return safe
}

func part2(data [][]int) int {
	safe := 0
	for _, line := range data {
		if isMonotonic(line, 1, 3) || isMonotonic(line, -3, -1) {
			safe += 1
			continue
		}

		line2 := make([]int, len(line)-1)
		//https://go.dev/play/p/-k4lZnPMSgg
		for i := 0; i < len(line); i += 1 {
			copy(line2[i:i+1], line[i:i+1])
			copy(line2[i+1:], line[i+2:])
			if isMonotonic(line2, 1, 3) || isMonotonic(line2, -3, -1) {
				safe += 1
				continue
			}
		}
	}
	return safe
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadIntTable("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}

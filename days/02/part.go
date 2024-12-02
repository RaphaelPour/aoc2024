package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/generator"
	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
)

func part1(data [][]int) int {
	safe := 0
	for _, line := range data {
		if math.IsMonotonicWithinRange(line, 1, 3) || math.IsMonotonicWithinRange(line, -3, -1) {
			safe += 1
		}
	}
	return safe
}

func part2(data [][]int) int {
	safe := 0
	for _, line := range data {
		if math.IsMonotonicWithinRange(line, 1, 3) || math.IsMonotonicWithinRange(line, -3, -1) {
			safe += 1
			continue
		}

		for line2 := range generator.SkipOneSeq(line) {
			if math.IsMonotonicWithinRange(line2, 1, 3) || math.IsMonotonicWithinRange(line2, -3, -1) {
				safe += 1
				break
			}
		}
	}
	return safe
}

func main() {
	data := input.LoadIntTable("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}

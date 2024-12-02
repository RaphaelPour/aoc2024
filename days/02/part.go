package main

import (
	"fmt"
	"iter"
	"slices"

	"github.com/RaphaelPour/stellar/input"
)

func SkipOne[S ~[]T, T any](in S) []S {
	return slices.Collect(SkipOneSeq(in))
}

func SkipOneSeq[S ~[]T, T any](in S) iter.Seq[S] {
	return func(yield func(S) bool) {
		out := make(S, len(in)-1)
		for i := 0; i < len(in); i += 1 {
			copy(out[:i], in[:i])
			copy(out[i:], in[i+1:])
			if !yield(out) {
				return
			}
		}
	}
}

func SkipOneSeq2[S ~[]T, T any](in S) iter.Seq2[int, S] {
	return func(yield func(int, S) bool) {
		out := make(S, len(in)-1)
		for i := 0; i < len(in); i += 1 {
			copy(out[:i], in[:i])
			copy(out[i:], in[i+1:])
			if !yield(i, out) {
				return
			}
		}
	}
}

func isMonotonic(report []int, minDiff, maxDiff int) bool {
	for i := 0; i < len(report)-1; i += 1 {
		diff := report[i+1] - report[i]
		if diff < minDiff || diff > maxDiff {
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

		//https://go.dev/play/p/-k4lZnPMSgg
		for line2 := range SkipOneSeq(line) {
			if isMonotonic(line2, 1, 3) || isMonotonic(line2, -3, -1) {
				safe += 1
				break
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

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
)

func part1(data []string) int {
	left := make([]int, len(data))
	right := make([]int, len(data))
	for i, line := range data {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("error parsing line", line)
			return -1
		}

		numLeft, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic(fmt.Sprintf("'%s' is not a number\n", parts[0]))
		}
		numRight, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			panic(fmt.Sprintf("'%s' is not a number\n", parts[1]))
		}

		left[i] = numLeft
		right[i] = numRight
	}

	sort.Ints(left)
	sort.Ints(right)

	var sum int
	for i := range left {
		sum += math.Abs(left[i] - right[i])
	}

	return sum
}

func part2(data []string) int {
	hist := make(map[int]int)
	left := make([]int, len(data))

	for i, line := range data {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("error parsing line", line)
			return -1
		}

		numLeft, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic(fmt.Sprintf("'%s' is not a number\n", parts[0]))
		}
		numRight, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			panic(fmt.Sprintf("'%s' is not a number\n", parts[1]))
		}

		left[i] = numLeft
		hist[numRight] = hist[numRight] + 1
	}

	var sum int
	for _, num := range left {
		sum += num * hist[num]
	}

	return sum
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}

package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	smath "github.com/RaphaelPour/stellar/math"
	sstrings "github.com/RaphaelPour/stellar/strings"
)

func M[in any, out any, inArr []in, outArr []out](input inArr, fn func(in) out) outArr {
	output := make(outArr, len(input))
	for i := range input {
		output[i] = fn(input[i])
	}
	return output
}

func part1(data []string) int {
	sum := 0
	for _, line := range data {
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			fmt.Println("unknown field count", len(fields), fields)
			return 0
		}

		testResult := sstrings.ToInt(strings.TrimSuffix(fields[0], ":"))
		numbers := M(strings.Fields(fields[1]), sstrings.ToInt)
		fmt.Println(numbers)

		found := false
		for op := 0; op <= smath.Pow(2, len(numbers)-1); op++ {
			result := numbers[0]
			fmt.Print(result)
			for i, num := range numbers[1:] {
				if op&(1<<(i)) == 0 {
					result += num
					fmt.Printf(" + %d", num)
				} else {
					result *= num
					fmt.Printf(" * %d", num)
				}
			}

			fmt.Printf("= %d\n", result)

			if result == testResult {
				found = true
				break
			}
		}
		if found {
			sum += testResult
		}
	}
	return sum
}

func part2(data []string) int {
	return 0
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	result := part1(data)
	if result <= 102900683395 {
		fmt.Printf("%d is too low", result)
	} else {
		fmt.Println(result)
	}

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}

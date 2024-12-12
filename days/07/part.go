package main

import (
	"fmt"
	"strconv"
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

		found := false
		for op := 0; op <= smath.Pow(2, len(numbers)-1); op++ {
			result := numbers[0]
			for i, num := range numbers[1:] {
				if op&(1<<(i)) == 0 {
					result += num
				} else {
					result *= num
				}
			}

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

func Default[S []T, T any](index int, arr S, def T) T {
	if len(arr) > index {
		return arr[index]
	}
	return def
}

func Concat(a, b int) int {
	return sstrings.ToInt(strconv.Itoa(a) + strconv.Itoa(b))
}

func R(result, goal int, numbers []int) bool {
	// result is growing monotonically, if
	// too big, goal can't be reached anymore
	if result > goal {
		return false
	}

	// no numbers left, goal must be reached or it never will
	if len(numbers) == 0 {
		// goal reached \o/ ?
		return result == goal
	}

	if R(result*numbers[0], goal, numbers[1:]) {
		return true
	} else if R(Concat(result, numbers[0]), goal, numbers[1:]) {
		return true
	}

	return R(result+numbers[0], goal, numbers[1:])
}

func part2(data []string) int {
	sum := 0
	for _, line := range data {
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			fmt.Println("unknown field count", len(fields), fields)
			return 0
		}
		testResult := sstrings.ToInt(strings.TrimSuffix(fields[0], ":"))
		numbers := M(strings.Fields(fields[1]), sstrings.ToInt)
		if R(numbers[0], testResult, numbers[1:]) {
			sum += testResult
		}
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

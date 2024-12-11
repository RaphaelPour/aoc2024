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

func Default[S []T, T any](index int, arr S, def T) T {
	if len(arr) > index {
		return arr[index]
	}
	return def
}

func Concat(a, b int) int {
	return sstrings.ToInt(strconv.Itoa(a) + strconv.Itoa(b))
}

func R(current, result, goal int, path string, numbers []int) bool {
	// result is growing monotonically, if
	// too big, goal can't be reached anymore
	if result > goal {
		fmt.Printf("%s%d=%d (%t)\n", path, current, result, false)
		return false
	}

	// no numbers left, goal must be reached or it never will
	if len(numbers) == 0 {
		// goal reached \o/
		fmt.Printf("%s%d=%d (%t)\n", path, current, result, result == goal)
		return result == goal
	}

	next := -1
	if len(numbers) > 0 {
		next = numbers[0]
	} else {
		next = current
	}

	mul := result
	if path == "" {
		mul = 1
	}

	newPath := fmt.Sprintf("%s%d", path, current)
	if R(next, mul*current, goal, fmt.Sprintf("%s*", newPath), numbers[1:]) {
		return true
	}
	//else if R(Concat(current, next), result, goal, fmt.Sprintf("%s(%d||%d) ", path, current, next), numbers[1:]) {
	//	return true
	//}

	return R(next, result+current, goal, fmt.Sprintf("%s+", newPath), numbers[1:])
}

func part2(data []string) int {
	sum := 0
	for i, line := range data {
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			fmt.Println("unknown field count", len(fields), fields)
			return 0
		}
		testResult := sstrings.ToInt(strings.TrimSuffix(fields[0], ":"))
		numbers := M(strings.Fields(fields[1]), sstrings.ToInt)
		fmt.Println(testResult)
		if R(numbers[0], 0, testResult, "", numbers[1:]) {
			sum += testResult
		}
		fmt.Println(100.0 / float64(len(data)) * float64(i+1))
		break
	}
	return sum
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadString("input_example")

	fmt.Println("== [ PART 1 ] ==")
	// result := part1(data)
	//if result <= 102900683395 {
	//	fmt.Printf("%d is too low", result)
	//} else {
	//	fmt.Println(result)
	//}

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
	fmt.Println("too low: 5905560574779")
}

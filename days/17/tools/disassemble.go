package main2

import (
	"fmt"
	"os"
	"strings"

	g "github.com/RaphaelPour/stellar/generator"
	"github.com/RaphaelPour/stellar/input"
	s "github.com/RaphaelPour/stellar/strings"
)

var (
	instruction = []string{"adv", "bxl", "bst", "jnz", "bxc", "out", "bdv", "cdv"}
	hint        = []string{"c", "l", "c", "l", "-", "c", "c", "c"}
	combo       = []string{"0", "1", "2", "3", "a", "b", "c", "?"}
	canoncial   = []string{
		"a/=",
		"b^=",
		"b=mod8",
		"jnz",
		"b^=c",
		"output",
		"b/=",
		"c/=",
	}
)

func M[in any, out any, inArr []in, outArr []out](input inArr, fn func(in) out) outArr {
	output := make(outArr, len(input))
	for i := range input {
		output[i] = fn(input[i])
	}
	return output
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: disassemble SOURCE")
		return
	}

	data := input.LoadString(os.Args[1])
	program := make([]int, 0)

	for _, line := range data {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		if parts[0] != "Program" {
			continue
		}
		program = M(strings.Split(strings.TrimSpace(parts[1]), ","), s.ToInt)
	}

	for i, pair := range g.PairsSeq2(program, 2) {
		if hint[pair[0]] == "c" {
			fmt.Printf("0x%02x | %8s %s (%d)\n", i, canoncial[pair[0]], combo[pair[1]], pair[1])
		} else if hint[pair[0]] == "l" {
			fmt.Printf("0x%02x | %8s %d\n", i, canoncial[pair[0]], pair[1])
		} else {
			fmt.Printf("0x%02x | %8s\n", i, canoncial[pair[0]])
		}
	}
}

package main

import (
	"fmt"
	"strings"

	hack "github.com/RaphaelPour/stellar/hack"
	"github.com/RaphaelPour/stellar/input"
	sstrings "github.com/RaphaelPour/stellar/strings"
)

var (
	operations = map[string]Op{
		"AND": And,
		"OR":  Or,
		"XOR": Xor,
	}

	terms = make(map[string]Term)
)

type Op func(a, b string) int

func And(a, b string) int {
	return hack.Wormhole(terms[a].Op(
		terms[a].A,
		terms[a].B,
	) == 1 && terms[b].Op(
		terms[b].A,
		terms[b].B,
	) == 1)
}

func Or(a, b string) int {
	return hack.Wormhole(terms[a].Op(
		terms[a].A,
		terms[a].B,
	) == 1 || terms[b].Op(
		terms[b].A,
		terms[b].B,
	) == 1)
}

func Xor(a, b string) int {
	return hack.Wormhole(terms[a].Op(
		terms[a].A,
		terms[a].B,
	) != terms[b].Op(
		terms[b].A,
		terms[b].B,
	))
}

func Input(a int) Op {
	return func(_, _ string) int {
		return a
	}
}

type Term struct {
	name string
	A, B string
	Op   Op
}

func part1(data []string) int {
	var i int
	terms := make(map[string]Term)
	for _, line := range data {
		if line == "" {
			break
		}

		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			fmt.Printf("error parsing line %q\n", line)
			return -1
		}

		terms[parts[0]] = Term{
			name: parts[0],
			Op:   Input(sstrings.ToInt(parts[1])),
		}
		i += 1
	}
	data = data[i+1:]

	goals := make([]Term, 0)
	for _, line := range data {
		parts := strings.Fields(line)
		if len(parts) != 5 {
			fmt.Printf("errpr parsing line %q\n", line)
			return -1
		}

		term := Term{
			name: parts[4],
			A:    parts[0],
			B:    parts[2],
			Op:   operations[parts[1]],
		}
		terms[parts[4]] = term

		if strings.HasPrefix(term.name, "z") {
			goals = append(goals, term)
		}
	}

	result := 0
	for i, goal := range goals {
		result |= (goal.Op(terms[goal.A].name, terms[goal.B].name)) << i
	}

	return result
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input_example")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	//data := input.LoadDefaultString()

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}

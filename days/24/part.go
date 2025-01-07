package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/RaphaelPour/stellar/hack"
	"github.com/RaphaelPour/stellar/input"
)

type ExprKind int

const (
	CONST ExprKind = iota
	AND
	OR
	XOR
)

func Str2Kind(s string) ExprKind {
	switch s {
	case "AND":
		return AND
	case "OR":
		return OR
	case "XOR":
		return XOR
	default:
		panic(fmt.Sprintf("unknown expression kind %q", s))
	}
}

type Expr struct {
	InputA, InputB string
	Value          bool
	Kind           ExprKind
}

func Eval(input string, expressions map[string]Expr) bool {
	node, ok := expressions[input]
	if !ok {
		panic(fmt.Sprintf("node %q not found in expression map", input))
	}

	switch node.Kind {
	case CONST:
		return node.Value
	case AND:
		return Eval(node.InputA, expressions) && Eval(node.InputB, expressions)
	case OR:
		return Eval(node.InputA, expressions) || Eval(node.InputB, expressions)
	case XOR:
		return Eval(node.InputA, expressions) != Eval(node.InputB, expressions)
	default:
		panic(fmt.Sprintf("unknown expression kind %q", node.Kind))
	}
}

func Parse(data []string) (map[string]Expr, []string) {
	var i int
	expressions := make(map[string]Expr)

	for _, line := range data {
		if line == "" {
			break
		}

		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			fmt.Printf("error parsing line %q\n", line)
			return nil, nil
		}
		expressions[parts[0]] = Expr{
			Kind:  CONST,
			Value: parts[1] == "1",
		}
		i += 1
	}
	data = data[i+1:]

	outputs := make([]string, 0)
	for _, line := range data {
		parts := strings.Fields(line)
		if len(parts) != 5 {
			fmt.Printf("errpr parsing line %q\n", line)
			return nil, nil
		}

		if strings.HasPrefix(parts[4], "z") {
			outputs = append(outputs, parts[4])
		}

		// x01 AND x02 -> y02
		// 0   1   2   3  4
		expressions[parts[4]] = Expr{
			Kind:   Str2Kind(parts[1]),
			InputA: parts[0],
			InputB: parts[2],
		}
	}

	sort.Strings(outputs)

	return expressions, outputs
}

func part1(data []string) int {
	expressions, outputs := Parse(data)
	var result int
	// iter through all output nodes and resolve them recursively
	for i, node := range outputs {
		result |= hack.Wormhole(Eval(node, expressions)) << i
	}

	return result
}

func Number(prefix string, expressions map[string]Expr) int {
	keys := make([]string, 0)
	for key := range expressions {
		if strings.HasPrefix(key, prefix) {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	var result int
	for i, key := range keys {
		result |= hack.Wormhole(expressions[key].Value) << i
	}
	return result
}

func Set(prefix string, i int, expressions map[string]Expr) {
	keys := make([]string, 0)
	for key := range expressions {
		if strings.HasPrefix(key, prefix) {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	for j, key := range keys {
		e := expressions[key]
		e.Value = (i == j)
		expressions[key] = e
	}
}

func part2(data []string) int {
	expressions, outputs := Parse(data)

	for i := 0; i < 44; i++ {
		Set("x", i, expressions)
		Set("y", i, expressions)

		var result int
		for i, node := range outputs {
			result |= hack.Wormhole(Eval(node, expressions)) << i
		}

		if result != Number("x", expressions)+Number("y", expressions) {
			fmt.Printf(
				"found bad sum x=%d + y=%d = %d\n",
				Number("x", expressions),
				Number("y", expressions),
				result,
			)
		}
	}
	return 0
}

func main() {
	data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	//data := input.LoadDefaultString()

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}

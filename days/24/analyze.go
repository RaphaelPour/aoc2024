package main

// go run analyze.go | dot -odiagram.png -Tpng

import (
	"fmt"
	"sort"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

type ExprKind int

const (
	CONST ExprKind = iota
	AND
	OR
	XOR
	NUM_KINDS
)

func (e ExprKind) String() string {
	switch e {
	case CONST:
		return "CONST"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case XOR:
		return "XOR"
	default:
		return "??"
	}
}

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

func main() {
	expressions, _ := Parse(input.LoadString("input"))

	fmt.Println("digraph D {\n")

	kindsIndex := make([]int, NUM_KINDS)
	for _, expr := range expressions {
		i := kindsIndex[expr.Kind]
		if expr.Kind == CONST {
			fmt.Printf("\t%s%d[label=\"%t\",shape=box]\n", expr.Kind, i, expr.Value)
		} else {
			fmt.Printf("\t%s%d[label=\"%s\",shape=box]\n", expr.Kind, i, expr.Kind)
			fmt.Printf("\t%s\n", expr.InputA)
			fmt.Printf("\t%s\n", expr.InputB)
		}

		kindsIndex[expr.Kind] += 1
	}
	fmt.Println("")

	kindsIndex = make([]int, NUM_KINDS)
	for key, expr := range expressions {
		i := kindsIndex[expr.Kind]

		fmt.Printf("\t%s%d -> %s\n", expr.Kind, i, key)
		if expr.Kind != CONST {
			fmt.Printf("\t%s -> %s%d\n", expr.InputA, expr.Kind, i)
			fmt.Printf("\t%s -> %s%d\n", expr.InputB, expr.Kind, i)
		}

		kindsIndex[expr.Kind] += 1
	}

	fmt.Println("}")
}

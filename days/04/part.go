package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	sstrings "github.com/RaphaelPour/stellar/strings"
)

var (
	pattern = regexp.MustCompile(`(XMAS|SMAX)`)
	grid    Grid
)

type Grid struct {
	fields [][]string
	marked map[int]int
}

func NewGrid(data []string) Grid {
	g := Grid{}
	g.fields = make([][]string, len(data))
	for y, line := range data {
		g.fields[y] = make([]string, len(line)+1)
		for x, ch := range line {
			g.fields[y][x] = string(ch)
		}
		g.fields[y][len(line)] = "_"
	}

	g.marked = make(map[int]int)

	return g
}

func (g Grid) String() string {
	fmt.Println(g.marked)
	out := ""
	for y := 0; y < len(g.fields); y++ {
		for x := 0; x < len(g.fields); x++ {
			if count, ok := g.marked[x+(len(g.fields)*y)]; ok {
				fmt.Println(x, len(g.fields)*y)
				out += "\033[0;32m" + fmt.Sprintf("\033[0;%dm%s\033[0m ", 31+count, g.fields[y][x])
			} else {
				out += g.fields[y][x] + " "
			}
		}
		out += "\n"
	}
	return out
}

func diag(total string, length int) int {
	expr1 := fmt.Sprintf(
		"(X%sM%sA%sS)",
		strings.Repeat(".", length),
		strings.Repeat(".", length),
		strings.Repeat(".", length),
	)
	expr2 := fmt.Sprintf(
		"(S%sA%sM%sX)",
		strings.Repeat(".", length),
		strings.Repeat(".", length),
		strings.Repeat(".", length),
	)
	return diag_(total, expr1, length) + diag_(sstrings.Reverse(total), expr2, length)
}
func diag_(total, expr string, length int) int {
	fmt.Println(expr)
	patternDiag := regexp.MustCompile(expr)

	index := 1
	sum := 0
	for {
		fmt.Println("rest: ", total[index-1:])
		match := patternDiag.FindStringIndex(total[index-1:])
		//fmt.Println(match)
		if match == nil {
			break
		}

		sum += 1
		index += match[0] + 1

		i := index - 2
		grid.marked[i] = grid.marked[i] + 1
		grid.marked[i+length*1+1] = grid.marked[i+length*1+1] + 1
		grid.marked[i+length*2+2] = grid.marked[i+length*2+2] + 1
		grid.marked[i+length*3+3] = grid.marked[i+length*3+3] + 1
		fmt.Println("match=", match, "index=", index)
	}
	//fmt.Println(sum)
	return sum
}

func part1(data []string) int {
	grid = NewGrid(data)
	//g := NewGrid(data)
	//fmt.Println(g)
	total := strings.Join(data, " ")
	fmt.Println(total)

	sum := diag(total, 0) + diag(total, len(data)+2) + diag(total, len(data)) + diag(total, len(data)-1)

	return sum
}

func part2(data []string) int {
	return 0
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadString("input_example")

	fmt.Println("too low: 2445")
	fmt.Println("too high: 2529")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))
	//fmt.Println(grid.String())

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}

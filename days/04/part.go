package main

import (
	"fmt"
	"iter"

	"github.com/RaphaelPour/stellar/input"
)

var (
	/* XMAS */
	horizontal = []Point{
		Point{0, 0, "X"}, Point{1, 0, "M"}, Point{2, 0, "A"}, Point{3, 0, "S"},
	}
	/* X
	 * M
	 * A
	 * S
	 */
	vertical = []Point{
		Point{0, 0, "X"}, Point{0, 1, "M"}, Point{0, 2, "A"}, Point{0, 3, "S"},
	}
	/* X
	 *  M
	 *   A
	 *    S
	 */
	diagonalDesc = []Point{
		Point{0, 0, "X"}, Point{1, 1, "M"}, Point{2, 2, "A"}, Point{3, 3, "S"},
	}
	/*    X
	 *   M
	 *  A
	 * S
	 */
	diagonalAsc = []Point{
		Point{0, 0, "X"}, Point{1, -1, "M"}, Point{2, -2, "A"}, Point{3, -3, "S"},
	}
	/* SAMX */
	horizontalR = []Point{
		Point{0, 0, "S"}, Point{1, 0, "A"}, Point{2, 0, "M"}, Point{3, 0, "X"},
	}
	/* S
	 * A
	 * M
	 * X
	 */
	verticalR = []Point{
		Point{0, 0, "S"}, Point{0, 1, "A"}, Point{0, 2, "M"}, Point{0, 3, "X"},
	}
	/* S
	 *  A
	 *   M
	 *    X
	 */
	diagonalDescR = []Point{
		Point{0, 0, "S"}, Point{1, 1, "A"}, Point{2, 2, "M"}, Point{3, 3, "X"},
	}
	/*    S
	 *   A
	 *  M
	 * X
	 */
	diagonalAscR = []Point{
		Point{0, 0, "S"}, Point{1, -1, "A"}, Point{2, -2, "M"}, Point{3, -3, "X"},
	}
	allChecks = [][]Point{
		horizontal, vertical, diagonalDesc, diagonalAsc,
		horizontalR, verticalR, diagonalDescR, diagonalAscR,
	}

	/*  M S
	 *   A
	 *  M S
	 */
	x0 = []Point{
		Point{-1, -1, "M"}, Point{0, 0, "A"}, Point{1, 1, "S"},
		Point{-1, 1, "M"}, Point{1, -1, "S"},
	}
	/*  S M
	 *   A
	 *  S M
	 */
	x1 = []Point{
		Point{-1, -1, "S"}, Point{0, 0, "A"}, Point{1, 1, "M"},
		Point{-1, 1, "S"}, Point{1, -1, "M"},
	}
	/*  M M
	 *   A
	 *  S S
	 */
	x2 = []Point{
		Point{-1, -1, "M"}, Point{0, 0, "A"}, Point{1, 1, "S"},
		Point{-1, 1, "S"}, Point{1, -1, "M"},
	}
	/*  S S
	 *   A
	 *  M M
	 */
	x3 = []Point{
		Point{-1, -1, "S"}, Point{0, 0, "A"}, Point{1, 1, "M"},
		Point{-1, 1, "M"}, Point{1, -1, "S"},
	}
	allX = [][]Point{
		x0, x1, x2, x3,
	}
)

type Point struct {
	x, y int
	ch   string
}

func (p Point) Add(x, y int) Point {
	p.x += x
	p.y += y
	return p
}

type Grid struct {
	// parsed input
	fields [][]string

	// count of matches for each field
	marked [][]int
}

func NewGrid(data []string) Grid {
	g := Grid{}
	g.fields = make([][]string, len(data))
	g.marked = make([][]int, len(data))
	for y, line := range data {
		g.fields[y] = make([]string, len(line))
		g.marked[y] = make([]int, len(line))
		for x, ch := range line {
			g.fields[y][x] = string(ch)
		}
	}

	return g
}

func (g Grid) Get(p Point) string {
	/* soft boundary check, just return any rubish that
	 * is not XMAS
	 */
	if p.x < 0 || p.x >= len(g.fields) {
		return "#"
	}
	if p.y < 0 || p.y >= len(g.fields) {
		return "#"
	}

	return g.fields[p.y][p.x]
}

func (g Grid) String() string {
	out := ""
	max := 0
	for y := 0; y < len(g.fields); y++ {
		for x := 0; x < len(g.fields); x++ {
			count := g.marked[y][x]
			if count > 0 {
				if count > max {
					max = count
				}
				out += fmt.Sprintf("\033[0;%dm%s\033[0m ", 31+count, g.fields[y][x])
			} else {
				out += fmt.Sprintf("\033[0;2m%s\033[0m ", g.fields[y][x])
			}
		}
		out += "\n"
	}

	out += "\n"
	// print legend
	for i := 0; i <= max; i++ {
		if i == 0 {
			out += fmt.Sprintf("\033[0;2mno match\033[0m\n")
			continue
		}
		out += fmt.Sprintf("\033[0;%dm%d matches\033[0m\n", 31+i, i)
	}

	return out
}

func (g Grid) Iter() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for y := range g.fields {
			for x := range g.fields[y] {
				if !yield(x, y) {
					return
				}
			}
		}
	}
}

func (g Grid) Search(checks [][]Point) int {
	sum := 0
	for x, y := range g.Iter() {
		for _, check := range checks {
			isMatch := true
			cache := make([]Point, 0)
			for _, point := range check {
				if g.Get(point.Add(x, y)) != point.ch {
					isMatch = false
					break
				}
				cache = append(cache, point.Add(x, y))
			}
			if isMatch {
				sum += 1
				for _, p := range cache {
					g.marked[p.y][p.x] += 1
				}
			}
		}
	}

	return sum
}

func part1(data []string) int {
	grid := NewGrid(data)

	result := grid.Search(allChecks)
	fmt.Println(grid.String())
	return result
}

func part2(data []string) int {
	grid := NewGrid(data)

	result := grid.Search(allX)
	fmt.Println(grid.String())
	return result
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}

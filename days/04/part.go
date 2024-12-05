package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Grid struct {
	fields [][]string
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
	for i := 0; i <= max; i++ {
		if i == 0 {
			out += fmt.Sprintf("\033[0;2mno match\033[0m\n")
			continue
		}
		out += fmt.Sprintf("\033[0;%dm%d matches\033[0m\n", 31+i, i)
	}

	return out
}

type Point struct {
	x, y int
	ch   string
}

func (p Point) Add(x, y int) Point {
	p.x += x
	p.y += y
	return p
}

func (g Grid) Search() int {
	horizontal := []Point{
		Point{0, 0, "X"}, Point{1, 0, "M"}, Point{2, 0, "A"}, Point{3, 0, "S"},
	}
	vertical := []Point{
		Point{0, 0, "X"}, Point{0, 1, "M"}, Point{0, 2, "A"}, Point{0, 3, "S"},
	}
	diagonalDesc := []Point{
		Point{0, 0, "X"}, Point{1, 1, "M"}, Point{2, 2, "A"}, Point{3, 3, "S"},
	}
	diagonalAsc := []Point{
		Point{0, 0, "X"}, Point{-1, -1, "M"}, Point{-2, -2, "A"}, Point{-3, -3, "S"},
	}
	horizontalR := []Point{
		Point{0, 0, "S"}, Point{1, 0, "A"}, Point{2, 0, "M"}, Point{3, 0, "X"},
	}
	verticalR := []Point{
		Point{0, 0, "S"}, Point{0, 1, "A"}, Point{0, 2, "M"}, Point{0, 3, "X"},
	}
	diagonalDescR := []Point{
		Point{0, 0, "S"}, Point{1, 1, "A"}, Point{2, 2, "M"}, Point{3, 3, "X"},
	}
	diagonalAscR := []Point{
		Point{0, 0, "S"}, Point{-1, -1, "A"}, Point{-2, -2, "M"}, Point{-3, -3, "X"},
	}

	checks := [][]Point{
		horizontal, vertical, diagonalDesc, diagonalAsc,
		horizontalR, verticalR, diagonalDescR, diagonalAscR,
	}

	sum := 0
	for y, line := range g.fields {
		for x := range line {
			//if field == "X" || field == "S" {
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
				//}
			}
		}
	}

	return sum
}

func part1(data []string) int {
	grid := NewGrid(data)

	result := grid.Search()
	fmt.Println(grid.String())
	return result
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

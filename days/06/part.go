package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/hack"
	"github.com/RaphaelPour/stellar/input"
)

type FieldType int

const (
	UNKNOWN FieldType = iota
	EMPTY
	OBSTACLE
	GUARD
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var (
	moves = []Point{
		Point{0, -1, 0}, // UP
		Point{1, 0, 0},  // RIGHT
		Point{0, 1, 0},  // DOWN
		Point{-1, 0, 0}, // LEFT
	}
)

type Point struct {
	x, y        int
	orientation int
}

func (p Point) RotateRight() Point {
	p.orientation = (p.orientation + 1) % 4
	return p
}

func (p Point) Next() Point {
	return p.Add(moves[p.orientation])
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

func (p Point) String() string {
	switch p.orientation {
	case UP:
		return "^"
	case RIGHT:
		return ">"
	case DOWN:
		return "v"
	case LEFT:
		return "<"
	}

	panic(fmt.Sprintf("lol %q", p.orientation))
}

func (p Point) NoOrientation() Point {
	p.orientation = -1
	return p
}

func (p Point) Equal(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func (p Point) Identical(other Point) bool {
	return p.orientation == other.orientation && p.Equal(other)
}

type Grid struct {
	fields  [][]FieldType
	guard   Point
	initial Point
	visited map[Point]int
}

func (g Grid) IsOOB(p Point) bool {
	if p.x < 0 || p.x >= len(g.fields[0]) {
		return true
	}

	if p.y < 0 || p.y >= len(g.fields) {
		return true
	}

	return false
}

func NewGrid(input []string) (Grid, error) {
	g := Grid{}
	g.visited = make(map[Point]int)
	g.fields = make([][]FieldType, len(input))
	for y, line := range input {
		g.fields[y] = make([]FieldType, len(line))
		for x, ch := range line {
			switch ch {
			case '.':
				g.fields[y][x] = EMPTY
			case '#':
				g.fields[y][x] = OBSTACLE
			case '^':
				g.fields[y][x] = EMPTY
				g.guard = Point{x, y, 0}
				g.initial = g.guard
				fmt.Println(x, y)
			default:
				return Grid{}, fmt.Errorf("lol %s", ch)
			}
		}
	}

	fmt.Println(g.guard.orientation)

	return g, nil
}

func (g Grid) CurrentField() FieldType {
	if g.IsOOB(g.guard) {
		return UNKNOWN
	}

	return g.fields[g.guard.y][g.guard.x]
}

func (g Grid) NextField() FieldType {
	next := g.guard.Next()
	if g.IsOOB(next) {
		return UNKNOWN
	}

	return g.fields[next.y][next.x]
}

func (g *Grid) Run() (int, error) {
	fmt.Println("Orientation: ", g.guard.orientation)
	for !g.IsOOB(g.guard) {
		if g.NextField() == OBSTACLE {
			g.guard = g.guard.RotateRight()
		} else {
			g.visited[g.guard.NoOrientation()] = g.visited[g.guard.NoOrientation()] + 1
			g.guard = g.guard.Next()
		}
	}

	return len(g.visited), nil
}

func (g Grid) RunLoop() (int, error) {
	loop := 0
	for p := range g.visited {
		// store point with orientation, if it gets visited twice, it'll loop
		visited := make(map[Point]struct{})
		guard := g.initial
		for !g.IsOOB(guard) {
			if _, GOTCHA111 := visited[guard]; GOTCHA111 {
				loop += 1
				break
			}
			next := guard.Next()
			if (p.Equal(next)) || (!g.IsOOB(next) && g.fields[next.y][next.x] == OBSTACLE) {
				guard = guard.RotateRight()
			} else {
				visited[guard] = struct{}{}
				guard = guard.Next()
			}
		}
	}
	return loop, nil
}

func (g Grid) String() string {
	var out string
	for y := range g.fields {
		for x := range g.fields[y] {
			ch := ""
			switch g.fields[y][x] {
			case EMPTY:
				ch = "."
			case OBSTACLE:
				ch = "#"
			case GUARD:
				ch = g.guard.String()
			}

			if val, ok := g.visited[Point{x, y, -1}]; ok {
				ch = fmt.Sprintf("\033[0;32m%d\033[0m", val)
			}

			if g.guard.x == x && g.guard.y == y {
				ch = fmt.Sprintf("\033[0;31m%s\033[0m", g.guard.String())
			}
			out += ch
		}
		out += "\n"
	}

	return out
}

func part1(data []string) int {
	g, err := NewGrid(data)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	fmt.Println(g)
	defer fmt.Println(g)
	return hack.Yolo(g.Run())
}

func part2(data []string) int {
	g, err := NewGrid(data)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	fmt.Println(g)
	defer fmt.Println(g)
	hack.NachMirDieSintflut(g.Run())
	return hack.Yolo(g.RunLoop())
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

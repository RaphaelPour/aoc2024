package main

import (
	"fmt"
	"iter"

	"github.com/RaphaelPour/stellar/input"
	m "github.com/RaphaelPour/stellar/math"
)

type Direction int

func (d Direction) Add(p m.Point) m.Point {
	switch d {
	case NORTH:
		p.Y -= 1
	case EAST:
		p.X += 1
	case SOUTH:
		p.Y += 0
	case WEST:
		p.X -= 1
	}

	return p
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (d Direction) Diff(other Direction) int {
	return (Abs(int(d)-int(other)) + 1) % 4
}

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

var (
	directions = []Direction{NORTH, EAST, SOUTH, WEST}
)

type Grid struct {
	fields     [][]bool
	start, end m.Point
	visited    map[m.Point]int
}

func NewGrid(data []string) Grid {
	g := Grid{}
	g.fields = make([][]bool, len(data))
	for y, line := range data {
		g.fields[y] = make([]bool, len(line))
		for x, field := range line {
			g.fields[y][x] = string(field) != "#"

			if string(field) == "S" {
				g.start = m.Point{X: x, Y: y}
			} else if string(field) == "E" {
				g.end = m.Point{X: x, Y: y}
			}
		}
	}

	return g
}

func (g Grid) Neighbors(p m.Point) iter.Seq2[Direction, m.Point] {
	return func(yield func(Direction, m.Point) bool) {
		for _, dir := range directions {
			newPoint := dir.Add(p)
			if newPoint.X < 0 || newPoint.Y < 0 {
				continue
			}

			if newPoint.Y >= len(g.fields) || newPoint.X >= len(g.fields[0]) {
				continue
			}

			if !yield(dir, newPoint) {
				return
			}
		}
	}
}

func (g *Grid) IDFS(pos m.Point, dir Direction, score, maxScore int) int {
	if score > maxScore {
		return -1
	}

	// Don't reject duplicates, but paths with greater score
	if visitedScore, visited := g.visited[pos]; visited && score > visitedScore {
		return -1
	}

	if pos.Equal(g.end) {
		return score
	}

	g.visited[pos] = score

	minScore := 0
	for newDir, neigh := range g.Neighbors(pos) {
		fmt.Println(pos, neigh, newDir)
		if newScore := g.IDFS(neigh, newDir, score+newDir.Diff(dir)+1, maxScore); newScore > -1 && newScore < minScore {
			minScore = newScore
		}
	}

	return minScore
}

func (g *Grid) Search() int {
	for score := 1; score < ((len(g.fields)-2)*(len(g.fields[0])-2))*1000; score += 1 {
		g.visited = make(map[m.Point]int)
		if result := g.IDFS(g.start, EAST, 0, score); result > -1 {
			return result
		}
	}
	return -1
}

func part1(data []string) int {
	g := NewGrid(data)
	return g.Search()
}

func part2(data []string) int {
	return 0
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadString("input_example")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}

package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Grid struct {
	fields []string
}

func (g Grid) OtherNeighbors(xc, yc int, me string) int {
	sum := 0
	for y := yc - 1; y <= yc+1; y += 1 {
		for x := xc - 1; x <= xc+1; x += 1 {
			if g.Get(x, y) == me {
				sum += 1
			}
		}
	}

	return sum
}

func (g Grid) Get(x, y int) string {
	if x < 0 || y < 0 {
		return ""
	}

	if x >= len(g.fields[0]) || y >= len(g.fields) {
		return ""
	}
	return string(g.fields[y][x])
}

type Plant struct {
	Area, Perimeter int
}

func (p Plant) Price() int {
	return p.Area * p.Perimeter
}

func (g Grid) Costs() int {
	hist := make(map[string]Plant)

	keys := make([]string, 0)
	for y := 0; y < len(g.fields); y++ {
		for x := 0; x < len(g.fields[0]); x++ {
			fmt.Println(g.Get(x, y), x, y)
			plant, ok := hist[g.Get(x, y)]
			if !ok {
				keys = append(keys, g.Get(x, y))
			}
			plant.Area += 1
			plant.Perimeter += g.OtherNeighbors(x, y, g.Get(x, y))
			hist[g.Get(x, y)] = plant
		}
	}

	price := 0
	for _, name := range keys {
		plant := hist[name]
		cost := plant.Price()
		fmt.Printf("A region of %s plants with price %d * %d = %d\n", name, plant.Area, plant.Perimeter, cost)
		price += cost
	}
	return price
}

func part1(data []string) int {
	grid := Grid{fields: data}
	return grid.Costs()
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

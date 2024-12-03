package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/strings"
)

func parts(data []string) (int, int) {
	pattern := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`)
	sum1 := 0
	sum2 := 0
	enabled := true
	for _, line := range data {
		for _, matches := range pattern.FindAllStringSubmatch(line, -1) {
			if matches[0] == "do()" {
				enabled = true
			} else if matches[0] == "don't()" {
				enabled = false
			} else {

				summand := strings.ToInt(matches[2]) * strings.ToInt(matches[3])
				sum1 += summand

				if enabled {
					sum2 += summand
				}
			}
		}
	}
	return sum1, sum2
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadString("input")
	p1, p2 := parts(data)
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(p1)

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(p2)
}

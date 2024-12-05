package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	sstrings "github.com/RaphaelPour/stellar/strings"
)

var (
	EmptyRule = Rule{-1, -1}
)

type Rule struct {
	former, later int
}

func NewRule(s string) (Rule, error) {
	fields := strings.Split(s, "|")
	if len(fields) != 2 {
		return EmptyRule, fmt.Errorf(
			"expected rule %q to have two fields, got %d",
			s,
			len(fields),
		)
	}

	return Rule{
		sstrings.ToInt(fields[0]),
		sstrings.ToInt(fields[1]),
	}, nil
}

func (r Rule) Validate(n []int) bool {
	formerIndex, laterIndex := -1, -1

	for i, num := range n {
		if num == r.former {
			formerIndex = i
		}
		if num == r.later {
			laterIndex = i
		}
	}

	if formerIndex == -1 {
		return true
	}
	if laterIndex == -1 {
		return true
	}
	return formerIndex < laterIndex
}

func Parse(s []string) ([][]int, []Rule, error) {
	rules := make([]Rule, 0)
	for ; len(s) > 0 && s[0] != ""; s = s[1:] {
		rule, err := NewRule(s[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing rule: %w", err)
		}
		rules = append(rules, rule)
	}

	// skip empty line
	s = s[1:]

	numberLists := make([][]int, 0)
	for ; len(s) > 0 && s[0] != ""; s = s[1:] {
		numberList := make([]int, 0)
		for _, num := range strings.Split(s[0], ",") {
			numberList = append(numberList, sstrings.ToInt(num))
		}
		numberLists = append(numberLists, numberList)
	}

	return numberLists, rules, nil
}

func part1(data []string) int {
	pagesList, rules, err := Parse(data)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	fmt.Println("Pageslist:", pagesList)
	fmt.Println("Rules:", rules)

	sum := 0
	for _, pages := range pagesList {
		valid := true
		for _, rule := range rules {
			if !rule.Validate(pages) {
				valid = false
				break
			}
		}
		if valid {
			fmt.Println(pages, "valid, add", pages[len(pages)/2])
			sum += pages[len(pages)/2]
		}
	}
	return sum
}

func part2(data []string) int {
	return 0
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}

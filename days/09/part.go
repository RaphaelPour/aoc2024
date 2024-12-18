package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

func Index(in []rune, needle rune, index int) int {
	for i := index; i < len(in); i++ {
		if in[i] == needle {
			return i
		}
	}
	return -1
}

func LastIndexNot(in []rune, needle rune, index int) int {
	for i := index; i >= 0; i-- {
		if in[i] != needle {
			return i
		}
	}
	return -1
}

func part1(data []string) int {
	var diskMapS string
	var id int
	var free int
	for i, ch := range data[0] {
		field := strconv.Itoa(id)
		count := int(ch - '0')
		if i&1 == 1 {
			field = "."
			free += count
		} else {
			id += 1
		}
		diskMapS += strings.Repeat(field, count)
	}

	diskMap := []rune(diskMapS)
	for {
		nextFree := Index(diskMap, '.', 0)
		if nextFree == -1 {
			panic(fmt.Sprintf("next free: %d", nextFree))
		}
		lastBlock := LastIndexNot(diskMap, '.', len(diskMap)-1)
		if lastBlock == -1 {
			panic(fmt.Sprintf("last block: %d", lastBlock))
		}
		if nextFree > lastBlock {
			break
		}
		diskMap[nextFree], diskMap[lastBlock] = diskMap[lastBlock], diskMap[nextFree]
		delta := (lastBlock - nextFree)
		if delta%1000 == 0 {
			fmt.Println(delta)
		}
	}

	var result int
	for pos, id := range diskMap {
		if id == '.' {
			continue
		}
		//fmt.Printf("%d * %d = %d\n", pos, id-'0', pos*int(id-'0'))
		result += pos * int(id-'0')
	}

	//fmt.Println(string(diskMap))

	return result
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
	result := part1(data)

	if result <= 89860529514 {
		fmt.Println("too low")
	}
	fmt.Println(result)
	// fmt.Println("== [ PART 2 ] ==")
	//fmt.Println(part2(data))
}

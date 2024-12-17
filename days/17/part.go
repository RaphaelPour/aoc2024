package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	m "github.com/RaphaelPour/stellar/math"
	s "github.com/RaphaelPour/stellar/strings"
)

func M[in any, out any, inArr []in, outArr []out](input inArr, fn func(in) out) outArr {
	output := make(outArr, len(input))
	for i := range input {
		output[i] = fn(input[i])
	}
	return output
}

type VM struct {
	a, b, c int
	ip      int
	program []int
	output  []int
}

type Operand func(vm *VM) int

func Literal(l int) Operand {
	return func(_ *VM) int {
		return l
	}
}

func Register(r int) Operand {
	return func(vm *VM) int {
		switch r {
		case 4:
			return vm.a
		case 5:
			return vm.b
		case 6:
			return vm.c
		default:
			panic(r)
		}
	}
}

func BOOM(n int) Operand {
	return func(_ *VM) int {
		panic(fmt.Sprintf("%d is not a valid Operand", n))
	}
}

var (
	comboOperands = []Operand{
		Literal(0),
		Literal(1),
		Literal(2),
		Literal(3),
		Register(4),
		Register(5),
		Register(6),
		BOOM(7),
	}
)

type Instruction func(vm *VM, arg int)

func ADV(vm *VM, arg int) {
	numerator := vm.a
	denominator := m.Pow(2, comboOperands[arg](vm))
	vm.a = int(float64(numerator / denominator))
	vm.ip += 2
}

func BXL(vm *VM, arg int) {
	vm.b ^= arg
	vm.ip += 2
}

func BST(vm *VM, arg int) {
	vm.b = comboOperands[arg](vm) % 8
	vm.ip += 2
}

func JNZ(vm *VM, arg int) {
	if vm.a == 0 {
		vm.ip += 2
		return
	}

	vm.ip = arg
}

func BXC(vm *VM, _ int) {
	vm.b ^= vm.c
	vm.ip += 2
}

func OUT(vm *VM, arg int) {
	vm.output = append(vm.output, comboOperands[arg](vm))
	vm.ip += 2
}

func BDV(vm *VM, arg int) {
	numerator := vm.a
	denominator := m.Pow(2, comboOperands[arg](vm))
	vm.b = int(float64(numerator / denominator))
	vm.ip += 2
}

func CDV(vm *VM, arg int) {
	numerator := vm.a
	denominator := m.Pow(2, comboOperands[arg](vm))
	vm.c = int(float64(numerator / denominator))
	vm.ip += 2
}

var (
	instructions = []Instruction{ADV, BXL, BST, JNZ, BXC, OUT, BDV, CDV}
)

func NewVM(data []string) VM {
	vm := VM{}
	vm.program = make([]int, 0)
	for _, line := range data {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		switch parts[0] {
		case "Register A":
			vm.a = s.ToInt(strings.TrimSpace(parts[1]))
		case "Register B":
			vm.b = s.ToInt(strings.TrimSpace(parts[1]))
		case "Register C":
			vm.c = s.ToInt(strings.TrimSpace(parts[1]))
		case "Program":
			vm.program = M(strings.Split(strings.TrimSpace(parts[1]), ","), s.ToInt)
		}
	}

	return vm
}

func (vm VM) String() string {
	return strings.Join(M(vm.output, strconv.Itoa), ",")
}

func (vm *VM) Run() {
	fmt.Printf(
		"a=%d b=%d c=%d ip=%d out=%s program=%v\n",
		vm.a,
		vm.b,
		vm.c,
		vm.ip,
		vm.String(),
		vm.program,
	)
	for vm.ip < len(vm.program) {
		fmt.Printf(
			"%d %d",
			vm.program[vm.ip],
			vm.program[vm.ip+1],
		)
		instructions[vm.program[vm.ip]](vm, vm.program[vm.ip+1])
		fmt.Printf(
			"-> a=%d b=%d c=%d ip=%d out=%s\n",
			vm.a,
			vm.b,
			vm.c,
			vm.ip,
			vm.String(),
		)
	}
}

func part1(data []string) string {
	vm := NewVM(data)
	vm.Run()
	return vm.String()
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input_example")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}

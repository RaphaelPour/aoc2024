package main

import (
	"fmt"
	"iter"
	"math"
	"strconv"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	m "github.com/RaphaelPour/stellar/math"
	s "github.com/RaphaelPour/stellar/strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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

func Spread(start, end, chunks int) iter.Seq2[int, int] {
	total := end - start
	chunkSize := total / chunks
	return func(yield func(int, int) bool) {
		for i := start; i <= end; i += chunkSize {
			if !yield(i, m.Min(i+chunkSize, end)) {
				return
			}
		}
	}
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
	vm.a /= m.Pow(2, comboOperands[arg](vm))
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
	vm.output = append(vm.output, comboOperands[arg](vm)%8)
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
	if denominator == 0 {
		panic(fmt.Sprintf(
			"integer divide by zero: denominator=%d combo=%d arg=%d a=%d b=%d c=%d",
			denominator,
			comboOperands[arg](vm), arg,
			vm.a, vm.b, vm.c),
		)
	}
	vm.c = int(float64(numerator / denominator))
	vm.ip += 2
}

var (
	instructions = []Instruction{ADV, BXL, BST, JNZ, BXC, OUT, BDV, CDV}
	//                           0    1    2    3    4    5    6    7
)

func λ(a int, goal string) string {
	b := 0
	c := 0
	out := ""

	for a != 0 { // 3,0 -> jnz 0 -> jmp to 0x0 if a == 0
		b = a % 8        // 2,4 -> bst a -> b = a mod 8
		b ^= 1           // 1,1 -> bxl 1 -> b = b xor 1
		c = a / (1 << b) // 7,5 -> cdv c -> c = a/(2**b)
		a /= 8           // 0,3 -> adv 3 -> a = a/(2**3)
		b ^= 4           // 1,4 -> bxl 4 -> b = b xor 4
		b ^= c           // 4,5 -> bxc _ -> b = b xor c
		// out += string(rune((b%8)+'0')) + "," // 5,5 -> out b -> append b to output
		ch := string(rune((b % 8) + '0'))
		if string(goal[len(out)]) != ch {
			break
		}
		out += ch + ","
	}

	return out[:max(0, len(out)-1)]
}

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

func (vm VM) IsQuine() bool {
	if len(vm.program) != len(vm.output) {
		return false
	}

	for i, ins := range vm.program {
		if vm.output[i] != ins {
			return false
		}
	}

	return true
}

func (vm VM) IsSoftQuine() bool {
	for i, ins := range vm.output {
		if vm.program[i] != ins {
			return false
		}
	}
	return true
}

func (vm VM) Matches() int {
	match := 0
	for i, ins := range vm.output {
		if vm.program[i] == ins {
			match += 1
		}
	}
	return match
}

func (vm VM) String() string {
	return strings.Join(M(vm.output, strconv.Itoa), ",")
}

func (vm *VM) Run() {
	for vm.ip < len(vm.program) {
		instructions[vm.program[vm.ip]](vm, vm.program[vm.ip+1])
	}
}

func (vm *VM) Run2() {
	for vm.ip < len(vm.program) && vm.IsSoftQuine() {
		instructions[vm.program[vm.ip]](vm, vm.program[vm.ip+1])
	}
}

func part1(data []string) string {
	vm := NewVM(data)
	vm.Run()
	return vm.String()
}

/*
func Solve(bytlePos int, goal string, bytleCombinations [][]int) []int {
	for i := 0; i < 8; i++ {
		candidate := i << bytlePos
		for j, predefBytles := range bytleCombinations {
			for k, predefBytle := range predefBytles {
				candidate
			}
		}
	}
}*/

func I(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for t21 := 0; t21 < 8; t21++ {
			for t20 := 0; t20 < 8; t20++ {
				for t19 := 0; t19 < 8; t19++ {
					for t18 := 0; t18 < 8; t18++ {
						for t17 := 0; t17 < 8; t17++ {
							for t16 := 0; t16 < 8; t16++ {
								for t15 := 0; t15 < 8; t15++ {
									for t14 := 0; t14 < 8; t14++ {
										for t13 := 0; t13 < 8; t13++ {
											for t12 := 0; t12 < 8; t12++ {
												for _, t1 := range []int{0b010, 0b101, 0b111} {
													//for t1 := 0; t1 < 8; t1++ {
													t2 := 0b101
													//for _, t2 := range []int{0b101, 0b110, 0b111} {
													//for t3 := 0; t3 < 8; t3++ {
													t3 := 0
													t4 := 0b101
													t5 := 0b010
													t6 := 0
													t7 := 0b110
													for _, t8 := range []int{0b101, 0b110} {
														for _, t9 := range []int{0b110, 0b111} {
															for _, t10 := range []int{0, 2, 5, 6} {
																t11 := 3

																i := (t1 | t2<<3 | t3<<6 | t4<<9 | t5<<12 | t6<<15 | t7<<18 | t8<<21 | t9<<24 | t10<<27)
																i |= t11<<30 | t12<<33 | t13<<36 | t14<<39 | t15<<42 | t16<<45 | t17<<48 | t18<<51 | t19<<54 | t20<<57 | t21<<60
																if i < 0 {
																	panic(fmt.Sprintf("i=%d", i))
																}
																if !yield(i) {
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func hexdump(n int) string {
	var out string
	tmp := fmt.Sprintf("%064b", uint(n)<<1)
	for j := 0; j < len(tmp)-1; j += 3 {
		out += tmp[j:j+3] + " "
	}
	return out
}

func similarity(a, b string) int {
	var result int
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			break
		}

		if a[i] != ',' {
			result += 1
		}
	}
	return result
}

func part2(data []string) int {
	rounds := 0

	procs := 1 //runtime.GOMAXPROCS(0)

	var result int
	done := make(chan struct{})
	goal := "2,4,1,1,7,5,0,3,1,4,4,5,5,5,3,0"

	for start, end := range Spread(0, math.MaxInt32, procs) {
		go func() {
			hist := make(map[int]int)
			for i := range I(start, end) {
				if rounds%10000000 == 0 {
					fmt.Printf("beep %d\n", rounds)
				}
				rounds += 1
				/*
					vm := NewVM(data)
					vm.a = i
					vm.Run2()
					if vm.Matches() > 8 {
						fmt.Printf("%s %s\n", hexdump(i), vm)
					}
					if vm.IsQuine() {
						result = i
						done <- struct{}{}
						break
					}
				*/
				out := λ(i, goal)
				if out == goal {
					result = i
					done <- struct{}{}
					break
				}

				s := similarity(out, goal)
				if s > 10 {
					shift := 10 * 3
					hist[(i>>shift)&7] = hist[(i>>shift)&7] + 1
					fmt.Println(hist)
					fmt.Printf("%s %s\n", hexdump(i), out)
				}
			}
		}()
	}

	<-done

	return result
}

func main() {
	data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}

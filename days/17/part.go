package main

import (
	"fmt"
	"iter"
	"math"
	"runtime"
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

func λ(a int) string {
	b := 0
	c := 0
	out := ""

	for a != 0 { // 3,0 -> jnz 0 -> jmp to 0x0 if a == 0
		b = a % 8                            // 2,4 -> bst a -> b = a mod 8
		b ^= 1                               // 1,1 -> bxl 1 -> b = b xor 1
		c = a / (1 << b)                     // 7,5 -> cdv c -> c = a/(2**b)
		a /= 8                               // 0,3 -> adv 3 -> a = a/(2**3)
		b ^= 4                               // 1,4 -> bxl 4 -> b = b xor 4
		b ^= c                               // 4,5 -> bxc _ -> b = b xor c
		out += string(rune((b%8)+'0')) + "," // 5,5 -> out b -> append b to output
	}

	return out[:len(out)-1]
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

func I(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for r := start; r < end; r++ {
			//i := r << 31

			for _, t1 := range []int{0b010, 0b101, 0b111} {
				//i |= (t1 << 0)

				/*
					t1 := (i >> 0) & 7
					if t1 != 0b010 && t1 != 0b101 && t1 != 0b111 {
						continue
					}*/

				//i |= (0b101 << 3)
				t2 := 0b101
				/*t2 := (i >> 3) & 7
				if t2 != 0b101 {
					continue
				}*/

				//for t3 := 0; t3 < 8; t3++ {
				t3 := 0
				/*
					t3 := (i >> 6) & 7
					if t3 != 0 {
						continue
					}*/

				//i |= (0b101 << 9)
				t4 := 0b101
				/*
					t4 := (i >> 9) & 7
					if t4 != 0b101 {
						continue
					}*/

				//i |= (0b010 << 12)
				t5 := 0b010
				/*
					t5 := (i >> 12) & 7
					if t5 != 0b010 {
						continue
					}*/

				/*
					t6 := (i >> 15) & 7
					if t6 != 0 {
						continue
					}*/
				//for t6 := 0; t6 < 8; t6++ {
				t6 := 0

				//for _, t7 := range []int{0b100, 0b110} {
				t7 := 0b110
				//i |= (t7 << 18)
				/*
					t7 := (i >> 18) & 7
					if t7 != 0b100 && t7 != 0b110 {
						continue
					}*/

				//for _, t8 := range []int{0b110, 0b010} {
				t8 := 0b110
				//i |= (t8 << 21)
				/*
					t8 := (i >> 21) & 7
					if t8 != 0b110 && t8 != 0b010 {
						continue
					}*/

				//	for _, t9 := range []int{0b100, 0b11, 0b111} {
				t9 := 0b111
				//i |= (t9 << 24)

				/*t9 := (i >> 24) & 7
				if t9 != 0 && t9 != 0b100 && t9 != 0b110 && t9 != 0b111 {
					continue
				}*/

				//for _, t10 := range []int{0b000, 0b001} {
				//	for t10 := 0; t10 < 8; t10++ {
				//i |= (t10 << 27)
				/*
					t10 := (i >> 27) & 7
					if t10 != 1 && t10 != 0 {
						continue
					}*/

				//for _, t11 := range []int{0b000, 0b010, 0b100, 0b110} {
				//for t11 := 0; t11 < 8; t11++ {
				//i |= (t11 << 31)
				/*
					t11 := (i >> 31) & 7
					if t11&2 == 0 {
						continue
					}*/

				i := r<<31 | (t1 | t2<<3 | t3<<6 | t4<<9 | t5<<12 | t6<<15 | t7<<18 | t8<<21 | t9<<24 | 0<<27)
				if i < 0 {
					panic(fmt.Sprintf("r=%d i=%d", r, i))
				}
				if !yield(i) {
					return
				}
				//}
				//}
				//	}
				//			}
				//		}
				//	}
				//	}
			}
		}
	}
}

func hexdump(n int) string {
	var out string
	tmp := fmt.Sprintf("%064b", n<<1)
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

	procs := runtime.GOMAXPROCS(0)

	var result int
	done := make(chan struct{})
	goal := "2,4,1,1,7,5,0,3,1,4,4,5,5,5,3,0"
	for start, end := range Spread(0, math.MaxInt32, procs) {
		go func() {
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
				out := λ(i)
				if out == goal {
					result = i
					done <- struct{}{}
					break
				}
				if s := similarity(out, goal); s > 8 {
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

package day17

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type cpu struct {
	a, b, c int
	pc      int
	memory  []int
	output  []int
}

func (c *cpu) Clone() *cpu {
	return &cpu{
		a:      c.a,
		b:      c.b,
		c:      c.c,
		memory: slices.Clone(c.memory),
		output: slices.Clone(c.output),
	}
}

var useComboOperand = aoc.SetWithValues([]int{0, 2, 5, 6, 7})
var opcodeToName = map[int]string{0: "adv", 1: "bxl", 2: "bst", 3: "jnz", 4: "bxc", 5: "out", 6: "bdv", 7: "scdv"}

func findValidStart(input string) int {
	initial := parse(input)
	n := len(initial.memory)
	program := slices.Clone(initial.memory)

outer:
	for i := 0; i < math.MaxInt64; i++ {
		c := initial.Clone()
		c.a = i

		for steps := 0; steps < 100000; steps++ {
			if i%100_000 == 0 {
				fmt.Printf("i: %d, steps: %d\n", i, steps)
			}
			ok := c.step()

			if len(c.output) > n {
				continue outer
			} else if len(c.output) == n && slices.Equal(c.output, program) {
				return i
			} else if len(c.output) > 0 && !slices.Equal(c.output, program[:len(c.output)]) {
				continue outer
			}

			if !ok {
				break
			}
		}
	}

	return -1
}

func (c *cpu) Run() string {
	for c.step() {

	}

	return aoc.JoinInts(c.output, ",")
}

func (c *cpu) step() bool {
	if c.pc >= len(c.memory) {
		return false
	}

	instr := c.memory[c.pc]
	n := c.memory[c.pc+1]
	s := opcodeToName[instr]
	_ = s
	operand := c.operand(instr, n)
	switch instr {
	case 0:
		// adv
		c.a = int(c.a / int(math.Pow(2, float64(operand))))
		c.pc += 2
	case 1:
		// bxl
		c.b ^= operand
		c.pc += 2
	case 2:
		// bst
		c.b = operand % 8
		c.pc += 2
	case 3:
		// jnz
		if c.a != 0 {
			c.pc = operand
		} else {
			c.pc += 2
		}
	case 4:
		// bxc
		c.b ^= c.c
		c.pc += 2
	case 5:
		// out
		o := operand % 8
		c.output = append(c.output, o)
		c.pc += 2
	case 6:
		// bdv
		c.b = int(c.a / int(math.Pow(2, float64(operand))))
		c.pc += 2
	case 7:
		// cdv
		c.c = int(c.a / int(math.Pow(2, float64(operand))))
		c.pc += 2
	}

	return true
}

func (c *cpu) operand(instr, op int) int {
	if !useComboOperand.Contains(instr) {
		return op
	}
	if op >= 0 && op <= 3 {
		return op
	}

	switch op {
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	case 7:
		return op
	default:
		panic(fmt.Errorf("unexpected combo operand op: %d", op))
	}
}

func (c *cpu) write(addr, value int) {
	c.memory[addr] = value
}

func (c *cpu) read(addr int) int {
	return c.memory[addr]
}

func parse(input string) *cpu {
	c := cpu{output: []int{}}
	for _, line := range aoc.Lines(input) {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Split(line, ": ")
		if strings.HasPrefix(line, "Register") {
			reg := strings.ReplaceAll(parts[0], "Register ", "")
			switch reg {
			case "A":
				c.a = aoc.MustAtoi(parts[1])
			case "B":
				c.b = aoc.MustAtoi(parts[1])
			case "C":
				c.c = aoc.MustAtoi(parts[1])
			}
		} else if strings.HasPrefix(line, "Program:") {
			c.memory = aoc.Ints(parts[1], ",")
		}
	}
	return &c
}

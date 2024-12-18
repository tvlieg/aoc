package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type opcode int

const (
	ADV opcode = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

type computer struct {
	a, b, c int
	stdout  []int
}

func (c *computer) String() string {
	return fmt.Sprintf("a: %d\tb: %d\tc: %d", c.a, c.b, c.c)
}

func newComputer(a, b, c int) *computer {
	return &computer{
		a: a,
		b: b,
		c: c,
	}
}

func (c *computer) compute(program []int) {
	var ip int
	for ip < len(program)-1 {
		oc := opcode(program[ip])
		op := program[ip+1]

		ip = c.operate(ip, oc, op)
	}
}

func (c *computer) print() string {
	var b strings.Builder
	for i, n := range c.stdout {
		b.WriteString(strconv.Itoa(n))
		if i < len(c.stdout)-1 {
			b.WriteRune(',')
		}
	}
	return b.String()
}

func (c *computer) operate(ip int, oc opcode, op int) int {
	switch oc {
	case ADV:
		c.adv(op)
	case BXL:
		c.bxl(op)
	case BST:
		c.bst(op)
	case JNZ:
		if newIP, ok := c.jnz(op); ok {
			return newIP
		}
	case BXC:
		c.bxc(op)
	case OUT:
		c.out(op)
	case BDV:
		c.bdv(op)
	case CDV:
		c.cdv(op)
	}

	return ip + 2
}

func (c *computer) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	default:
		panic(fmt.Sprintf("invalid operand: %d", operand))
	}
}

// The adv instruction (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A register.
func (c *computer) adv(op int) {
	c.a = dv(c.a, c.combo(op))
}

// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the result in register B.
func (c *computer) bxl(op int) {
	c.b = c.b ^ op
}

// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.
func (c *computer) bst(op int) {
	c.b = c.combo(op) % 8
}

// The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.
func (c *computer) jnz(op int) (int, bool) {
	if c.a == 0 {
		return 0, false
	}
	return op, true
}

// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)
func (c *computer) bxc(_ int) {
	c.b = c.b ^ c.c
}

// The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value. (If a program outputs multiple values, they are separated by commas.)
func (c *computer) out(op int) {
	c.stdout = append(c.stdout, c.combo(op)%8)
}

// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register. (The numerator is still read from the A register.)
func (c *computer) bdv(op int) {
	c.b = dv(c.a, c.combo(op))
}

// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register. (The numerator is still read from the A register.)
func (c *computer) cdv(op int) {
	c.c = dv(c.a, c.combo(op))
	// fmt.Printf("op: %d; a: %b; c: %b\n", c.b, c.a, c.c)
}

func dv(nom, denom int) int {
	return int(float64(nom) / math.Pow(2, float64(denom)))
}

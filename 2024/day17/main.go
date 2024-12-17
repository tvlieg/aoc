package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var (
	//go:embed example_input
	example string

	//go:embed input
	input string

	//go:embed example_input_2
	example2 string
)

func main() {
	fmt.Println("Part 1 (example):", part1(example))
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2 (example):", part2(example2))
	fmt.Println("Part 2:", part2(input))
}

func part1(input string) string {
	a, b, c, p := parseInput(input)
	computer := newComputer(a, b, c)
	computer.compute(p)
	return computer.print()
}

func part2(input string) int {
	_, _, _, p := parseInput(input)

	// n := int(math.Pow(8, 16)) + 1
	// m := int(math.Pow(8, 17))
	// for i := n; i <= m; i++ {

	for i := 0; ; i++ {
		c := newComputer(i, 0, 0)
		c.compute(p)

		// fmt.Println(c.print())
		if !slices.Equal(p, c.stdout) {
			continue
		}
		return i
	}
}

func parseInput(input string) (a, b, c int, p []int) {
	parseReg := func(line string) (reg int) {
		_, s, _ := strings.Cut(line, ": ")
		reg, _ = strconv.Atoi(s)
		return reg
	}
	parseProgram := func(line string) []int {
		_, s, _ := strings.Cut(line, ": ")
		numStrings := strings.Split(s, ",")
		program := make([]int, len(numStrings))
		for i, ns := range numStrings {
			n, _ := strconv.Atoi(ns)
			program[i] = n
		}
		return program
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		s := scanner.Text()
		switch {
		case strings.HasPrefix(s, "Register A: "):
			a = parseReg(s)
		case strings.HasPrefix(s, "Register B: "):
			b = parseReg(s)
		case strings.HasPrefix(s, "Register C: "):
			c = parseReg(s)
		case strings.HasPrefix(s, "Program: "):
			p = parseProgram(s)
		}
	}

	return a, b, c, p
}

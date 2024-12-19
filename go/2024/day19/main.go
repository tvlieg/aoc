package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed example_input
var example string

//go:embed input
var input string

func main() {
	fmt.Println("Part 1 (example):", part1(example))
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2 (example):", part2(example))
	fmt.Println("Part 2:", part2(input))
}

func part1(input string) int {
	towels, designs := parse(input)

	var count int
	for _, d := range designs {
		if canArrange(d, towels) {
			count++
		}
	}

	return count
}

func part2(input string) int {
	towels, designs := parse(input)

	var count int
	for _, d := range designs {
		count += combinations(d, towels)
	}

	return count
}

func parse(input string) (towels, designs []string) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		towels = strings.Split(line, ", ")
	}

	for scanner.Scan() {
		line := scanner.Text()
		designs = append(designs, line)
	}

	return towels, designs
}

func canArrange(design string, towels []string) bool {
	failed := make(map[string]struct{})
	return recurse(design, towels, failed)
}

func recurse(design string, towels []string, failed map[string]struct{}) bool {
	if design == "" {
		return true
	}
	if _, ok := failed[design]; ok {
		return false
	}

	for _, t := range towels {
		if len(design) < len(t) {
			continue
		}
		if design[:len(t)] == t {
			if recurse(design[len(t):], towels, failed) {
				return true
			}
			failed[design[len(t):]] = struct{}{}
		}
	}

	return false
}

func combinations(design string, towels []string) int {
	failed := make(map[string]int)
	return recurseComb(design, towels, failed)
}

func recurseComb(design string, towels []string, seen map[string]int) int {
	if design == "" {
		return 1
	}
	if n, ok := seen[design]; ok {
		return n
	}

	var count int
	for _, t := range towels {
		if len(design) < len(t) {
			continue
		}

		if design[:len(t)] == t {
			n := recurseComb(design[len(t):], towels, seen)
			seen[design[len(t):]] = n
			count += n
		}
	}

	return count
}

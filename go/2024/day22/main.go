package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example_input
var example string

//go:embed input
var input string

func main() {
	fmt.Println("Part 1 (example):", part1(example))
	fmt.Println("Part 1 :", part1(input))

	// fmt.Println(evolve(123))
}

func part1(input string) int {
	numbers := parse(input)

	var sum int
	for _, n := range numbers {
		// fmt.Printf("%d: ", n)
		for range 2000 {
			n = evolve(n)
		}
		// fmt.Println(n)
		sum += n
	}
	return sum
}

func parse(input string) (numbers []int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}
	return
}

func evolve(n int) int {
	n = ((n * 64) ^ n) % 16777216
	n = ((n / 32) ^ n) % 16777216
	return ((n * 2048) ^ n) % 16777216
}

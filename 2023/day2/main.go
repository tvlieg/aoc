package main

import (
	"bufio"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

func process(fileName string, fn func(string) int) (sum int, err error) {
	file, err := f.Open(fileName)
	if err != nil {
		fmt.Println("could not open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += fn(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w\n", err)
	}

	return sum, nil
}

func part1(s string) int {
	g := newGame(s)
	if g.isPossible() {
		return g.id
	}
	return 0
}

func part2(s string) int {
	g := newGame(s)
	return g.power()
}

func main() {
	// Part 1
	sum, err := process("input.txt", part1)
	if err != nil {
		fmt.Println("Part 1 failed:", err)
	}
	fmt.Println(sum)

	// Part 2
	sum, err = process("input.txt", part2)
	if err != nil {
		fmt.Println("Part 2 failed:", err)
	}
	fmt.Println(sum)
}
